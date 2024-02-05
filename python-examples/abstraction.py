from abc import ABC, abstractmethod


class Vehicle(ABC):
    def __init__(self, name):
        self.name = name

    @abstractmethod
    def mode_of_transport(self):
        pass

    @abstractmethod
    def fuel_source(self):
        pass


class ElectricCar(Vehicle):
    def mode_of_transport(self):
        return "land"

    def fuel_source(self):
        return "electricity"


class Sailboat(Vehicle):
    def mode_of_transport(self):
        return "water"

    def fuel_source(self):
        return "wind"


# Instantiate objects of the subclasses
tesla = ElectricCar("Tesla Model S")
sailboat = Sailboat("Beneteau Oceanis")

# Using the abstract methods
print(f"{tesla.name}: Mode of transport - {tesla.mode_of_transport()
                                           }, Fuel source - {tesla.fuel_source()}")
print(f"{sailboat.name}: Mode of transport - {sailboat.mode_of_transport()
                                              }, Fuel source - {sailboat.fuel_source()}")

# Attempting to instantiate an object of the abstract class will raise an error
# vehicle = Vehicle("Generic Vehicle")  # This will raise an error
