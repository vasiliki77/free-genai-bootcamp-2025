## Table of Contents

- [Table of Contents](#table-of-contents)
- [Listening Learning App](#listening-learning-app)

> 2025-02-24
## Listening Learning App

1. Extracted Ancient Greek Text from [Perseus Digital Library](https://en.wikipedia.org/wiki/Perseus_Digital_Library) because there was no ancient transcription in YouTube.
- Downloaded the XML file of Hommer Iliad book 1.
- Converted the XML URL so it could be fetched programmatically.
- Python script to download XML:

    ```python
    import requests

    # XML URL from Perseus
    url = "https://www.perseus.tufts.edu/hopper/xmlchunk?doc=Perseus:text:1999.01.0135:book=1:card=1"

    response = requests.get(url)

    if response.status_code == 200:
        with open("listening-learning/ancient_greek_text.xml", "wb") as f:
            f.write(response.content)
        print("✅ Ancient Greek XML file downloaded successfully!")
    else:
        print("❌ Failed to download.")
    ```

- Command to run the script:
    ```bash
    python listening-learning/xml-download.py
    ```

---

2. Created a New SQLite Database for Storage
- Created a separate SQLite database inside `listening-learning/` instead of modifying `words.test.db`.
- Created SQLite database (`listening-learning.db`):
    ```bash
    touch listening-learning/listening-learning.db
    ```
- Created a table to store Ancient Greek text:

    ```python
    import sqlite3

    conn = sqlite3.connect("listening-learning/listening-learning.db")
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

    print("✅ Database schema created!")
    ```
- Command to run the script:
    ```bash
    python listening-learning/init_db.py
    ```

---

 3. Parsed and Stored Ancient Greek Text from XML
- Extracted Ancient Greek text from the XML file.
- Stored it in `listening-learning.db`.

    ```python
    import sqlite3
    import xml.etree.ElementTree as ET

    DB_PATH = "listening-learning/listening-learning.db"
    XML_PATH = "listening-learning/ancient_greek_text.xml"

    def extract_greek_text(xml_file):
        tree = ET.parse(xml_file)
        root = tree.getroot()

        greek_passages = []
        for elem in root.iter():
            if elem.text:
                text = elem.text.strip()
                if len(text) > 5:
                    greek_passages.append(text)

        return "\n".join(greek_passages)

    def store_in_database(source, greek_text):
        conn = sqlite3.connect(DB_PATH)
        cursor = conn.cursor()

        cursor.execute("INSERT INTO greek_texts (source, passage) VALUES (?, ?)", (source, greek_text))

        conn.commit()
        conn.close()
        print("✅ Ancient Greek text inserted successfully!")

    greek_text = extract_greek_text(XML_PATH)
    store_in_database("Perseus Library", greek_text)
    ```
- Command to run the script:
    ```bash
    python listening-learning/parse_and_store.py
    ```

---

 4. Using [Coqui.ai](https://github.com/coqui-ai/TTS) to generate audio for the Ancient Greek text.

- Downloaded the docker image 
```
docker run --rm -it -p 5002:5002 --entrypoint /bin/bash ghcr.io/coqui-ai/tts-cpu
python3 TTS/server/server.py --list_models #To get the list of available models
python3 TTS/server/server.py --model_name tts_models/el/cv/vits # To start a greek server
```


- Used this simple Python script to send text and get audio:
```python
import requests

# The Greek phrase to transcribe
text = "ἐξ οὗ δὴ τὰ πρῶτα διαστήτην ἐρίσαντε"

# Send POST request to the TTS server
response = requests.post(
    "http://localhost:5002/api/tts",
    data={
        "text": text,
    }
)

# Save the audio if successful
if response.status_code == 200:
    with open("output.wav", "wb") as f:
        f.write(response.content)
    print("✅ Audio saved to output.wav")
```

- Audio was successfully generated and saved to `output.wav`. I was able to play the audio using any media player.
Although not perfect, the pronunciation is good enough for now.


- Then created a script to randomly select and [transcribe words from the database](../listening-comp/transcribe_random_sequence.py) which also worked. 
```
python transcribe_random_sequence.py 
🔎 Getting random sequence from database...

📜 Selected text:
οὔ τις ὁρᾶτο: θάμβησεν δ᾽ Ἀχιλεύς, μετὰ

🔊 Converting to speech...
✅ Audio saved as random_sequence.wav
```

