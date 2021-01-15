from dataclasses import dataclass


@dataclass
class Mangle:
    __name: str = ""


mangle = Mangle()
print(mangle.__name)
