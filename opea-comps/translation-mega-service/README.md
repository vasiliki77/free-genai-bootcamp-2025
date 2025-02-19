# Translation Mega Service

A service for translating English to Ancient Greek.

For detailed steps see the following entry in the journal:

[Translation Mega Service Development](#translation-mega-service-development)
## Example Usage


1. Using Python requests:
````python
import requests

response = requests.post(
    "http://localhost:8000/v1/translate",
    json={"text": "Where is the dog? He needs to eat."}
)
print(response.json())
````

1. Using httpie (more readable than curl):
````bash
http POST localhost:8000/v1/translate \
    text="Where is the dog? He needs to eat."
````

1. Or a cleaner curl format:
````bash
curl localhost:8000/v1/translate \
  -X POST \
  -H "Content-Type: application/json" \
  -d @- << EOF
{
  "text": "Where is the dog? He needs to eat."
}
EOF
````

```bash
 curl -X POST http://localhost:8000/v1/translate   -H "Content-Type: application/json"   -d '{"text": "Where is the dog? He needs to eat."}'
 ```





