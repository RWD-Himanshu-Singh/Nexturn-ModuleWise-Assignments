import sqlite3

DATABASE = "books.db"

def get_db_connection():
    conn = sqlite3.connect(DATABASE)
    conn.row_factory = sqlite3.Row
    return conn

def init_db():
    # Initialize the database with the books table
    connection = sqlite3.connect(DATABASE)
    cursor = connection.cursor()
    cursor.execute("""
        CREATE TABLE IF NOT EXISTS books (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT NOT NULL,
            author TEXT NOT NULL,
            published_year INTEGER NOT NULL,
            genre TEXT NOT NULL
        )
    """)
    sample_books = [
        ("The Great Gatsby", "F. Scott Fitzgerald", 1925, "Fiction"),
        ("To Kill a Mockingbird", "Harper Lee", 1960, "Fiction"),
        ("1984", "George Orwell", 1949, "Sci-Fi"),
        ("The Catcher in the Rye", "J.D. Salinger", 1951, "Fiction"),
        ("Sapiens: A Brief History of Humankind", "Yuval Noah Harari", 2011, "Non-Fiction"),
    ]

    cursor.executemany("""
        INSERT INTO books (title, author, published_year, genre)
        VALUES (?, ?, ?, ?)
    """, sample_books)
    connection.commit()
    connection.close()
