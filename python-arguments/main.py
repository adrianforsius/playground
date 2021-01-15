from dataclasses import dataclass


@dataclass
class MasterTask:
    test: str


task = MasterTask("my task config")
task.x = "y"
new_task = MasterTask(**task.__dict__)