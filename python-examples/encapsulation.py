class BankAccount:
    def __init__(self, account_number, balance=0):
        self.__account_number = account_number  # private attribute
        self.__balance = balance  # private attribute

    # Getter method for balance
    def get_balance(self):
        return self.__balance

    # Setter method for balance
    def set_balance(self, balance):
        if balance >= 0:
            self.__balance = balance
        else:
            print("Balance cannot be negative.")

    # Method to deposit money
    def deposit(self, amount):
        if amount > 0:
            self.__balance += amount
            print(f"Amount {amount} deposited.")
        else:
            print("Deposit amount must be positive.")

    # Method to withdraw money
    def withdraw(self, amount):
        if 0 < amount <= self.__balance:
            self.__balance -= amount
            print(f"Amount {amount} withdrawn.")
        else:
            print("Insufficient balance or invalid withdrawal amount.")


# Create a BankAccount object
account = BankAccount("123456789", 1000)

# Access and modify balance using public methods
print(account.get_balance())  # 1000
account.deposit(500)
print(account.get_balance())  # 1500
account.withdraw(200)
print(account.get_balance())  # 1300

# Attempt to access private attribute directly (will raise an AttributeError)
# print(account.__balance)  # This will raise an error
