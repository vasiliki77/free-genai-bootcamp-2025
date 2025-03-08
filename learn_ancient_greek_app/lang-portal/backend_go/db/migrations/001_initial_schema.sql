-- Initial database schema

-- Words table
CREATE TABLE words (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    ancient_greek TEXT NOT NULL,
    greek TEXT NOT NULL,
    english TEXT NOT NULL,
    parts JSON
);

-- Groups table
CREATE TABLE groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    name TEXT NOT NULL UNIQUE
);

-- Words-Groups join table
CREATE TABLE words_groups (
    word_id INTEGER,
    group_id INTEGER,
    PRIMARY KEY (word_id, group_id),
    FOREIGN KEY (word_id) REFERENCES words(id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);

-- Study Activities table
CREATE TABLE study_activities (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    name TEXT NOT NULL,
    description TEXT,
    thumbnail_url TEXT
);

-- Study Sessions table
CREATE TABLE study_sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    group_id INTEGER NOT NULL,
    study_activity_id INTEGER NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups(id),
    FOREIGN KEY (study_activity_id) REFERENCES study_activities(id)
);

-- Word Reviews table
CREATE TABLE word_reviews (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    word_id INTEGER NOT NULL,
    study_session_id INTEGER NOT NULL,
    correct BOOLEAN NOT NULL,
    FOREIGN KEY (word_id) REFERENCES words(id),
    FOREIGN KEY (study_session_id) REFERENCES study_sessions(id)
);

-- Create indexes
CREATE INDEX idx_words_deleted_at ON words(deleted_at);
CREATE INDEX idx_groups_deleted_at ON groups(deleted_at);
CREATE INDEX idx_study_activities_deleted_at ON study_activities(deleted_at);
CREATE INDEX idx_study_sessions_deleted_at ON study_sessions(deleted_at);
CREATE INDEX idx_word_reviews_deleted_at ON word_reviews(deleted_at); 