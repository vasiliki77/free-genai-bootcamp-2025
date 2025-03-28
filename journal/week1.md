## Table of Contents

- [Table of Contents](#table-of-contents)
- [Frontend and Backend Technical Specs](#frontend-and-backend-technical-specs)
- [Resolving Backend Startup Issues and API Response Validation](#resolving-backend-startup-issues-and-api-response-validation)
- [Exploring Testing Strategies: Go Unit Tests vs. Ruby/RSpec Integration Tests](#exploring-testing-strategies-go-unit-tests-vs-rubyrspec-integration-tests)
- [Summary of Work on branch reimplementing\_backend](#summary-of-work-on-branch-reimplementing_backend)
- [Summary of Work on endpoints using test database](#summary-of-work-on-endpoints-using-test-database)
- [Technical Specs Analysis \& Frontend Implementation](#technical-specs-analysis--frontend-implementation)
- [Installing openvino on my Intel AI PC](#installing-openvino-on-my-intel-ai-pc)
- [LLM Deployment Session Summary](#llm-deployment-session-summary)
- [Exploration of different models](#exploration-of-different-models)
- [Translation Mega Service Development](#translation-mega-service-development)


> 2025-02-11

## Frontend and Backend Technical Specs

Today I created comprehensive technical specification documents for the language learning portal project:
1. Frontend Technical Specs
- Defined 10 main pages including Dashboard, Study Activities, Words, Groups, and Settings
- Detailed the purpose, components, and required API endpoints for each page
- Outlined key features like study progress tracking, activity launching, and word management
2. Backend Technical Specs
- Established core business goals for a language learning portal
- Defined technical stack: Go backend with SQLite3 database using Gin framework
- Designed database schema with 6 main tables:
    - words
    - groups
    - words_groups
    - study_sessions
    - study_activities
    - word_review_items
- Specified 15+ API endpoints to support the frontend functionality

The specifications create a foundation for a single-user learning portal that will serve as both a vocabulary inventory system and a learning record store, while providing a unified launch point for various learning activities.

> 2025-02-12

## Resolving Backend Startup Issues and API Response Validation

Today the Main Goal was to get the Go backend server (`lang-portal/backend_go`) to start successfully and respond to basic API requests.

I chose to use Cursor for this. Even when I ran out of credits for some models, I was still able to use gemini model. 

**Key Issues I Faced and Resolved:**

1.  **`undefined: service` Errors:**
    *   **Problem:** When running `mage dev`, I encountered errors like `internal/handlers/dashboard.go:12:20: undefined: service`. This indicated that the handlers were unable to find the `internal/service` package.
    *   **Root Cause:**  Initially, I suspected import path issues and potentially Go modules not being fully active.
    *   **Troubleshooting Steps:**
        *   I verified import paths in all handler files.
        *   Checked the `go.mod` file to ensure the module path was correct.
        *   Confirmed the directory structure was as expected (`internal/service` directory present).
        *   Cleaned the Go build cache multiple times.
        *   Explicitly set the `GO111MODULE=on` environment variable, both in the terminal and directly within the `magefile.go` to force Go modules to be active.
        *   I even created a minimal "Hello World" service and handler to isolate the issue.
    *   **Resolution:** Initially, when Composer was creating the files, the folder created was `backeng_go/internal/services` which I renamed manually to `service` but I failed to update it everywhere. So the files were referencing `package services` and not `package service`

2.  **`sql: unknown driver "sqlite3" (forgotten import?)` Error:**
    *   **Problem:** After resolving the package name conflict, I encountered the error `Failed to initialize database: sql: unknown driver "sqlite3" (forgotten import?)`.
    *   **Root Cause:** This error meant that the `sqlite3` database driver was not being correctly loaded at runtime. This is almost always due to a missing or incorrect import in `cmd/server/main.go`.
    *   **Solution:** I identified that the `_ "github.com/mattn/go-sqlite3"` import was missing from `cmd/server/main.go`.  Adding this underscore import (for side effects) correctly registered the `sqlite3` driver.

**Current Status:**

*   **Server Starts Successfully:**  After adding the `sqlite3` import, your backend server now starts without any of the previous errors when I run `mage dev`.
*   **API Endpoints Respond:**  Basic API endpoints like `/api/words`, `/api/dashboard/quick-stats`, and `/api/study_activities` are now responding with `200 OK` status codes and JSON responses (though currently empty or with placeholder data).
*   **Backend Functionality Confirmed (Basic Level):**  I have confirmed that the fundamental server setup, routing, and handler structure are working correctly.

**Next Steps:**

*   **Implement Service Logic:** The next major step is to implement the actual business logic within your services (e.g., in `internal/service/words.go`, `dashboard.go`, etc.). This involves:
    *   Connecting to the SQLite database using the `models` package.
    *   Writing SQL queries to fetch data from the database.
    *   Returning the fetched data from the service functions.
*   **Update Handlers to Use Services:** Modify handlers to call the service functions to get data and then format that data into JSON responses to send back to the client.

> 2025-02-13

## Exploring Testing Strategies: Go Unit Tests vs. Ruby/RSpec Integration Tests

Today I started working on creating tests for the code. Initially I created tests in Go. Afterwards I realised that Andrew chose to use Ruby for testing. 

Unsure about which method to follow, I asked cursor chat and it gave me the following interesting analogy:
> 
> **Analogy:**
> 
> Think of it like testing a car:
> 
> * **Go `*_test.go` files:**  Are like testing individual car components (engine, brakes, steering) in the factory to make sure each part works correctly in isolation and when assembled together.
> * 
> * **Ruby/RSpec tests:** Are like test driving the assembled car on a track to see if it performs as expected in real-world driving conditions (acceleration, braking, handling, etc.).
> * 
> You wouldn't remove the factory component tests just because you are also doing test drives! Both are important for ensuring the car is reliable.
> 
> **In summary:**
> 
> Keep your `*_test.go` files! They are valuable for testing the internal logic of your Go backend. Use Ruby/RSpec tests **in addition** to your Go tests to provide endpoint/system-level testing and ensure your API behaves as expected from an external client's perspective.  This combination will give you a more robust and well-tested application.
> 
> If you were to *only* choose one type of testing (which is not recommended), for a Go backend project, it would generally be more efficient and maintainable to primarily rely on Go's testing tools for most of your testing needs, and use Ruby/RSpec (or similar tools) for specific higher-level end-to-end or API contract verification tests if needed. But ideally, use both for a more complete testing strategy.

> 2025-02-14

## Summary of Work on branch reimplementing_backend

1. **Fixed Database Migration Issues**
- Simplified the SQLite foreign key constraints in the schema
- Disabled GORM's auto-migration to prevent conflicts with our manual schema
- Fixed issues with the words_groups join table

2. **Implemented Word Response Format**
- Created proper JSON response structures for word endpoints
- Added `WordWithStats` struct for list responses
- Added `WordResponse` struct for single word details
- Implemented custom JSON marshaling to match API spec

3. **Fixed Data Formatting**
- Moved stats (correct_count, wrong_count) to top level of response
- Ensured Parts field always includes all verb forms (present, future, aorist, perfect)
- Removed omitempty tags to ensure all fields are present in JSON
- Initialized empty strings for verb parts instead of null values

4. **Code Organization**
- Separated response types from database models
- Added constructor method `NewWordWithStats` for converting database models to response format
- Updated both list and detail endpoints to use the new response formats

The main challenge was getting the JSON response format to exactly match the API specification, particularly with nested fields like `parts` and handling of statistics. I also had to deal with some SQLite-specific database migration issues.

Next step is to run rspec for the rest of the endpoints without errors.


> 2025-02-15

## Summary of Work on endpoints using test database

1. After fixing the words endpoints, moved on to implementing study sessions and dashboard functionality:
   - Added proper pagination for study sessions
   - Implemented study session details endpoint
   - Added word listing for study sessions
   - Created dashboard endpoints for stats and last session

2. Added test data management to help with testing:
   - Created a ResetHistory endpoint to clear study data
   - Added ReloadTestData endpoint to restore test data
   - Implemented proper error handling for these endpoints
   - Fixed success rate calculation to return proper float values

3. The commands used are important for testing:
```bash
rm -f words.test.db        # Removes any existing test database
./scripts/init_test_db.sh  # Creates fresh test database with schema
DB_PATH=words.test.db go run cmd/server/main.go  # Runs server with test database
```

These commands ensure:
- Starting with a clean slate (rm -f)
- The database schema is properly initialized (init_test_db.sh)
- The server uses the test database instead of production (DB_PATH)

This setup allows us to:
- Run tests against a known state
- Avoid affecting production data
- Reset the database between test runs
- Have consistent test data across all tests

The key improvement was moving from hardcoded test data to a proper test database management system, making the tests more reliable and maintainable.

There are no more failures in the test database. 🍻

## Technical Specs Analysis & Frontend Implementation

Today I analyzed two frontend technical specification documents for a language learning portal project. The key findings were:
1. Specs Comparison
- Original spec focused on page-by-page requirements and API endpoints
- New spec provided comprehensive technical stack and implementation details
- Recommended using both specs together as they complement each other
2. Backend Review
- Examined the existing Go backend implementation
- Confirmed API endpoints match frontend requirements
- Noted SQLite database schema and data structures
3. Frontend Implementation
- Started with a Lovable-generated React project
- Added key missing components:
  - API client configuration for backend communication
  - TypeScript interfaces matching backend data structures
  - Sample implementation of StudyActivities page
4. Tech Stack Alignment
- Confirmed frontend tech choices align with requirements:
  - React with TypeScript
  - Tailwind CSS
  - ShadcN UI components
  - Vite.js
  - React Query for API data fetching


Next Steps:
- Implement remaining page components
- Add error handling and loading states
- Unit test the frontend code
- Implement pagination for lists
- Add breadcrumb navigation 
- Set up dark mode toggle


> 2025-02-17

## Installing openvino on my Intel AI PC

https://docs.openvino.ai/2025/get-started/install-openvino/install-openvino-pip.html
In powershell:
```bash
python -m venv openvino_env
openvino_env\Scripts\activate
(openvino_env) PS C:\Users\xxxxx> python -m pip install --upgrade pip

(openvino_env) PS C:\Users\xxxxx> python -m pip install openvino

(openvino_env) PS C:\Users\xxxxx> python -c "from openvino import Core; print(Core().available_devices)"
ZE_LOADER_DEBUG_TRACE:Using Loader Library Path:
ZE_LOADER_DEBUG_TRACE:Tracing Layer Library Path: ze_tracing_layer.dll
['CPU', 'GPU', 'NPU']
```

> 2025-02-19

## LLM Deployment Session Summary

For detailed steps see the following file:
- [OPEA Components summary](../opea-comps/vllm/README.md)


**Goal**
Deploy **Mistral-7B-Instruct-v0.3** using **vLLM** with **Intel Arc GPU acceleration** in **WSL2**.

**Challenges**
- **WSL2 GPU passthrough for Intel Arc** isn’t straightforward.
- OpenVINO GPU support requires specific setup.
- **vLLM** and **OpenVINO** version compatibility issues.

**Issues Encountered**
- **[GPU] Can't get PERFORMANCE_HINT**: Error related to missing GPU configuration or drivers.
- **OpenVINO GPU support** in container environment needed special setup.

**Possible Solutions**
- **CPU-Only Deployment**: Works immediately, simpler setup, but slower inference.
- **Native Windows Deployment**: Allows direct access to the Intel Arc GPU, but requires different setup.
- **Alternative Frameworks**: Explore **Intel’s optimized transformers** or **PyTorch with Intel extensions**.

**Next Steps**
1. Short-term: Deploy with **CPU-only** version for testing.
2. Long-term: Either **set up native Windows** environment for GPU acceleration or implement with **Intel’s optimized frameworks**.


**References**
- [vLLM Documentation](https://github.com/vllm-project/vllm)
- [OpenVINO GPU Support](https://docs.openvino.ai/latest/openvino_docs_install_guides_installing_openvino_docker.html)
- [Intel Arc GPU Setup](https://www.intel.com/content/www/us/en/developer/articles/guide/getting-started-with-intel-oneapi-base-toolkit-in-wsl-2.html) 


## Exploration of different models

1. **Considered MarianMT Model**
   - Found it's a specialized translation model
   - Smaller (298MB) than Mistral
   - Specifically designed for translation tasks

2. **Explored OPEA Components**
   - Looked at `/comps/third_parties`
   - Considered TGI (Text Generation Inference)
   - Found TGI only supports Xeon and Gaudi, not Intel NPU

3. **Back to vLLM**
   - Decided to try vLLM again since it has Intel hardware support
   - Attempted to build the container using `Dockerfile.intel_gpu`
   - Currently encountering BuildKit issues with Docker

**Current Status:**
- Trying to build vLLM with Intel GPU support
- Getting BuildKit-related errors in Docker build
- Need to resolve Docker build issues to proceed



## Translation Mega Service Development

Created a FastAPI service that translates English to Ancient Greek using Mistral.

**Setup and Development**

1. Created project structure:
```bash
mkdir -p opea-comps/translation-mega-service
cd opea-comps/translation-mega-service
```

2. Created requirements.txt [requirements.txt](../opea-comps/translation-mega-service/requirements.txt)


3. Created Dockerfile [Dockerfile](../opea-comps/translation-mega-service/Dockerfile)


4. Built and ran the service:
```bash
# Set Hugging Face token
export HF_TOKEN="my_token"

# Build image
docker build -t translation-mega-service .

# Run container
docker run -d \
    --name translation-mega-service \
    -p 8000:8000 \
    -e HF_TOKEN=${HF_TOKEN} \
    --memory=8g \
    --memory-swap=16g \
    translation-mega-service

# Check logs
docker logs -f translation-mega-service
```

**Testing the Service**

Multiple ways to test:

1. Using curl:
```bash
curl -X POST http://localhost:8000/v1/translate \
  -H "Content-Type: application/json" \
  -d '{"text": "Where is the dog? He needs to eat."}'
```

2. Using Python:
```python
import requests

response = requests.post(
    "http://localhost:8000/v1/translate",
    json={"text": "Where is the dog? He needs to eat."}
)
print(response.json())
```

**Key Features**

- Uses Mistral-7B-Instruct-v0.3 for translation
- FastAPI for the web service
- Docker containerization
- Few-shot prompting for better translations
- Optimized for CPU usage
- Memory-efficient configuration

**Challenges Solved**

1. Model loading optimization:
   - Used CPU-specific PyTorch
   - Added memory limits
   - Optimized token generation

2. Translation improvements:
   - Added few-shot examples
   - Optimized prompt structure
   - Added proper error handling

3. Docker configuration:
   - Security best practices
   - Memory management
   - Environment variable handling


I tried to use other smaller models like TinyLlama-1.1B-Chat but the translation was not good.
For every update:
```
# Stop and remove old container
docker stop translation-mega-service
docker rm translation-mega-service

# Build with new name
docker build -t translation-mega-service .

# Run with new name and ensure HF_TOKEN is passed
docker run -d \
    --name translation-mega-service \
    -p 8000:8000 \
    -e HF_TOKEN=${HF_TOKEN} \
    --memory=8g \
    translation-mega-service


docker logs -f translation-mega-service
```

**It works, but...**

1. The container still takes too long to start.
2. The model takes long to generate the translation.