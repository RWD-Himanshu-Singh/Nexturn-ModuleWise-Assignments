
from customer_management import Customer

class Transaction(Customer):
    def __init__(self, name, email, phone, book_title, quantity_sold):
        super().__init__(name, email, phone)
        self.book_title = book_title
        self.quantity_sold = quantity_sold


sales = []  # List to store all sales transactions


def sell_book(customer_name, customer_email, customer_phone, book_title, quantity):
    from book_management import books 

    book = next((b for b in books if b.title.lower() == book_title.lower()), None)

    if not book:
        return "Error: Book not found."
    if book.quantity < quantity:
        return f"Error: Only {book.quantity} copies available. Sale cannot be completed."
    
    book.quantity -= quantity
    transaction = Transaction(customer_name, customer_email, customer_phone, book_title, quantity)
    sales.append(transaction)
    return f"Sale successful! Remaining quantity: {book.quantity}"


def view_sales():
    if not sales:
        return "No sales records available."
    return "\n".join(
        f"Customer: {sale.name}, Book: {sale.book_title}, Quantity: {sale.quantity_sold}"
        for sale in sales
    )
