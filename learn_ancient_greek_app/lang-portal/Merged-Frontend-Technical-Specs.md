# Frontend Technical Specification

## Business Goal: 

A language learning school wants to build a prototype of learning portal which will act as three things:
- Inventory of possible vocabulary that can be learned
- Act as a  Learning record store (LRS), providing correct and wrong score on practice vocabulary
- A unified launchpad to launch different learning apps

## Technology Stack
The technical stack should be the following:
- React.js as the frontend library
- Tailwind CSS as the css framework
- Vite.js as the local development server
- Typescript for the programming language
- ShadCN for components

## Frontend Routes
This is a list of routes for our web-app we are building
Each of these routes are a page and we'll describe them in more details under the pages heading.

- /dashboard
- /study_activities
- /study_activity/:id
- /words
- /words/:id
- /groups
- /groups/:id
- /study_sessions
- /study_session/:id
- /settings

The default route / should forward to /dashboard

## Navigation

There will be horizontal navigation bar with the following links:
- Dashboard
- Study Activities
- Words
- Word Groups
- Sessions
- Settings

## Breadcrumbs

Beneath the navigation there will be breadcrumbs so users can easily see where they are. Examples of breadcrumbs:

Dashboard
Study Activities > Adventure MUD
Study Activities > Typing Tutor
Words > Ἄρξαι
Word Groups > Core Verbs

## Pages

### Dashboard
This page provides a summary of the student's progression.

- **Last Session**

### Study Activities Index

#### Route
The route for this page: `/study-activities`

#### Description
This is a grid of cards that represent an activity.

#### Card Structure
A card has:
- **Thumbnail**
- **Title**
- **"Launch" button**
- **"View" button**

### Launch Button Behavior
- The **Launch** button will open a new address in a new tab.
- Study activities are their own apps, but in order for them to launch, they need to be provided a `group_id`.

#### Example:

`localhost:8080?group_id=4`

### Study Activities Show

This page requires no pagination because there is unlikely to be more than 20 possible study activities.

#### Navigation
- The **View** button will go to the Student Activities Show Page.

#### Route
- The route for this page: `/study-activities/:id`

#### Information Section
This page will have an information section that contains:
- **Thumbnail**
- **Title**
- **Description**
- **Launch Button**

## Study Sessions List
There will be a list of sessions for this study activity. Each session item will contain:

- **Group Name**: So you know what group name was used for the sessions.
  - This will be a **link** to the Group Show Page.
- **Start Time**: When the session was created in `YYYY-MM-DD HH:MM` format (12-hour time).
- **End Time**: When the last `word_review_item` was created.
- **Review Items**: The number of review items.

### Word Groups Index

#### Route
The route for this page: `/groups`

#### Table Structure
This is a table of words with the following columns:

- **Ancient Greek**: The word in Ancient Greek script
  - This will also contain a small button to play the pronunciation of the word.
  - The Ancient Greek word will be a link to the **Word Details** page.

- **Greek**: The corresponding Modern Greek version of the word.

- **English**: The English translation of the word.

- **Correct**: The number of times the word was reviewed correctly.

- **Wrong**: The number of times the word was reviewed incorrectly.

## Display Limit
- Only **50 words** should be displayed at a time.

## Pagination
- **Previous button**: Greyed out if you cannot go further back.
- **Page X of Y**: The current page should be **bolded**.
- **Next button**: Greyed out if you cannot go further forward.

## Sorting Functionality
- All table **headings should be sortable**. Clicking a heading toggles between **ascending (ASC) and descending (DESC) order**.
- An **ASCII arrow** should indicate sorting direction:
  - **ASC (A → Z or 1 → 9)**: Arrow **pointing down (↓)**.
  - **DESC (Z → A or 9 → 1)**: Arrow **pointing up (↑)**.

## Words Show

The route for this page: `/words/:id`

# Word Groups Index

## Route
The route for this page: `/word-groups`

## Table Structure
This is a table of word groups with the following columns:

