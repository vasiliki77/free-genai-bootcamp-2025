-- Test Groups
INSERT INTO groups (name, created_at, updated_at) VALUES 
('Basic Verbs', DATETIME('now'), DATETIME('now')),
('Common Nouns', DATETIME('now'), DATETIME('now')),
('Greetings', DATETIME('now'), DATETIME('now'));

-- Test Words
INSERT INTO words (ancient_greek, greek, english, parts, created_at, updated_at) VALUES 
('λύω', 'λύνω', 'to loose, destroy', 
 '{"present": "λύω", "future": "λύσω", "aorist": "ἔλυσα", "perfect": "λέλυκα"}',
 DATETIME('now'), DATETIME('now')),
('γράφω', 'γράφω', 'to write', 
 '{"present": "γράφω", "future": "γράψω", "aorist": "ἔγραψα", "perfect": "γέγραφα"}',
 DATETIME('now'), DATETIME('now')),
('χαίρω', 'χαίρω', 'to rejoice', 
 '{"present": "χαίρω", "future": "χαιρήσω", "aorist": "ἐχάρην", "perfect": "κεχάρηκα"}',
 DATETIME('now'), DATETIME('now'));

-- Link Words to Groups
INSERT INTO words_groups (word_id, group_id) VALUES 
(1, 1),  -- λύω in Basic Verbs
(2, 1),  -- γράφω in Basic Verbs
(3, 1),  -- χαίρω in Basic Verbs
(3, 3);  -- χαίρω also in Greetings

-- Add study activities with full details
INSERT INTO study_activities (name, description, thumbnail_url, created_at, updated_at) VALUES 
('Vocabulary Quiz', 'Test your knowledge of Greek vocabulary', 
 'https://example.com/vocab-quiz.jpg', DATETIME('now'), DATETIME('now')),
('Word Explorer', 'Interactive exploration of word meanings and forms', 
 'https://example.com/word-explorer.jpg', DATETIME('now'), DATETIME('now')),
('Writing Practice', 'Practice writing Greek words and sentences', 
 'https://example.com/writing.jpg', DATETIME('now'), DATETIME('now'));

-- Add study sessions across multiple days for streak testing
INSERT INTO study_sessions (group_id, study_activity_id, created_at, updated_at) VALUES
-- Today's sessions
(1, 1, DATETIME('now'), DATETIME('now')),
(2, 1, DATETIME('now'), DATETIME('now')),
-- Yesterday's sessions
(1, 2, DATETIME('now', '-1 day'), DATETIME('now', '-1 day')),
(3, 2, DATETIME('now', '-1 day'), DATETIME('now', '-1 day')),
-- 2 days ago sessions
(2, 3, DATETIME('now', '-2 days'), DATETIME('now', '-2 days')),
(1, 3, DATETIME('now', '-2 days'), DATETIME('now', '-2 days')),
-- 3 days ago sessions
(1, 1, DATETIME('now', '-3 days'), DATETIME('now', '-3 days')),
-- 5 days ago sessions (gap for streak testing)
(2, 2, DATETIME('now', '-5 days'), DATETIME('now', '-5 days'));

-- Add word reviews with mixed success rates
INSERT INTO word_reviews (word_id, study_session_id, correct, created_at, updated_at) VALUES
-- Today's reviews (75% success)
(1, 1, true, DATETIME('now'), DATETIME('now')),
(2, 1, true, DATETIME('now'), DATETIME('now')),
(3, 1, true, DATETIME('now'), DATETIME('now')),
(1, 2, false, DATETIME('now'), DATETIME('now')),
-- Yesterday's reviews (50% success)
(2, 3, true, DATETIME('now', '-1 day'), DATETIME('now', '-1 day')),
(3, 3, false, DATETIME('now', '-1 day'), DATETIME('now', '-1 day')),
-- 2 days ago reviews (100% success)
(1, 5, true, DATETIME('now', '-2 days'), DATETIME('now', '-2 days')),
(2, 5, true, DATETIME('now', '-2 days'), DATETIME('now', '-2 days')),
-- 3 days ago reviews (0% success)
(3, 7, false, DATETIME('now', '-3 days'), DATETIME('now', '-3 days')),
-- 5 days ago reviews (mixed success)
(1, 8, true, DATETIME('now', '-5 days'), DATETIME('now', '-5 days')),
(2, 8, false, DATETIME('now', '-5 days'), DATETIME('now', '-5 days')); 