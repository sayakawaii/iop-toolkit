import logging
import xml.etree.ElementTree as ET
from concurrent.futures import ThreadPoolExecutor, as_completed
from components.netconf import NetconfClient
from utils.rpc import RpcGenerator
from utils.olttype import is_df_platform
from models.oltcomponent import OltComponent, ComponentState, OltInterface, OltInterfaceVani, OltInterfaceVaniOnuPresentOnThisOlt, PonInfo, LtInfo, NtInfo
from models.host import Netconf
from dataclasses import asdict

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s][%(name)s] %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S"
)
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)

lt_netconf_port_mapping = {
    "Slot-Lt-1": 833,
    "Slot-Lt-2": 834,
    "Slot-Lt-3": 835,
    "Slot-Lt-4": 836,
    "Slot-Lt-5": 837,
    "Slot-Lt-6": 838,
    "Slot-Lt-7": 839,
    "Slot-Lt-8": 840,
}

def get_netconf_port_by_lt_parent(lt_parent: str) -> int:
    return lt_netconf_port_mapping.get(lt_parent, 833)

def get_lt_list_from_response(response: str):
    return _parse_components_from_response(response, "lt")

def get_nt_list_from_response(response: str):
    return _parse_components_from_response(response, "nt")

def get_pon_list_from_response(response: str):
    return _parse_components_from_response(response, "transceiver")

def _parse_components_from_response(response: str, component_type: str) -> list:
    """
    Parse XML response and extract components of specified type (lt, nt, transceiver, etc.)
    
    Args:
        response: XML response string
        component_type: Component type to filter ('lt', 'nt', 'transceiver')
    
    Returns:
        List of OltComponent objects
    """
    components = []
    
    try:
        root = ET.fromstring(response)
        
        # Define namespaces
        namespaces = {
            'hw': 'urn:ietf:params:xml:ns:yang:ietf-hardware',
            'nokia-hwi': 'http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-hardware-identities',
            'bbf-hwt': 'urn:bbf:yang:bbf-hardware-types'
        }
        
        # Find all component elements
        for comp_elem in root.findall('.//hw:component', namespaces):
            # Extract class and check if it matches the component type
            class_elem = comp_elem.find('hw:class', namespaces)
            if class_elem is None or class_elem.text is None:
                continue
            
            class_text = class_elem.text
            
            # Check if class matches the component type
            # Support both nokia-hwi and bbf-hwt namespaces
            match_found = False
            if component_type.lower() == 'transceiver':
                match_found = f"bbf-hwt:{component_type}" in class_text
            else:
                match_found = f"nokia-hwi:{component_type}" in class_text
            
            if not match_found:
                continue
            
            # Extract required fields
            name = comp_elem.findtext('hw:name', '', namespaces)
            parent = comp_elem.findtext('hw:parent', '', namespaces)
            model_name = comp_elem.findtext('hw:model-name', '', namespaces)
            
            # Extract state information
            state_elem = comp_elem.find('hw:state', namespaces)
            if state_elem is not None:
                state_last_changed = state_elem.findtext('hw:state-last-changed', '', namespaces)
                admin_state = state_elem.findtext('hw:admin-state', '', namespaces)
                oper_state = state_elem.findtext('hw:oper-state', '', namespaces)
                standby_state = state_elem.findtext('hw:standby-state', '', namespaces)
            else:
                state_last_changed = admin_state = oper_state = standby_state = ''
            
            # Create ComponentState object
            component_state = ComponentState(
                state_last_changed=state_last_changed,
                admin_state=admin_state,
                oper_state=oper_state,
                standby_state=standby_state
            )
            
            # Create OltComponent object
            component = OltComponent(
                name=name,
                type=component_type.lower(),
                parent=parent,
                model=model_name,
                state=component_state,
            )
            
            components.append(component)
    
    except ET.ParseError as e:
        logger.error(f"Failed to parse XML response: {e}")
    
    return components

def get_onu_list_from_response(response: str):
    return _parse_interface_from_response(response, "v-ani")

