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
else:
    print("❌ Error:", response.text)
    print("Status code:", response.status_code)
    print("Response headers:", response.headers)
    print("Request data sent:", text) 