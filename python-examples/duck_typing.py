class Duck:
    def quack(self):
        print("Quack, quack!")


class Frog:
    def quack(self):
        print("The person imitates a duck's quack.")


def make_it_quack(duck):
    duck.quack()


duck = Duck()
frog = Frog()

make_it_quack(duck)   # Output: Quack, quack!
make_it_quack(frog)  # Output: The person imitates a duck's quack.
