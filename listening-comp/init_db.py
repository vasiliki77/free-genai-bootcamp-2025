import sqlite3

# Connect to the SQLite database (creates if it doesn't exist)
conn = sqlite3.connect("listening-learning.db")
cursor = conn.cursor()

# Create the table
cursor.execute("""
CREATE TABLE IF NOT EXISTS greek_texts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    source TEXT,
    passage TEXT,
    translation TEXT,
    audio_path TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
""")

# Commit and close connection
conn.commit()
conn.close()

print("✅ Database schema created successfully!")
