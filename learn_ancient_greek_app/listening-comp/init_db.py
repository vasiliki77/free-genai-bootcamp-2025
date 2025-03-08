import sqlite3
import xml.etree.ElementTree as ET
import os

# Get the absolute path to the database
BASE_DIR = os.path.dirname(os.path.abspath(__file__))
DB_PATH = os.path.join(BASE_DIR, 'backend', 'listening-learning.db')

def group_lines(lines, lines_per_passage=4):
    """Group lines into passages of specified size with overlap"""
    passages = []
    for i in range(0, len(lines) - lines_per_passage + 1, 2):  # Step by 2 for overlap
        passage = ' '.join(lines[i:i + lines_per_passage])
        if passage.strip():
            passages.append(passage)
    return passages

def init_database():
    # Create database directory if it doesn't exist
    os.makedirs(os.path.dirname(DB_PATH), exist_ok=True)
    
    # Check if XML file exists
    xml_path = os.path.join(BASE_DIR, 'ancient_greek_text.xml')
    if not os.path.exists(xml_path):
        print(f"❌ Error: XML file not found at {xml_path}")
        return
    
    # Create database and table
    print(f"Creating database at: {DB_PATH}")
    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()
    
    # Drop table if exists and create new one
    cursor.execute('DROP TABLE IF EXISTS greek_texts')
    cursor.execute('''
        CREATE TABLE greek_texts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            passage TEXT NOT NULL
        )
    ''')
    
    # Parse XML and insert texts
    try:
        print("Parsing XML file...")
        tree = ET.parse(xml_path)
        root = tree.getroot()
        
        # Find all line elements and extract their text
        lines = []
        for line in root.findall('.//l'):
            if line is not None and line.text:
                text = line.text.strip()
                if text:
                    lines.append(text)
                    print(f"Found line: {text[:50]}...")
        
        if not lines:
            print("❌ No valid lines found in XML file")
            return
        
        # Group lines into passages
        passages = group_lines(lines)
        if not passages:
            print("❌ No valid passages could be created from lines")
            return
        
        # Insert all passages
        print(f"\nCreating {len(passages)} passages from {len(lines)} lines...")
        cursor.executemany('INSERT INTO greek_texts (passage) VALUES (?)', 
                         [(p,) for p in passages])
        
        # Commit changes and close
        conn.commit()
        print(f"✅ Successfully initialized database with {len(passages)} passages")
        print(f"Sample passage: {passages[0][:100]}...")
        
    except ET.ParseError as e:
        print(f"❌ Error parsing XML file: {str(e)}")
        conn.rollback()
    except Exception as e:
        print(f"❌ Error initializing database: {str(e)}")
        conn.rollback()
    finally:
        conn.close()

if __name__ == '__main__':
    print("🚀 Initializing database...")
    init_database()
