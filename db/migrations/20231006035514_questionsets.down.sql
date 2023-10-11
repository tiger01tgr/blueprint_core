BEGIN;

-- Drop the foreign key constraint from Questions to QuestionSets
ALTER TABLE Questions
DROP CONSTRAINT IF EXISTS questions_questionsetid_fkey;

-- Drop the foreign key constraint from QuestionSets to Employers
ALTER TABLE QuestionSets
DROP CONSTRAINT IF EXISTS questionsets_employerid_fkey;

-- Drop the foreign key constraint from Employers to Industries
ALTER TABLE Employers
DROP CONSTRAINT IF EXISTS employers_industryid_fkey;

-- Drop the unique constraint on Employers' name
ALTER TABLE Employers
DROP CONSTRAINT IF EXISTS employers_name_key;

-- Drop the unique constraint on Employers' id
ALTER TABLE Employers
DROP CONSTRAINT IF EXISTS employers_pkey;

-- Drop the unique constraint on Industries' id
ALTER TABLE Industries
DROP CONSTRAINT IF EXISTS industries_pkey;

ALTER TABLE Industries
DROP CONSTRAINT IF EXISTS industries_name_key;

-- Drop the Employers table
DROP TABLE IF EXISTS Employers;

-- Drop the QuestionSets table
DROP TABLE IF EXISTS QuestionSets;

-- Drop the Questions table
DROP TABLE IF EXISTS Questions;

-- Drop the Industries table
DROP TABLE IF EXISTS Industries;

COMMIT;
