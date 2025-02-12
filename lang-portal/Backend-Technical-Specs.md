# Backend Server Technical Specs

## Business Goal: 

A language learning school wants to build a prototype of learning portal which will act as three things:
- Inventory of possible vocabulary that can be learned
- Act as a  Learning record store (LRS), providing correct and wrong score on practice vocabulary
- A unified launchpad to launch different learning apps

## Technical Requirements

- The backend will be built using Go
- The database will be SQLite3
- The API will be built using Gin
- Mage is a task runner for Go.
- The API will always return JSON
- There will no authentication or authorization
- Everything be treated as a single user

## Directory Structure

```text
backend_go/
├── cmd/
│   └── server/           # Main application entry point
├── internal/
│   ├── handlers/             # API handlers
│   ├── models/          # Data models and database operations
│   └── services/        # Business logic
├── db/                  # Database-related files
│   ├── migrations/      # SQL migration files
│   └── seeds/          # Seed data JSON files
├── magefile.go          # Mage task definitions
├── words.db            # SQLite database
└── go.mod              # Go module file
```

## Database Schema

Our database will be a single sqlite database called `words.db` that will be in the root of the project folder of `backend_go`

We have the following tables:

- words - stores vocabulary words
  - id integer
  - ancient greek string
  - greek string
  - english string
  - parts json
- words_groups - join table for words and groups many-to-many
  - id integer
  - word_id integer
  - group_id integer
- groups - thematic groups of words
  - id integer
  - name string
- study_sessions - records of study sessions grouping word_review_items
  - id integer
  - group_id integer
  - created_at datetime
  - study_activity
- study_activities - a specific study activity, linking a study session to group
  - id integer
  - study_session_id integer
  - group_id integer
  - created_at datetime
- word_review_items - a record of word practice, determining if the word was correct or not
  - id integer
  - word_id integer
  - study_session_id integer
  - correct boolean
  - created_at datetime


### API Endpoints
- GET /api/dashboard/last_study_session
- GET /api/dashboard/study_progress
- GET /api/dashboard/quick-stats
- GET /api/study_activities
- GET /api/study_activities/:id
- GET /api/study_activities/:id/study_sessions
- POST /api/study_activities
  - required params: group_id, study_activity_id
- GET /api/words
  - pagination with 100 items per page
- GET /api/words/:id
- GET /api/groups
  - pagination with 100 items per page
- GET /api/groups/:id
- GET /api/groups/:id/words
- GET /api/groups/:id/study_sessions
- GET /api/study_sessions
  - pagination with 100 items per page
- GET /api/study_sessions/:id
- GET /api/study_sessions/:id/words
- POST /api/reset_history
- POST /api/full_reset
- POST /api/study_sessions/:id/words/:word_id/review
  - required params: 
    - id (study_session_id) integer
    - word_id integer
    - correct boolean


### API Endpoints

#### GET /api/dashboard/last_study_session
Returns information about the most recent study session.
Response:
```json
{
  "id": 123,
  "group_id": 456,
  "group_name": "Basic Greetings",
  "study_activity_id": 456,
  "created_at": "2024-03-20T15:30:00Z",
}
```

#### GET /api/dashboard/study_progress
Returns study progress statistics over time. Please note that the frontend will determine progress bar based on total words studied and total available words.
Response:
```json
{
  "total_words_studied": 15,
  "total_available_words": 300
}
```

#### GET /api/dashboard/quick-stats
Returns summary statistics for the learning progress.
Response:
```json
{
  "success_rate": 80.0,
  "total_study_sessions": 4,
  "total_active_groups": 3,
  "study_streak_days": 4
}
```

#### GET /api/study_activities
Returns list of all study activities.
Response:
```json
{
  "items": [
    {
      "id": 1,
      "name": "Vocabulary Quiz",
      "description": "Test your knowledge of Greek vocabulary",
      "thumbnail_url": "https://example.com/thumbnail.jpg"
    },
    {
      "id": 2,
      "name": "Word Explorer",
      "description": "Interactive exploration of word meanings and forms",
      "thumbnail_url": "https://example.com/word-explorer.jpg"
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "total_items": 2,
    "items_per_page": 100
  }
}
```

#### GET /api/study_activities/:id
Returns details of a specific study activity.
Response:
```json
{
  "id": 1,
  "name": "Vocabulary Quiz",
  "description": "Test your knowledge of Greek vocabulary",
  "thumbnail_url": "https://example.com/thumbnail.jpg"
}
```

#### GET /api/study_activities/:id/study_sessions
Returns all study sessions for a specific activity.
Response:
```json
{
  "items": [
    {
      "id": 123,
      "activity_name": "Vocabulary Quiz",
      "group_name": "Basic Greetings",
      "start_time": "2025-02-12T17:20:23-05:00",
      "end_time": "2025-02-12T17:30:23-05:00",
      "review_items_count": 20
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "total_items": 100,
    "items_per_page": 20
  }
}
```

#### POST /api/study_activities
Creates a new study activity session.
Request:
```json
{
  "group_id": 456,
  "study_activity_id": 1
}
```
Response:
```json
{
  "id": 123,
  "group_id": 456,
  "study_activity_id": 1,
  "created_at": "2024-03-20T15:30:00Z"
}
```

#### GET /api/words
Returns paginated list of vocabulary words.
Response:
```json
{
  "items": [
    {
      "id": 123,
      "ancient_greek": "Χαῖρε",
      "greek": "Γειά",
      "english": "Hello",
      "parts": {
        "present": "χαίρω",
        "future": "χαιρήσω",
        "aorist": "ἐχάρην",
        "perfect": "κεχάρηκα"
      },
      "correct_count": 5,
      "wrong_count": 2
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "total_items": 500,
    "items_per_page": 100
  }
}
```

