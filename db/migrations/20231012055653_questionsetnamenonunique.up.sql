BEGIN;

ALTER TABLE QuestionSets
DROP CONSTRAINT IF EXISTS questionsets_name_unique;

COMMIT;