def _parse_interface_from_response(response: str, interface_type: str) -> list:
    """
    Parse XML response and extract interfaces of specified type (v-ani, etc.)
    
    Args:
        response: XML response string
        interface_type: Interface type to filter ('v-ani')
    
    Returns:
        List of OltInterface objects
    """
    interfaces = []
    
    try:
        root = ET.fromstring(response)
        
        # Define namespaces
        namespaces = {
            'if-state': 'urn:ietf:params:xml:ns:yang:ietf-interfaces',
            'bbf-xponift': 'urn:bbf:yang:bbf-xpon-if-type',
            'bbf-xponvani': 'urn:bbf:yang:bbf-xponvani',
            'bbf-xpon-onu-types': 'urn:bbf:yang:bbf-xpon-onu-types',
            'nokia-sdan': 'http://www.nokia.com/Fixed-Networks/BBA/yang/nokia-sdan-if-xponvani-aug',
            'bbf-xpon-def': 'urn:bbf:yang:bbf-xpon-defects'
        }
        
        # Find all interface elements
        for if_elem in root.findall('.//if-state:interface', namespaces):
            # Extract type element
            type_elem = if_elem.find('if-state:type', namespaces)
            if type_elem is None or type_elem.text is None:
                continue
            
            # Check if type matches the interface type
            if f"bbf-xponift:{interface_type}" not in type_elem.text:
                continue
            
            # Extract interface basic fields
            name = if_elem.findtext('if-state:name', '', namespaces)
            admin_status = if_elem.findtext('if-state:admin-status', '', namespaces)
            oper_status = if_elem.findtext('if-state:oper-status', '', namespaces)
            last_change = if_elem.findtext('if-state:last-change', '', namespaces)
            if_index = if_elem.findtext('if-state:if-index', '', namespaces)
            lower_layer_if = if_elem.findtext('if-state:lower-layer-if', '', namespaces)
            
            # Extract v-ani specific data
            vani_elem = if_elem.find('bbf-xponvani:v-ani', namespaces)
            onu_id = ''
            management_tcont_alloc_id = ''
            management_gemport_id = ''
            onu_presence_state = ''
            onu_fiber_distance = ''
            detected_serial_number = ''
            detected_registration_id = ''
            
            if vani_elem is not None:
                onu_id = vani_elem.findtext('bbf-xponvani:onu-id', '', namespaces)
                management_tcont_alloc_id = vani_elem.findtext('bbf-xponvani:management-tcont-alloc-id', '', namespaces)
                management_gemport_id = vani_elem.findtext('bbf-xponvani:management-gemport-id', '', namespaces)
                onu_presence_state = vani_elem.findtext('bbf-xponvani:onu-presence-state', '', namespaces)
                
                # Extract onu-present-on-this-olt data
                onu_present_elem = vani_elem.find('bbf-xponvani:onu-present-on-this-olt', namespaces)
                if onu_present_elem is not None:
                    onu_fiber_distance = onu_present_elem.findtext('bbf-xponvani:onu-fiber-distance', '', namespaces)
                    detected_serial_number = onu_present_elem.findtext('bbf-xponvani:detected-serial-number', '', namespaces)
                    detected_registration_id = onu_present_elem.findtext('bbf-xponvani:detected-registration-id', '', namespaces)
            
            # Create OltInterfaceVaniOnuPresentOnThisOlt object
            onu_present_on_this_olt = OltInterfaceVaniOnuPresentOnThisOlt(
                onu_fiber_distance=int(onu_fiber_distance) if onu_fiber_distance else 0,
                detected_serial_number=detected_serial_number,
                detected_registration_id=detected_registration_id
            )
            
            # Create OltInterfaceVani object
            vani = OltInterfaceVani(
                onu_id=int(onu_id) if onu_id else 0,
                management_tcont_alloc_id=int(management_tcont_alloc_id) if management_tcont_alloc_id else 0,
                management_gemport_id=int(management_gemport_id) if management_gemport_id else 0,
                onu_presence_state=onu_presence_state,
                onu_presentt=onu_present_on_this_olt
            )
            
            # Create OltInterface object
            interface = OltInterface(
                name=name,
                type=interface_type,
                admin_state=admin_status,
                oper_state=oper_status,
                last_changed=last_change,
                if_index=if_index,
                lower_layer_if=lower_layer_if,
                vani=vani
            )
            
            interfaces.append(interface)
    
    except ET.ParseError as e:
        logger.error(f"Failed to parse XML response: {e}")
    
    return interfaces

