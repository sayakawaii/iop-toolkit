from concurrent.futures import ThreadPoolExecutor
import logging

logger = logging.getLogger(__name__)

class BaseHandler:
    def __init__(self, max_workers=4):
        self.executor = ThreadPoolExecutor(max_workers=max_workers)

    def submit(self, key, value):
        self.executor.submit(self.handle, key, value)

    def handle(self, key, value):
        raise NotImplementedError

    def shutdown(self):
        self.executor.shutdown(wait=True)
