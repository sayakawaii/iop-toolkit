import os
import shutil
from datetime import datetime
import logging

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s][%(name)s] %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S"
)
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)

class Remotelog:
    def __init__(self, jump_ip : str|None, target_ip : str, localpath : str):
        self.jump_ip = jump_ip
        self.target_ip = target_ip
        self.localpath = localpath
        self.cleanup_old_logs()
    def new(self, appname_list: list[str]):
        object_list = []
        file_path_list = []
        for appname in appname_list:
            logger.info(f"{appname} remotelog initialized.")
            date = datetime.now().strftime("%Y-%m-%d")
            base_root = os.path.abspath(self.localpath)
            if self.jump_ip:
                object_name = os.path.join(date, self.jump_ip, self.target_ip, appname + ".log")
                localpath = os.path.join(base_root, date, self.jump_ip, self.target_ip)
            else:
                localpath = os.path.join(base_root, date, self.target_ip)
                object_name = os.path.join(date, self.target_ip, appname + ".log")
            os.makedirs(localpath, exist_ok=True)
            file_path = os.path.join(localpath, appname + ".log")
            # Clear file contents if it exists
            if os.path.exists(file_path):
                open(file_path, 'w').close()
            object_list.append(object_name)
            file_path_list.append(file_path)
        return object_list, file_path_list
    
    def cleanup_old_logs(self):
        """Delete all directories that are not from the current date."""
        current_date = datetime.now().strftime("%Y-%m-%d")
        base_root = os.path.abspath(self.localpath)
        
        if not os.path.exists(base_root):
            logger.warning(f"Base log directory does not exist: {base_root}")
            return
        
        for entry in os.listdir(base_root):
            entry_path = os.path.join(base_root, entry)
            if os.path.isdir(entry_path) and entry != current_date:
                try:
                    shutil.rmtree(entry_path)
                    logger.info(f"Deleted old log directory: {entry}")
                except Exception as e:
                    logger.error(f"Failed to delete directory {entry}: {e}")