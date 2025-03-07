# free-genai-bootcamp-2025

## Table of Contents
- [Table of Contents](#table-of-contents)
- [Preweek Journal](journal/week0.md)
- [Week 1 Journal](journal/week1.md)
- [Week 2 Journal](journal/week2.md)
- [Week 3 Journal](journal/week3.md)
- [Week 4 Journal](journal/week4.md)


## Ancient Greek Translation API with Mistral 7B Instruct v0.3 on Lightning.ai

This guide provides step-by-step instructions to reproduce and deploy the Mistral 7B Instruct v0.3 translation model using Lightning.ai and FastAPI. The model translates English sentences into Ancient Greek.


### Prerequisites

- Python 3.10
- [Lightning.ai](https://lightning.ai/) account
- Hugging Face account and [API token](https://huggingface.co/settings/tokens)

### 1. Clone the Repository into a Lightning studio

```bash
git clone <repository-url>
cd mistral_translation_app
```

### 2. Install Dependencies

```bash
pip install -r requirements.txt
```

### 3. Set Hugging Face Token

Set your Hugging Face token as an environment variable:

```bash
export HF_TOKEN='your_huggingface_token_here'
```

### 4. Set API key


```bash
export SECRET_API_KEY='a_secret_key'
```

### 4. Run Locally (Testing)

To run your API locally and test:

```bash
python app.py
```


Test translation example:

```bash
⚡ ~ curl -X POST "http://localhost:8080/translate?sentence=Wildboars%20are%20at%20the%20door,%20did%20you%20leave%20the%20garbage%20out?" -H "api-key: a_secret_key"
```



### Expected JSON Response:

```json
{"original":"Wildboars are at the door, did you leave the garbage out?","translation":"and","full_response":"\nTranslate the following English sentence accurately into Classical Ancient Greek (Attic dialect). \nUse correct grammar, vocabulary, and polytonic diacritics. \nProvide ONLY the Ancient Greek translation enclosed by <START> and <END> tags, nothing else.\n\nExamples:\nEnglish: Wisdom is virtue. → <START>Σοφία ἐστὶν ἀρετή.<END>\nEnglish: Life is short. → <START>Ὁ βίος βραχύς ἐστιν.<END>\nEnglish: Know thyself. → <START>Γνῶθι σεαυτόν.<END>\nEnglish: Hello world. → <START>Χαῖρε, ὦ κόσμε!<END>\nEnglish: I love philosophy. → <START>Φιλοσοφίαν φιλῶ.<END>\n\nNow translate accurately:\nEnglish: Wildboars are at the door, did you leave the garbage out? → <START>Ἄιμνοι εἰσὶν τῇ θύρᾳ, ἦλθες λιπὼν τὸ σκόροπον;<END>"}
```
### Translation Quality and Limitations

The app uses Mistral-7B-Instruct-v0.3, a general-purpose instruction-following model. Prompt engineering with explicit examples significantly improves translation accuracy into Ancient Greek. However, because the model is not fine-tuned specifically for Ancient Greek translation, complex sentences or unusual vocabulary might lead to occasional inaccuracies.

**Recommended next steps:**  
- Fine-tune a model specifically on a dataset of English-to-Ancient Greek translations for enhanced accuracy.
- Extend and diversify prompt examples to improve generalization further.

The current setup demonstrates clearly how prompt engineering enhances performance and provides a stable foundation suitable for the bootcamp evaluation.


## Deploy language portal on your local machine

