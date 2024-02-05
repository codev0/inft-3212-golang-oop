# Define the superclass
class Animal:
    def __init__(self, name):
        self.name = name

    def speak(self):
        raise NotImplementedError("Subclass must implement abstract method")

# Define the subclass


class Dog(Animal):
    def speak(self):
        return f"{self.name} says Woof!"


# Create an instance of Dog
dog = Dog("Buddy")
print(dog.speak())
