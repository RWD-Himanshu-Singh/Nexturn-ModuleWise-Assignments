from flask import Flask
from app.models.database import init_db
from app.routes.books import books_bp

def create_app():
    app = Flask(__name__)

  
    init_db()

  
    app.register_blueprint(books_bp)

    return app
