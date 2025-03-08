import requests

# URL for Book 1 of the Iliad in XML format
url = "http://www.perseus.tufts.edu/hopper/xmlchunk?doc=Perseus:text:1999.01.0133:book=1"

# Send GET request to fetch the XML content
response = requests.get(url)

# Save the XML file locally
if response.status_code == 200:
    with open("ancient_greek_text.xml", "wb") as f:
        f.write(response.content)
    print("✅ Ancient Greek XML file downloaded successfully!")
else:
    print(f"❌ Failed to download. Status code: {response.status_code}")