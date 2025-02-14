-- Test Words
INSERT INTO words (ancient_greek, greek, english, parts, created_at, updated_at) VALUES 
('λύω', 'λύνω', 'to loose, destroy', 
 '{"present": "λύω", "future": "λύσω", "aorist": "ἔλυσα", "perfect": "λέλυκα"}',
 DATETIME('now'), DATETIME('now')),
('γράφω', 'γράφω', 'to write', 
 '{"present": "γράφω", "future": "γράψω", "aorist": "ἔγραψα", "perfect": "γέγραφα"}',
 DATETIME('now'), DATETIME('now'));

-- Test Groups
INSERT INTO groups (name, created_at, updated_at) VALUES 
('Basic Verbs', DATETIME('now'), DATETIME('now')),
('Common Nouns', DATETIME('now'), DATETIME('now'));

-- Link Words to Groups
INSERT INTO words_groups (word_id, group_id) VALUES 
(1, 1),  -- λύω in Basic Verbs
(2, 1);  -- γράφω in Basic Verbs 