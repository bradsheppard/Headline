from abc import ABC, abstractmethod
from typing import List


class Producer(ABC):

    @abstractmethod
    def update(self, interests: List[str], user_id: int) -> None:
        pass
