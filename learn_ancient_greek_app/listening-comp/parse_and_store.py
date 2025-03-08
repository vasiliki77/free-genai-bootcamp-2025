import sqlite3
import xml.etree.ElementTree as ET

# Database file path
DB_PATH = "listening-learning.db"
# XML file path
XML_PATH = "ancient_greek_text.xml"

def extract_greek_text(xml_file):
    """Extracts Ancient Greek text from the XML file."""
    tree = ET.parse(xml_file)
    root = tree.getroot()
    
    greek_passages = []
    
    # Iterate over XML nodes to find text content
    for elem in root.iter():
        if elem.text:
            text = elem.text.strip()
            # Ensure text is in Greek (skip metadata)
            if len(text) > 5:  # Avoid short, irrelevant elements
                greek_passages.append(text)

    return "\n".join(greek_passages)

def store_in_database(source, greek_text):
    """Stores extracted Greek text into SQLite."""
    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()

    # Ensure table exists
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

    # Insert extracted text into database
    cursor.execute("INSERT INTO greek_texts (source, passage) VALUES (?, ?)", (source, greek_text))

    conn.commit()
    conn.close()
    print("✅ Ancient Greek text inserted successfully!")

# Run the process
greek_text = extract_greek_text(XML_PATH)
store_in_database("Perseus Library", greek_text)
