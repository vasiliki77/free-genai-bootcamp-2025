# Project Title
**GenAI Bootcamp Final “SHRED” Presentation**

## 1. Introduction & Motivation


Ancient Greek is more than just an old language; it is a cornerstone of Western literature, philosophy, science, and history. Many foundational texts—from Homer’s epics to Aristotle’s treatises—were originally composed in Ancient Greek, and their ideas continue to influence modern thought, language, and culture. By studying Ancient Greek, learners gain direct access to primary sources, preserving not only the language itself but also the intricate worldview it embodies.

Keeping Ancient Greek alive ensures that future generations can engage with these classic works in their original form, deepening our collective understanding of the roots of democracy, ethics, and the humanities. Additionally, the study of Ancient Greek can sharpen linguistic skills, enhance critical thinking, and foster cross-cultural awareness. As new technologies emerge, harnessing them to make ancient languages more accessible breathes fresh life into texts that might otherwise remain confined to academic circles—helping to bridge the gap between our modern world and the rich heritage of the past.

  **Context**: 
- **Domain / Problem Addressed**:  
  This project focuses on using modern digital tools—such as LLM models—to support learning and preservation of Ancient Greek. While many educational technologies and research efforts center on contemporary languages, resources for ancient or “dead” languages are often limited. This gap puts culturally significant texts and traditions at risk of becoming inaccessible to new generations.

- **Why It’s Relevant / Important**:  
  Ancient Greek has played a pivotal role in shaping Western literature, philosophy, and scientific thought. By making the language more approachable through modern techniques—like AI-driven tutoring, automated translations, or interactive grammar exercises—we help safeguard a key part of our shared intellectual and cultural heritage. Engaging with Ancient Greek texts in their original form enriches our understanding of foundational ideas that still influence society today, whether in ethics, politics, science, or the arts. Beyond preservation, this work highlights how cutting-edge technology can bridge centuries, fostering a dialogue between past and present for the benefit of scholars, students, and anyone with a passion for history and linguistics.

 **Motivation**:
- **Rationale for Addressing the Challenge**:  
  Despite being Greek, I had limited formal exposure to Ancient Greek because of the educational structure in Greece. While basic ancient Greek is introduced in early high school, students later choose from three academic streams: Technology, Scientific, or Humanities. Only the Humanities track continues advanced Ancient Greek studies, so my decision to pursue Scientific precluded further learning in this field. However, I’ve always been motivated by the desire to read and appreciate Ancient Greek texts in their original form.

- **Beneficiaries of the Solution**:  
  This approach targets students following the Humanities track who are required to pass Ancient Greek for academic advancement, as well as anyone—regardless of academic background—who aspires to learn or deepen their understanding of the Greek language.

---


## 2. Scientific/Technological Objective

- **Objective**  
  My main goal is to develop an interactive, user-friendly app that lowers the barrier to learning Ancient Greek. This involves integrating linguistic databases and AI-driven learning resources into a cohesive platform, so learners can practice reading, writing and translation in real-time while receiving immediate feedback.

- **Significance**  
  Achieving this objective advances digital language education by offering a modern, accessible gateway to a traditionally specialized field of study. By consolidating diverse resources into an easy-to-use application, the project helps preserve Ancient Greek and democratize classical studies, making the language more approachable for high school students, independent learners, and anyone eager to explore the foundations of Western literature and thought.

---


## 3. Uncertainty or Knowledge Gap

- **Uncertainty**  
  At the outset, I had limited proficiency in Ancient Greek, making it challenging to evaluate the accuracy of more complex translations. Additionally, I lack a formal background in software development, so I relied on AI tools like Lovable to help build the frontend and used Claude Sonnet extensively throughout this project. While these resources significantly boosted productivity, I initially had minimal experience in leveraging AI solutions—my credentials were limited to an AI Practitioner Certificate and the Gen AI Essentials course taken prior to the bootcamp. On top of this, I spent considerable time researching pretrained models capable of translating and transcribing Ancient Greek, experimenting with various options to determine which would best meet the project’s needs.

- **Why It Was Not Trivial to Solve**  
  Evaluating translations for a language that I do not speak fluently and integrating AI-driven development tools without a strong programming background created a multifaceted challenge. I had to learn and adapt rapidly in multiple domains—ancient languages, software engineering, and AI development workflows—which required extensive trial, error, and iteration.

