BEGIN;

-- Remove the foreign key constraint and 'roleId' column from 'QuestionSets' table
ALTER TABLE QuestionSets DROP CONSTRAINT fk_questionsets_roles;
ALTER TABLE QuestionSets DROP COLUMN roleId;

-- Add back the 'role' column to the 'QuestionSets' table
ALTER TABLE QuestionSets ADD COLUMN role TEXT;

-- Add the 'deleted' column to the 'Questions' table
ALTER TABLE Questions ADD COLUMN deleted BOOLEAN;

-- Remove the 'name' column from the 'QuestionSets' table
ALTER TABLE QuestionSets DROP COLUMN name;

-- Remove the unique constraint on the 'name' column
ALTER TABLE QuestionSets DROP CONSTRAINT questionsets_name_unique;

-- Add back the 'questions' column to the 'QuestionSets' table
ALTER TABLE QuestionSets ADD COLUMN questions BIGINT[];

-- Commit the changes
COMMIT;