- **Group Name**: The name of the group.
  - This will be a **link** to the Word Groups Show page.

- **Words**: The number of words associated with this group.

This page contains the same **sorting** and **pagination** logic as the **Words Index** page.

---

## Word Groups Show

### Route
The route for this page: `/word-groups/:id`

This has the same components as the **Words Index**, but it is **scoped to only show words associated with this group**.

---

## Sessions Index

### Route
The route for this page: `/sessions`

- This page contains a **list of sessions**, similar to the **Study Activities Show** page.
- This page follows the same **sorting** and **pagination** logic as the **Words Index** page.

---

## Settings Page

### Route
The route for this page: `/settings`

## Reset History Button

This has a button that allows us to **reset the entire database**.

- We need to **confirm** this action in a dialog.
- The user must **type the word "reset me"** to confirm.

---

## Dark Mode Toggle

This is a **toggle** that changes from **light to dark theme**.




## Project Structure
```text
frontend/
├── src/
│   ├── api/              # API client and types
│   ├── components/       # Reusable UI components
│   ├── pages/           # Page components
│   ├── hooks/           # Custom React hooks
│   ├── types/           # TypeScript interfaces
│   └── utils/           # Utility functions
```

## API Types
```typescript
interface Pagination {
  current_page: number;
  total_pages: number;
  total_items: number;
  items_per_page: number;
}

interface WordParts {
  present: string;
  future: string;
  aorist: string;
  perfect: string;
}

interface Word {
  id: number;
  ancient_greek: string;
  greek: string;
  english: string;
  parts: WordParts;
  correct_count?: number;
  wrong_count?: number;
}
```

## Pages and Components

### Dashboard Page `/dashboard`
**API Endpoints:**
- GET /api/dashboard/last_study_session
- GET /api/dashboard/study_progress
- GET /api/dashboard/quick-stats

**Components:**
- LastStudySession
- StudyProgress
- QuickStats
- StartStudyButton

### Study Activities Page `/study_activities`
**API Endpoints:**
- GET /api/study_activities
- POST /api/study_activities

**Components:**
- ActivityCard
- ActivityList
- LaunchActivityForm

### Words Page `/words`
**API Endpoints:**
- GET /api/words
- GET /api/words/:id

**Components:**
- WordList
- WordDetails
- PaginationControls

### Groups Page `/groups`
**API Endpoints:**
- GET /api/groups
- GET /api/groups/:id
- GET /api/groups/:id/words
- GET /api/groups/:id/study_sessions

**Components:**
- GroupList
- GroupDetails
- GroupWordList
- GroupStudySessions

### Study Sessions Page `/study_sessions`
**API Endpoints:**
- GET /api/study_sessions
- GET /api/study_sessions/:id
- GET /api/study_sessions/:id/words
- POST /api/study_sessions/:id/words/:word_id/review

**Components:**
- SessionList
- SessionDetails
- WordReviewForm

### Settings Page `/settings`
**API Endpoints:**
- POST /api/reset_history
- POST /api/full_reset

**Components:**
- ResetHistoryButton
- FullResetButton
- ConfirmationDialog

## Data Fetching Strategy
```typescript
// Example using React Query
export function useWords(page: number) {
  return useQuery(['words', page], () => 
    axios.get(`/api/words?page=${page}&per_page=100`)
  );
}

export function useGroup(id: number) {
  return useQuery(['group', id], () => 
    axios.get(`/api/groups/${id}`)
  );
}
```

## Error Handling
- API error responses follow backend format:
```typescript
interface ErrorResponse {
  error: string;
}
```
- Global error boundary for React errors
- Toast notifications for API errors

## Pagination
- All lists use server-side pagination
- Default page size: 100 items
- Pagination controls component

## Testing Strategy
- Jest for unit tests
- React Testing Library for component tests
- MSW for API mocking

## Build and Development
```bash
npm run dev     # Start development server
npm run build   # Build for production
npm run test    # Run tests
``` 