import sqlite3

def init_db():
    conn = sqlite3.connect("listening-learning.db")
    cursor = conn.cursor()
    
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
    
    conn.commit()
    conn.close()
    print("✅ Database initialized")

if __name__ == "__main__":
    init_db() 