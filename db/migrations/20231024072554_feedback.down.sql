BEGIN;

-- Revert the 'status' column data type change and remove the 'closed' status
ALTER TABLE PracticeSessions
ALTER COLUMN status
TYPE TEXT
USING status::TEXT;

-- Remove the 'closed' status from the 'status' column
UPDATE PracticeSessions
SET status = 'completed'
WHERE status = 'closed';

-- Remove and re-add the 'chk_practicesessions_status' constraint
ALTER TABLE PracticeSessions
DROP CONSTRAINT IF EXISTS chk_practicesessions_status;
ALTER TABLE PracticeSessions
ADD CONSTRAINT chk_practicesessions_status CHECK (status IN ('in_progress', 'completed'));

-- Remove and re-add the 'chk_practicesessions_completedat' constraint
ALTER TABLE PracticeSessions
DROP CONSTRAINT IF EXISTS chk_practicesessions_completedat;
ALTER TABLE PracticeSessions
ADD CONSTRAINT chk_practicesessions_completedat CHECK (completedAt IS NULL OR status = 'completed');

-- Drop the 'Feedback' and 'FeedbackEntries' tables
DROP TABLE IF EXISTS Feedback;
DROP TABLE IF EXISTS FeedbackEntries;

COMMIT;
