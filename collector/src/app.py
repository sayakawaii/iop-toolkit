from logic.fsm import FiniteStateMachine
import logging

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s][%(name)s] %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S"
)
logger = logging.getLogger("main")
logger.setLevel(logging.INFO)

def main():
    fsm = FiniteStateMachine()
    fsm.run()

if __name__ == "__main__":
    main()
