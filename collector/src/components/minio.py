from minio import Minio
from minio.error import S3Error
import os
import socketserver
import logging

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s][%(name)s] %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S"
)
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)

# MinIO connection is fully env-driven so the same image runs on the
# server (238) and on a local WSL host without code changes.
minio_endpoint = os.environ.get("MINIO_ENDPOINT", "localhost:9000")
minio_access_key = os.environ.get("MINIO_ACCESS_KEY", "eonu")
minio_secret_key = os.environ.get("MINIO_SECRET_KEY", "eonu#1234")
minio_secure = os.environ.get("MINIO_SECURE", "false").lower() == "true"

client = Minio(
    minio_endpoint,
    access_key=minio_access_key,
    secret_key=minio_secret_key,
    secure=minio_secure
)

class MinIOLogHandler(socketserver.BaseRequestHandler):
    def __init__(self, miniotarget : str = None):
        self.miniotarget = miniotarget
    def upload_file(self, bucket_name: str, object_name: str, file_path: str):
        try:
            # 确保 bucket 存在
            if not client.bucket_exists(bucket_name):
                client.make_bucket(bucket_name)
            
            # 上传文件
            client.fput_object(bucket_name, object_name, file_path)
            logger.info(f"File '{file_path}' uploaded to bucket '{bucket_name}' as '{object_name}'.")

        except S3Error as e:
            logger.info(f"Error occurred: {e}")