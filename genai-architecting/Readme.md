# Conceptual Diagram: Language Learning with GenAI

## Functional Requirements
- Enable students to access study activities through a language learning portal.
- Utilize a Retrieval-Augmented Generation (RAG) model for sentence construction.
- Implement guardrails for input and output validation.
- The company wants to invest in owning their infrastructure because there is a concern about the privacy of user data.
- The company will choose on-prem server as the cost of managed services for GenAI might greatly rise in the future.
- They want to invest an AI PC where they can afford to spend 10-15K. They have 300 active students, and students are located within the city of Athens.

## Assumptions
- The system assumes that users have authenticated access via the language portal.
- The RAG model is pre-trained and optimized for sentence construction.
- The selected open-source LLMs will be powerful enough to run on hardware within a 10-15K investment.
- We will connect a single server in our office to the internet, assuming sufficient bandwidth to support 300 students.

## Data Strategy
- Store core vocabulary words in a structured database.
- Use a vector database for optimized retrieval of relevant language constructs.
- Ensure data security and privacy compliance for student interactions.
- We will procure and store licensed content in our database for controlled access to address concerns regarding copyrighted materials.

## Key Considerations
- **Scalability:** The system should support an increasing number of students.
- **Integration:** The study activities should seamlessly integrate with the language portal.
- **Performance:** Optimize response times for AI-generated study activities.
- **Security:** Implement guardrails to ensure appropriate AI-generated content.
- **Model Selection:** We're considering using IBM Granite because it is a truly open-source model with traceable training data, helping to avoid copyright issues while providing transparency into model operations. [IBM Granite on Hugging Face](https://huggingface.co/ibm-granite).