class OltInfoManager():
    def __init__(self, client : NetconfClient):
        self.client = client

    def get_hardware(self):
        return self.client.get_config(RpcGenerator().get_hardware_component_list())

    def get_config_nt_list(self):
        return self.client.get_config(RpcGenerator().get_nt_list())

    def get_config_lt_list(self):
        return self.client.get_config(RpcGenerator().get_lt_list())

    def get_hardware_state(self):
        return self.client.get(RpcGenerator().get_hardware_state_component_list())

    def get_online_nt_list(self):
        return self.client.get(RpcGenerator().get_online_nt_list())

    def get_online_lt_list(self):
        return self.client.get(RpcGenerator().get_online_lt_list())
    
    def _process_single_lt(self, lt_component, use_netconf: bool = False) -> LtInfo:
        """
        Process a single LT component to get its topology information.
        This method is designed to be called concurrently.
        
        Args:
            lt_component: LT component to process
            use_netconf: Whether to create a separate NETCONF connection
            
        Returns:
            LtInfo object with PON and ONU information
        """
        try:
            if use_netconf:
                netconf = Netconf(
                    self.client.device["host"], 
                    get_netconf_port_by_lt_parent(lt_component.parent), 
                    self.client.device["username"], 
                    self.client.device["password"],
                    timeout = 300
                )
                client = NetconfClient(netconf)
            else:
                client = self.client
            
            lt_info = LtInfoManager(client)
            
            # Get PON ports
            out = lt_info.get_pon_port_list()
            pons = get_pon_list_from_response(out)
            
            # Get ONU list
            out = lt_info.get_onu_list()
            onus = get_onu_list_from_response(out)
            
            # Create PON info only if pons data exists
            if pons:
                pon_info = PonInfo.from_component(pons[0], onuinfo=onus)
                poninfo_list = [pon_info]
            else:
                logger.warning(f"No PON data found for LT {lt_component.name}")
                poninfo_list = []
            
            # Create LT info
            lt_info_obj = LtInfo.from_component(lt_component, poninfo=poninfo_list)
            
            # Close connection if we created a new one
            if use_netconf:
                client.close()
            
            logger.info(f"Successfully processed LT: {lt_component.name}")
            return lt_info_obj
            
        except Exception as e:
            logger.error(f"Error processing LT {lt_component.name}: {e}", exc_info=True)
            # Return empty LT info on error
            return LtInfo.from_component(lt_component, poninfo=[])
    
    def get_lt_topology(self, lt_components: list, use_netconf: bool = False) -> list[LtInfo]:
        """
        Get topology information for multiple LT components concurrently.
        
        Args:
            lt_components: List of LT components to process
            use_netconf: Whether to create separate NETCONF connections for each LT
            
        Returns:
            List of LtInfo objects
        """
        if not lt_components:
            return []
        
        lt_info_list = []
        
        # Process LTs concurrently using ThreadPoolExecutor
        with ThreadPoolExecutor(max_workers=min(len(lt_components), 8)) as executor:
            # Submit all tasks
            future_to_lt = {
                executor.submit(self._process_single_lt, lt_comp, use_netconf): lt_comp 
                for lt_comp in lt_components
            }
            
            # Collect results as they complete
            for future in as_completed(future_to_lt):
                lt_component = future_to_lt[future]
                try:
                    lt_info = future.result()
                    lt_info_list.append(lt_info)
                except Exception as e:
                    logger.error(f"Failed to get result for LT {lt_component.name}: {e}")
                    # Add empty LT info on error
                    lt_info_list.append(LtInfo.from_component(lt_component, poninfo=[]))
        
        return lt_info_list

    def get_olt_topology(self, action: str) -> NtInfo:
        logger.info(f" OLT {self.client.device['host']} action: {action} ")
        get_onus = action == "get_topology_with_onus"
        out = self.get_online_nt_list()
        nt = get_nt_list_from_response(out)
        nt_component = nt[0] if nt else None
        
        if not nt_component:
            return NtInfo(None, ltinfo=[])
        
        if not is_df_platform(nt_component.model):
            out = self.get_online_lt_list()
            lt = get_lt_list_from_response(out)
            healthy_lt = filter_healthy_lt_list(lt)
            
            if not healthy_lt or not get_onus:
                return NtInfo.from_component(nt_component, healthy_lt)
            
            lt_info_list = self.get_lt_topology(healthy_lt, use_netconf=True)
            
            return NtInfo.from_component(nt_component, lt_info_list)
        elif get_onus:
            lt_info_list = self.get_lt_topology([nt_component], use_netconf=False)
            
            return NtInfo.from_component(nt_component, lt_info_list)
    

class LtInfoManager():
    def __init__(self, client : NetconfClient):
        self.client = client

    def get_pon_port_list(self):
        return self.client.get(RpcGenerator().get_pon_port_list())

    def get_onu_list(self):
        return self.client.get(RpcGenerator().get_online_onu_list())

def filter_healthy_lt_list(lt_list: list) -> list:
    return [
        lt for lt in lt_list
        if lt.state.admin_state == 'unlocked' and lt.state.oper_state == 'enabled'
    ]