#### GET /api/words/:id
Returns details of a specific word.
Response:
```json
{
  "id": 123,
  "ancient_greek": "Χαῖρε",
  "greek": "Γειά",
  "english": "Hello",
  "parts": {
    "present": "χαίρω",
    "future": "χαιρήσω",
    "aorist": "ἐχάρην",
    "perfect": "κεχάρηκα"
  },
  "stats": {
    "correct_count": 5,
    "wrong_count": 2
  },
  "groups": [
    {
      "id": 123,
      "name": "Basic Greetings"
    }
  ]
}

```

#### GET /api/groups
Returns paginated list of word groups.
Response:
```json
{
  "items": [
    {
      "id": 123,
      "name": "Basic Greetings",
      "word_count": 20
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "total_items": 10,
    "items_per_page": 100
  }
}
```

#### GET /api/groups/:id
Returns details of a specific group.
Response:
```json
{
  "id": 123,
  "name": "Basic Greetings",
  "stats": {
    "total_word_count": 20
  }
}
```

#### GET /api/groups/:id/words
Returns all words in a specific group.
Response:
```json
{
  "items": [
    {
      "id": 123,
      "ancient_greek": "Χαῖρε",
      "greek": "Γειά",
      "english": "Hello",
      "parts": {
        "present": "χαίρω",
        "future": "χαιρήσω",
        "aorist": "ἐχάρην",
        "perfect": "κεχάρηκα"
      },
      "correct_count": 5,
      "wrong_count": 2
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "total_items": 20,
    "items_per_page": 100
  }
}
```

#### GET /api/groups/:id/study_sessions
Returns all study sessions for a specific group.
Response:
```json
{
  "items": [
    {
      "id": 123,
      "activity_name": "Vocabulary Quiz",
      "group_name": "Basic Greetings",
      "start_time": "2025-02-12T17:20:23-05:00",
      "end_time": "2025-02-12T17:30:23-05:00",
      "review_items_count": 20
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "total_items": 5,
    "items_per_page": 100
  }
}
```

#### GET /api/study_sessions
Returns paginated list of all study sessions.
Response:
```json
{
  "items": [
    {
      "id": 123,
      "activity_name": "Vocabulary Quiz",
      "group_name": "Basic Greetings",
      "start_time": "2025-02-12T17:20:23-05:00",
      "end_time": "2025-02-12T17:30:23-05:00",
      "review_items_count": 20
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "total_items": 100,
    "items_per_page": 100
  }
}
```

#### GET /api/study_sessions/:id
Returns details of a specific study session.
Response:
```json
{
  "id": 123,
  "activity_name": "Vocabulary Quiz",
  "group_name": "Basic Greetings",
  "start_time": "2025-02-12T17:20:23-05:00",
  "end_time": "2025-02-12T17:30:23-05:00",
  "review_items_count": 20
}
```

#### GET /api/study_sessions/:id/words
Returns all words reviewed in a specific study session.
Response:
```json
{
  "items": [
    {
      "ancient_greek": "Χαῖρε",
      "greek": "Γειά",
      "english": "Hello",
      "correct_count": 5,
      "wrong_count": 2
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "total_items": 20,
    "items_per_page": 100
  }
}
```

#### POST /api/reset_history
Resets all study history while keeping words and groups.
Response:
```json
{
  "success": true,
  "message": "Study history has been reset"
}
```

#### POST /api/full_reset
Resets everything in the database.
Response:
```json
{
  "success": true,
  "message": "Database has been reset to initial state"
}
```

#### POST /api/study_sessions/:id/words/:word_id/review
Records a word review result.
Request Payload:
```json
{
  "correct": true
}
```
Response:
```json
{
  "word_id": 1,
  "success": true,
  "study_session_id": 1,
  "correct": true,
  "created_at": "2024-02-12T15:30:00Z"
}


## Task Runner Tasks

Mage is a make/rake-like build tool using Go that allows writing build tasks in pure Go code.
It will be used to handle database initialization, migrations, and other maintenance tasks.

### Initialize Database
This task will initialize the sqlite database called `words.db`

### Migrate Database
This task will run a series of migrations SQL files on the database.

Migrations live in the `migrations` folder. The migration files will be run in order of their file name. The file names should look like this:

```
0001_init.sql
0002_create_words_table.sql
0003_create_groups_table.sql
0004_add_word_parts_column.sql
```

### Seed Data
This task will import JSON files and transform them into target data for our database.

All seed files live in the `seeds` folder.
All seed files should be loaded.
In our task we should have DSL to specify each seed file and its expected group word name. The JSON structure for Ancient Greek vocabulary should look like this:

```json
[
  {
    "ancient_greek": "λύω",
    "greek": "λύνω",
    "english": "to loose, destroy",
    "parts": {
      "present": "λύω",
      "future": "λύσω",
      "aorist": "ἔλυσα",
      "perfect": "λέλυκα"
    }
  },
  ...
]
```

Each seed file can be associated with a specific group (e.g., "Basic Verbs", "Common Nouns", "Greetings") during the import process.

### Required Tasks
- Reset: Clear all data and reinitialize the database
- ResetHistory: Clear only study history while preserving vocabulary and groups
- Backup: Create a backup of the current database
- Restore: Restore database from a backup file
- ImportWords: Import vocabulary from CSV/JSON files
- ExportWords: Export vocabulary to CSV/JSON files
- ValidateDB: Check database integrity and constraints
- Dev: Start the development server with hot reload

