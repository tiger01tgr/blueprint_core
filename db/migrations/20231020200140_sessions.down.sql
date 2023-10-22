BEGIN;

-- Drop the constraints
ALTER TABLE PracticeSessionSubmissions DROP CONSTRAINT IF EXISTS fk_practicesessionsubmissions_questions;
ALTER TABLE PracticeSessionSubmissions DROP CONSTRAINT IF EXISTS fk_practicesessionsubmissions_practicesessions;

ALTER TABLE PracticeSessions DROP CONSTRAINT IF EXISTS fk_practicesessions_questions;
ALTER TABLE PracticeSessions DROP CONSTRAINT IF EXISTS fk_practicesessions_questionsets;
ALTER TABLE PracticeSessions DROP CONSTRAINT IF EXISTS fk_practicesessions_users;

-- Drop the tables
DROP TABLE IF EXISTS PracticeSessionSubmissions;
DROP TABLE IF EXISTS PracticeSessions;

COMMIT;
