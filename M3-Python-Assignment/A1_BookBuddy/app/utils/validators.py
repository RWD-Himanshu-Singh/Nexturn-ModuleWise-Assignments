def validate_book_data(data):
    required_fields = {"title", "author", "published_year", "genre"}
    genres = {"Fiction", "Non-Fiction", "Mystery", "Sci-Fi"}

    if not all(field in data for field in required_fields):
        return False, "Missing required fields."

    if not isinstance(data["published_year"], int) or data["published_year"] <= 0:
        return False, "Invalid published_year."

    if data["genre"] not in genres:
        return False, f"Invalid genre. Allowed genres: {', '.join(genres)}"

    return True, None
