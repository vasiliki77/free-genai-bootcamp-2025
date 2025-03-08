-- Insert study activities
INSERT INTO study_activities (name, description, thumbnail_url) VALUES
    ('Vocabulary Quiz', 'Test your knowledge of Greek vocabulary', 'https://example.com/vocab-quiz.jpg'),
    ('Word Explorer', 'Interactive exploration of word meanings and forms', 'https://example.com/word-explorer.jpg'),
    ('Writing Practice', 'Practice writing Greek words and sentences', 'https://example.com/writing.jpg');

-- Insert word groups
INSERT INTO groups (name) VALUES
    ('Basic Verbs'),
    ('Common Nouns'),
    ('Greetings');

-- Insert words
INSERT INTO words (ancient_greek, greek, english, parts) VALUES
    ('λύω', 'λύνω', 'to loose, destroy', 
        '{"present": "λύω", "future": "λύσω", "aorist": "ἔλυσα", "perfect": "λέλυκα"}'),
    ('γράφω', 'γράφω', 'to write', 
        '{"present": "γράφω", "future": "γράψω", "aorist": "ἔγραψα", "perfect": "γέγραφα"}'),
    ('Χαῖρε', 'Γειά', 'Hello', 
        '{"present": "χαίρω", "future": "χαιρήσω", "aorist": "ἐχάρην", "perfect": "κεχάρηκα"}');

-- Link words to groups
INSERT INTO words_groups (word_id, group_id) VALUES
    (1, 1), -- λύω -> Basic Verbs
    (2, 1), -- γράφω -> Basic Verbs
    (3, 3); -- Χαῖρε -> Greetings 