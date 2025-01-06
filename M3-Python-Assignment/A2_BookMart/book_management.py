
class Book:
    def __init__(self, title, author, price, quantity):
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity

    def display_details(self):
        return f"Title: {self.title}, Author: {self.author}, Price: {self.price}, Quantity: {self.quantity}"


books = []  # List to store all books


def add_book(title, author, price, quantity):
    try:
        if price <= 0 or quantity <= 0:
            raise ValueError("Price and Quantity must be positive numbers.")
        book = Book(title, author, price, quantity)
        books.append(book)
        return "Book added successfully!"
    except ValueError as e:
        return f"Error: {e}"


def view_books():
    if not books:
        return "No books available."
    return "\n".join(book.display_details() for book in books)


def search_book(keyword):
    results = [book.display_details() for book in books if keyword.lower() in book.title.lower() or keyword.lower() in book.author.lower()]
    if not results:
        return "No books found."
    return "\n".join(results)
