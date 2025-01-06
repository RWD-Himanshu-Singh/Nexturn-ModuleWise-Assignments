
class Customer:
    def __init__(self, name, email, phone):
        self.name = name
        self.email = email
        self.phone = phone

    def display_details(self):
        return f"Name: {self.name}, Email: {self.email}, Phone: {self.phone}"


customers = []  # List to store all customers


def add_customer(name, email, phone):
    if not name or not email or not phone:
        return "Error: All customer details must be provided."
    customer = Customer(name, email, phone)
    customers.append(customer)
    return "Customer added successfully!"


def view_customers():
    if not customers:
        return "No customers available."
    return "\n".join(customer.display_details() for customer in customers)
