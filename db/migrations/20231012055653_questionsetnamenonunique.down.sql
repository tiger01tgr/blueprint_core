BEGIN;

-- Recreate the unique constraint
ALTER TABLE QuestionSets
ADD CONSTRAINT questionsets_name_unique UNIQUE (name);

COMMIT;