- **Why It Matters**  
  By working through these complexities, I gained invaluable insights into selecting, deploying, and researching pretrained models for niche language tasks. I tested a range of approaches—running solutions locally, via Docker, and in cloud environments—and learned best practices for harnessing AI within a specialized domain. These findings not only strengthened my technical proficiency but also demonstrated how AI-driven frameworks can make less commonly studied languages more accessible and ensure their continued vitality.

---


## 4. Methodology / Systematic Investigation

1. **Planning & Research**  
   - **Initial Exploration**: Rather than conducting a formal literature review, I focused on practical sources like Hugging Face, searching for models potentially capable of handling Ancient Greek.  
   - **Tool & Framework Selection**: To identify the most effective model, I set up tests in Google Colab and ran detailed prompts to evaluate each candidate’s ability to translate accurately into Ancient Greek.  

2. **Experimental Design / Implementation**  
   - **Prompt Engineering**: I used few-shot prompt engineering to guide the models in generating precise Ancient Greek text, including correct polytonic diacritics rather than modern Greek.  
   - **Technical Stack**: My experiments spanned multiple platforms and libraries, including ChatGPT, Claude, Mistral, MarianMT, Facebook/OPT-125M, VLLM, OpenVINO notebooks (on both WSL and Windows), and the OPEA services.  

3. **Testing & Iteration**  
   - **Evaluation Methods**: I measured translation quality, transcription accuracy, and response latency across different models to determine their suitability for Ancient Greek tasks.  
   - **Continuous Refinement**: Based on performance data and observed issues, I iterated on prompts, model configurations, and system settings to achieve more reliable outputs.  

4. **Documentation & Version Control**  
   - **Project Tracking**: I maintained a weekly journal detailing each day’s progress, tests conducted, and lessons learned.  
   - **Version Control Workflow**: Each new task was tracked in a dedicated branch, ensuring organized records of changes and facilitating easier rollbacks if needed.

---


## 5. Observations, Results & Insights

- **Key Findings**  
  - Through comparative testing, Mistral Instruct 7B emerged as the most accurate model for translating Ancient Greek texts. However, it required a GPU environment to run efficiently, as CPU-based operation proved too slow.  
  - For transcription across multiple languages—including Greek—Coqui-TTS offered broad functionality, making it a useful component for capturing spoken input or generating audio outputs in Ancient Greek.

- **Challenges**  
  - While earlier sections addressed some hurdles—such as limited domain expertise, tool integration, and the complexity of accurately assessing Ancient Greek translations—these issues continued to influence both testing and implementation. Ongoing refinements and additional resources may be necessary to achieve greater reliability and performance across all features.


---

## 6. Conclusion & Next Steps

- **Conclusion**  
  - **Achievement of Initial Objective**: While the original goal was to create a comprehensive app covering grammar, vocabulary, and other language-learning components, resource constraints limited the final product to core translation (English to Ancient Greek), listening, speaking features, and diacritic-based writing practice (the latter not driven by AI). Despite these limitations, the core aim—developing a functional proof of concept—was fulfilled.  
  - **Key Takeaways**: This project demonstrated that with pre-trained models and AI tools, a single individual can undertake tasks traditionally requiring a specialized team (development, DevOps, QA, etc.). In essence, the barriers to building sophisticated applications are dramatically lower when leveraging modern AI frameworks.

- **Future Work**  
  - **Refinement & Expansion**: To offer a richer user experience, additional time and collaboration with an experienced Ancient Greek instructor would be beneficial. This would enable the integration of more robust learning features, including structured grammar lessons, vocabulary drills, and culturally informed content.  
  - **Potential Domains for Exploration**: Beyond Ancient Greek, this methodology can be applied to other niche or “dead” languages, leveraging AI to revitalize under-resourced linguistic communities and potentially broaden interest in cultural preservation.

---

## 7. References


- **References**:
  1. [Mistral Instruct 7B on Hugging Face](https://huggingface.co)  
  2. [Coqui TTS](https://huggingface.co/mistralai/Mistral-7B-Instruct-v0.3)  
  3. [OpenVINO Notebooks](https://github.com/openvinotoolkit/openvino_notebooks)  
  4. [OPEA](https://opea-project.github.io/latest/index.html)
  5. [Perseus Digital Library](https://www.perseus.tufts.edu/hopper/)
  6. [Perseus Digital Library on wikipedia](https://en.wikipedia.org/wiki/Perseus_Digital_Library)

  7. [Lovable](https://lovable.dev/)


---

**End of Presentation**

