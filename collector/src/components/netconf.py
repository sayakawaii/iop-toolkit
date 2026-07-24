import paramiko
from ncclient import manager
from ncclient.xml_ import to_ele
from ncclient.transport.errors import AuthenticationError
from concurrent.futures import ThreadPoolExecutor, as_completed
from models.host import Netconf
import logging

logging.basicConfig(
    level=logging.WARNING,
    format="%(asctime)s [%(levelname)s][%(name)s] %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S"
)
logger = logging.getLogger(__name__)
logger.setLevel(logging.WARNING)

logging.getLogger('ncclient').setLevel(logging.WARNING)
logging.getLogger('ncclient.operations').setLevel(logging.WARNING)
logging.getLogger('ncclient.transport').setLevel(logging.WARNING)

# 避免因服务器使用较旧的密钥类型而无法连接
paramiko.Transport._preferred_keys = ('ssh-rsa', 'ssh-dss')

class NetconfClient:
    def __init__(self, device : Netconf):
        self.device = device.as_params()
        self.original_device = device  # Keep reference to original device object
        self.manager = None
        self._connect_with_password_fallback()
    
    def _connect_with_password_fallback(self):
        """
        Establish NETCONF connection with concurrent password attempts.
        Try all possible passwords in parallel, use the first successful connection.
        """
        fallback_passwords = ['admin', 'Netconf#150']
        original_password = self.device['password']
        
        # Build list of passwords to try
        passwords_to_test = [original_password]
        for pwd in fallback_passwords:
            if pwd != original_password:
                passwords_to_test.append(pwd)
        
        logger.info(f"Attempting NETCONF connection with {len(passwords_to_test)} password(s) concurrently for {self.device['host']}")
        
        successful_connection = None
        correct_password = None
        
        with ThreadPoolExecutor(max_workers=len(passwords_to_test)) as executor:
            # Submit all connection attempts concurrently
            future_to_password = {
                executor.submit(self._try_connect, self.device['host'], self.device['port'], 
                               self.device['username'], pwd, self.device.get('hostkey_verify', False),
                               self.device.get('device_params'), self.device.get('timeout', 30)): pwd
                for pwd in passwords_to_test
            }
            
            # Check results as they complete
            for future in as_completed(future_to_password):
                password = future_to_password[future]
                try:
                    result = future.result()
                    if result is not None:
                        # Found successful connection!
                        successful_connection = result
                        correct_password = password
                        
                        if password != original_password:
                            logger.info(f"✓ NETCONF connection established with fallback password for {self.device['host']}")
                        else:
                            logger.info(f"✓ NETCONF connection established with original password for {self.device['host']}")
                        
                        # Cancel remaining futures
                        for f in future_to_password:
                            f.cancel()
                        break
                except Exception as e:
                    # This attempt failed, continue checking others
                    logger.debug(f"Connection attempt failed for {self.device['host']}: {e}")
                    continue
        
        if successful_connection and correct_password:
            self.manager = successful_connection
            self.device['password'] = correct_password
            self.original_device.password = correct_password
            logger.info(f"NETCONF connection established to {self.device['host']}")
            return
        
        # All attempts failed
        logger.error(f"All NETCONF connection attempts failed for {self.device['host']}")
        raise AuthenticationError("Authentication failed with all attempted passwords")
    
    def _try_connect(self, host, port, username, password, hostkey_verify, device_params, timeout):
        """
        Try to establish a single NETCONF connection with given credentials.
        Returns manager object if successful, None otherwise.
        """
        try:
            params = {
                'host': host,
                'port': port,
                'username': username,
                'password': password,
                'hostkey_verify': hostkey_verify,
                'timeout': timeout
            }
            if device_params:
                params['device_params'] = device_params
            
            m = manager.connect(**params)
            return m
        except AuthenticationError:
            # Authentication failed, return None
            return None
        except Exception as e:
            # Other errors should be raised
            logger.debug(f"Connection error for {host}:{port}: {e}")
            raise
    
    def _connect(self):
        """Establish NETCONF connection with current password"""
        try:
            if self.manager is not None:
                try:
                    self.manager.close_session()
                except:
                    pass
            self.manager = manager.connect(**self.device)
            logger.info(f"NETCONF connection established to {self.device['host']}")
        except Exception as e:
            logger.error(f"Failed to establish NETCONF connection: {e}")
            raise
    
    def _ensure_connected(self):
        """Ensure the connection is active, reconnect if needed"""
        if self.manager is None or not self.manager.connected:
            logger.warning(f"Connection lost to {self.device['host']}, reconnecting...")
            self._connect()
    
    def close(self):
        """Close the NETCONF connection"""
        if self.manager is not None:
            try:
                self.manager.close_session()
                logger.info(f"NETCONF connection closed to {self.device['host']}")
            except:
                pass
            finally:
                self.manager = None
    
    def __enter__(self):
        """Support context manager protocol"""
        return self
    
    def __exit__(self, exc_type, exc_val, exc_tb):
        """Support context manager protocol"""
        self.close()
        return False

    def dispatch(self, xml_str):
        self._ensure_connected()
        logging.info(f"Dispatching custom RPC:\n{xml_str}")
        logging.info(f"after to ele:\n{to_ele(xml_str)}")
        reply = self.manager.dispatch(to_ele(xml_str))
        return str(reply)

    def edit_config(self, xml_str):
        self._ensure_connected()
        reply = self.manager.edit_config(target="running", config=xml_str)
        return str(reply)

    def get(self, xml_str):
        self._ensure_connected()
        reply = self.manager.get(filter=xml_str)
        return str(reply)

    def get_capabilities(self):
        self._ensure_connected()
        self.manager.get()
        return f"\n".join(self.manager.server_capabilities)

    def get_config(self, xml_str):
        self._ensure_connected()
        reply = self.manager.get_config(source="running", filter=xml_str)
        return str(reply)