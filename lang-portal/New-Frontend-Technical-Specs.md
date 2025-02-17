# Frontend Technical Specification

## Business Goal

The frontend will serve as an interactive language learning portal that:
- Displays vocabulary inventory for study and practice.
- Tracks and visualizes learning progress based on backend study records.
- Provides a launchpad for different learning applications.

## Technology Stack

- **Framework**: React (Next.js for server-side rendering and performance optimization)
- **State Management**: Redux Toolkit or React Context API
- **UI Library**: Tailwind CSS and Headless UI
- **Data Fetching**: Axios or React Query
- **Routing**: Next.js built-in routing
- **Charts and Graphs**: Recharts for visualizing study progress
- **Build Tool**: Vite (if Next.js is not used) or Webpack
- **Task Runner**: NPM/Yarn scripts

## Directory Structure

```text
frontend/
├── public/                  # Static assets
├── src/
│   ├── components/          # Reusable UI components
│   ├── pages/               # Page components
│   ├── hooks/               # Custom React hooks
│   ├── context/             # Context Providers for global state
│   ├── services/            # API interaction logic
│   ├── utils/               # Utility functions
│   ├── styles/              # Global and component styles
│   ├── store/               # Redux store (if using Redux)
├── package.json             # Project dependencies and scripts
├── next.config.js           # Next.js configuration (if used)
└── tailwind.config.js       # Tailwind configuration
```

## Core Features and UI Components

### Dashboard
- Displays **last study session** details via `GET /api/dashboard/last_study_session`
- Shows **study progress** via `GET /api/dashboard/study_progress`
- Quick stats on **success rate, total study sessions, active groups, and study streak** via `GET /api/dashboard/quick-stats`
- UI Components:
  - Progress bar for overall learning progress
  - Study streak counter
  - Quick stats cards

### Vocabulary Explorer
- Fetches and paginates words using `GET /api/words`
- Word details page via `GET /api/words/:id`
- UI Components:
  - Search and filter options for words
  - Card-based word listing
  - Word details modal with translations and grammatical parts

### Groups and Thematic Learning
- Fetches groups via `GET /api/groups`
- Lists words in a group via `GET /api/groups/:id/words`
- Displays study sessions per group via `GET /api/groups/:id/study_sessions`
- UI Components:
  - Group selection dropdown
  - Word list per group
  - Study session history table

### Study Activities
- Lists available study activities via `GET /api/study_activities`
- Shows study sessions per activity via `GET /api/study_activities/:id/study_sessions`
- Starts a new study session via `POST /api/study_activities`
- UI Components:
  - Activity selection grid
  - Activity details page
  - Start study session button

### Study Sessions
- Lists all study sessions via `GET /api/study_sessions`
- Fetches details of a study session via `GET /api/study_sessions/:id`
- Displays words practiced in a session via `GET /api/study_sessions/:id/words`
- UI Components:
  - Study session history table
  - Session details modal
  - Correct/Wrong answer statistics

### Word Review & Practice
- Records word review using `POST /api/study_sessions/:id/words/:word_id/review`
- UI Components:
  - Flashcard-based word practice interface
  - Correct/Wrong buttons for answer validation
  - Streak and accuracy indicators

### Reset & Maintenance Tools
- Reset history via `POST /api/reset_history`
- Full reset via `POST /api/full_reset`
- UI Components:
  - Reset history button (with confirmation modal)
  - Full reset button (with admin permissions warning)

## API Integration
- **Data Fetching Strategy**: React Query or Axios
- **Error Handling**: Toast notifications for API errors
- **Caching**: Next.js ISR (Incremental Static Regeneration) for frequently used endpoints

## Navigation and Routing
- `/dashboard` - User’s main landing page with progress overview
- `/words` - Vocabulary explorer with filtering
- `/words/:id` - Word details
- `/groups` - List of word groups
- `/groups/:id` - Group-specific word list and progress
- `/activities` - Available learning activities
- `/activities/:id` - Activity-specific sessions
- `/sessions` - Study session history
- `/sessions/:id` - Study session details

## UI/UX Considerations
- **Mobile Responsiveness**: Fully responsive with Tailwind CSS
- **Dark Mode Support**: Using Tailwind’s dark mode utilities
- **Performance Optimizations**:
  - Lazy loading images with `next/image`
  - Client-side pagination for large datasets
  - Memoization of API results

## Testing Strategy
- **Unit Testing**: Jest & React Testing Library
- **Integration Testing**: Cypress for E2E tests
- **Linting & Code Quality**: ESLint & Prettier

## Deployment & Hosting
- **Hosting Platform**: Vercel or Netlify (if using Next.js)
- **CI/CD**: GitHub Actions for automatic testing & deployment
- **Environment Variables**: `.env.local` for API base URL

## Task Runner Commands
- `npm run dev` - Start development server
- `npm run build` - Build project for production
- `npm run start` - Start production server
- `npm run lint` - Run ESLint checks
- `npm run test` - Run unit tests
- `npm run e2e` - Run end-to-end tests

## Future Enhancements
- **User Accounts**: Add authentication when needed
- **Leaderboard & Gamification**: Track top learners
- **Speech Recognition**: Use Web Speech API for pronunciation practice
- **Offline Mode**: Implement PWA support for offline learning

This frontend spec ensures alignment with the backend API while focusing on usability and performance for the language learning portal.

