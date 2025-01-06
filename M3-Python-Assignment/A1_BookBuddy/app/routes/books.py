from flask import Blueprint, request, jsonify
from app.models.database import get_db_connection
from app.utils.validators import validate_book_data

books_bp = Blueprint("books", __name__)

@books_bp.route("/")
def index():
    return jsonify({"message": "Welcome to the Bookstore!"})


@books_bp.route("/books", methods=["POST"])
def add_book():
    data = request.get_json()
    is_valid, error = validate_book_data(data)
    if not is_valid:
        return jsonify({"error": "Invalid data", "message": error}), 400

    conn = get_db_connection()
    cursor = conn.cursor()
    cursor.execute("""
        INSERT INTO books (title, author, published_year, genre)
        VALUES (?, ?, ?, ?)
    """, (data["title"], data["author"], data["published_year"], data["genre"]))
    conn.commit()
    book_id = cursor.lastrowid
    conn.close()

    return jsonify({"message": "Book added successfully", "book_id": book_id}), 201

@books_bp.route("/books", methods=["GET"])
def get_books():
    conn = get_db_connection()
    books = conn.execute("SELECT * FROM books").fetchall()
    conn.close()

    return jsonify([dict(book) for book in books])

@books_bp.route("/books/<int:book_id>", methods=["GET"])
def get_book(book_id):
    conn = get_db_connection()
    book = conn.execute("SELECT * FROM books WHERE id = ?", (book_id,)).fetchone()
    conn.close()

    if book is None:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

    return jsonify(dict(book))

@books_bp.route("/books/<int:book_id>", methods=["PUT"])
def update_book(book_id):
    data = request.get_json()
    is_valid, error = validate_book_data(data)
    if not is_valid:
        return jsonify({"error": "Invalid data", "message": error}), 400

    conn = get_db_connection()
    cursor = conn.cursor()
    cursor.execute("""
        UPDATE books
        SET title = ?, author = ?, published_year = ?, genre = ?
        WHERE id = ?
    """, (data["title"], data["author"], data["published_year"], data["genre"], book_id))
    conn.commit()

    if cursor.rowcount == 0:
        conn.close()
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

    conn.close()
    return jsonify({"message": "Book updated successfully"})

@books_bp.route("/books/<int:book_id>", methods=["DELETE"])
def delete_book(book_id):
    conn = get_db_connection()
    cursor = conn.cursor()
    cursor.execute("DELETE FROM books WHERE id = ?", (book_id,))
    conn.commit()

    if cursor.rowcount == 0:
        conn.close()
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

    conn.close()
    return jsonify({"message": "Book deleted successfully"})
