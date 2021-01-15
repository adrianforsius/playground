from dataclasses import dataclass


@dataclass
class Mangle:
    __name: str = ""

    def hello(self):
        return self.__name


mangle = Mangle("hello")
print(mangle.hello())

print(mangle._Mangle__name)

print(mangle.__name)