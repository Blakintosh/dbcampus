DROP TABLE IF EXISTS TeamManager CASCADE;
CREATE TABLE TeamManager(
    teamManagerID SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(120) NOT NULL,
    sessionID VARCHAR(120),
    jiraEmail VARCHAR(120),
    jiraApiToken VARCHAR(120),
    managerExperience DECIMAL(5,3)
);

DROP TABLE IF EXISTS Project CASCADE;
CREATE TABLE Project(
    projectCode VARCHAR(3) NOT NULL,
    username VARCHAR(50) NOT NULL REFERENCES TeamManager(username),
    projectName VARCHAR(120),
    projSuccess DECIMAL(3,2),
    budget DECIMAL(17,2),
    monthlyExpenses DECIMAL(17,2),
    customSpendings DECIMAL(17,2),
    deadline TIMESTAMP,
    teamMeanExperience DECIMAL(5,3),
    weeklyTeamMeetings DECIMAL(5,3),
    clientMeetingsPerMonth DECIMAL(5,3),
    jiraProjectCode VARCHAR(120),
    jiraURL VARCHAR(120),
    PRIMARY KEY (projectCode, username)
);

DROP TABLE IF EXISTS TeamSurveys CASCADE;
CREATE TABLE TeamSurveys(
    teamMetricsID SERIAL PRIMARY KEY,
    projectCode VARCHAR(50) UNIQUE NOT NULL REFERENCES Project(projectCode),
    surveyLink VARCHAR(120),
    formID VARCHAR(50) UNIQUE NOT NULL,
    supportFromTopManagement DECIMAL(3,2),
    testingQuality DECIMAL(3,2),
    documentationQuality DECIMAL(3,2),
    clarityOfRequirements DECIMAL(3,2),
    taskTooMuch DECIMAL(3,2),
    teamSatisfaction DECIMAL(3,2)
);

DROP TABLE IF EXISTS Client CASCADE;
CREATE TABLE Client(
    clientID SERIAL PRIMARY KEY
);

-- Insert dummy team manager data
INSERT INTO TeamManager (username, password, managerExperience) VALUES ('admin', '$2a$10$dW7ufHTFVBJdTh5LAqmmJ.PyLVOG6CvqUkawV/keN.PxtrgcMuRGq', 5.0);

-- Insert dummy project data
INSERT INTO Project (projectCode, username, projectName, budget, monthlyExpenses, customSpendings, deadline, teamMeanExperience, weeklyTeamMeetings, clientMeetingsPerMonth) VALUES ('ABC', 'admin', 'Project ABC', 1000000.00, 1000.00, 10000.00, '2023-12-31 23:59:59', 5.0, 1.0, 1.0);


-- DROP TABLE IF EXISTS ClientSurveys CASCADE;
-- CREATE TABLE ClientSurveys(
--     clientMetricsID SERIAL PRIMARY KEY,
--     clientID INTEGER NOT NULL REFERENCES Client(clientID),
--     projectID INTEGER NOT NULL REFERENCES Project(projectID),
--     currentProductSatisfaction DECIMAL(3,2),
--     scopeSatisfaction DECIMAL(3,2),
--     numberOfMeetings INTEGER,
--     CONSTRAINT unique_client_project UNIQUE (clientID, projectID)
-- );

-- DROP TABLE IF EXISTS ProjectCode CASCADE;
-- CREATE TABLE ProjectCode(
--     projectCodeID SERIAL PRIMARY KEY,
--     projectID INTEGER NOT NULL REFERENCES Project(projectID),
--     numberOfLangs INTEGER,
--     reuseLevel DECIMAL(3,2),
--     interfacingLevel DECIMAL(3,2),
--     testQuality DECIMAL(3,2),
--     codeErrorDensity DECIMAL(3,2)
-- );

-- DROP TABLE IF EXISTS ProjectRequirements CASCADE;
-- CREATE TABLE ProjectRequirements(
--     projectRequirementsID SERIAL PRIMARY KEY,
--     projectID INTEGER NOT NULL REFERENCES Project(projectID),
--     stability DECIMAL(3,2),
--     complexity DECIMAL(3,2),
--     clarity DECIMAL(3,2),
--     dependence DECIMAL(3,2),
--     realisticSchedule DECIMAL(3,2),
--     clearObjectives DECIMAL(3,2),
--     flexiDevEnviro DECIMAL(3,2)
-- );


-- Generate IDs ?

-- CREATE VIEW AllProjects AS
-- SELECT * FROM Project
-- ORDER BY projectID ASC;

-- CREATE VIEW AllClients AS
-- SELECT * FROM ClientMetrics
-- ORDER BY clientID ASC;

-- CREATE VIEW ProjectDetails AS
-- SELECT ProjectCode.*, ProjectRequirements.*
-- FROM ProjectCode
-- JOIN ProjectRequirements
-- ON ProjectCode.projectID = ProjectRequirements.projectID
-- WHERE ProjectCode.projectID = [your_projectID];

-- CREATE OR REPLACE FUNCTION prevent_duplicate_projectID()
-- RETURNS trigger AS $$
-- BEGIN
--     IF EXISTS (SELECT 1 FROM Project WHERE projectID = NEW.projectID) THEN
--         RAISE EXCEPTION 'Project with projectID % already exists.', NEW.projectID;
--     ELSE
--         RETURN NEW;
--     END IF;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE TRIGGER prevent_duplicate_projectID_trigger
-- BEFORE INSERT ON Project
-- FOR EACH ROW
-- EXECUTE FUNCTION prevent_duplicate_projectID();

-- CREATE OR REPLACE FUNCTION prevent_duplicate_TeamMetrics_projectID()
-- RETURNS trigger AS $$
-- BEGIN
--     IF EXISTS (SELECT 1 FROM TeamMetrics WHERE projectID = NEW.projectID) THEN
--         RAISE EXCEPTION 'TeamMetrics entry with projectID % already exists.', NEW.projectID;
--     ELSE
--         RETURN NEW;
--     END IF;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE TRIGGER prevent_duplicate_TeamMetrics_projectID_trigger
-- BEFORE INSERT ON TeamMetrics
-- FOR EACH ROW
-- EXECUTE FUNCTION prevent_duplicate_TeamMetrics_projectID();

-- CREATE OR REPLACE FUNCTION prevent_duplicate_ClientMetrics_projectID()
-- RETURNS trigger AS $$
-- BEGIN
--     IF EXISTS (SELECT 1 FROM ClientMetrics WHERE projectID = NEW.projectID) THEN
--         RAISE EXCEPTION 'ClientMetrics entry with projectID % already exists.', NEW.projectID;
--     ELSE
--         RETURN NEW;
--     END IF;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE TRIGGER prevent_duplicate_ClientMetrics_projectID_trigger
-- BEFORE INSERT ON ClientMetrics
-- FOR EACH ROW
-- EXECUTE FUNCTION prevent_duplicate_ClientMetrics_projectID();

-- CREATE OR REPLACE FUNCTION prevent_invalid_projectID()
-- RETURNS trigger AS $$
-- BEGIN
--     IF NOT EXISTS (SELECT 1 FROM Project WHERE projectID = NEW.projectID) THEN
--         RAISE EXCEPTION 'Invalid projectID %.', NEW.projectID;
--     ELSE
--         RETURN NEW;
--     END IF;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE TRIGGER prevent_invalid_projectID_trigger
-- BEFORE INSERT ON ProjectCode
-- FOR EACH ROW
-- EXECUTE FUNCTION prevent_invalid_projectID();

-- CREATE TRIGGER prevent_invalid_projectID_trigger
-- BEFORE INSERT ON ProjectRequirements
-- FOR EACH ROW
-- EXECUTE FUNCTION prevent_invalid_projectID();


-- CREATE OR REPLACE FUNCTION get_project_variables(p_project_id INTEGER)
-- RETURNS TABLE(
--     project_id INTEGER,
--     team_manager_id INTEGER,
--     proj_success DECIMAL(3,2),
--     budget INTEGER(64),
--     current_spend INTEGER(64),
--     next_deadline TIMESTAMP,
--     final_deadline TIMESTAMP,
--     prev_deadlines_met INTEGER(64),
--     deliver_frequency DECIMAL(3,2),
--     team_capability DECIMAL(3,2),
--     documentation_level DECIMAL(3,2),
--     perceived_task_difficulty DECIMAL(3,2),
--     project_satisfaction DECIMAL(3,2),
--     team_motivation DECIMAL(3,2),
--     team_happiness DECIMAL(3,2),
--     turnover INTEGER(64),
--     oauth_sub_id VARCHAR(100),
--     current_product_satisfaction DECIMAL(3,2),
--     scope_satisfaction DECIMAL(3,2),
--     number_of_meetings INTEGER(64),
--     manager_capability DECIMAL(3,2),
--     previous_project_success DECIMAL(3,2),
--     number_of_langs INTEGER(64),
--     reuse_level DECIMAL(3,2),
--     interfacing_level DECIMAL(3,2),
--     test_quality DECIMAL(3,2),
--     code_error_density DECIMAL(3,2),
--     stability DECIMAL(3,2),
--     complexity DECIMAL(3,2),
--     clarity DECIMAL(3,2),
--     dependence DECIMAL(3,2),
--     realistic_schedule DECIMAL(3,2),
--     clear_objectives DECIMAL(3,2),
--     flexi_dev_enviro DECIMAL(3,2)
-- ) AS $$
-- BEGIN
--     SELECT
--         projectID,
--         teamManagerID,
--         projSuccess,
--         budget,
--         currentSpend,
--         nextDeadline,
--         finalDeadline,
--         prevDeadlinesMet,
--         deliverFrequency,
--         teamCapability,
--         documentationLevel
--     INTO
--         project_id,
--         team_manager_id,
--         proj_success,
--         budget,
--         current_spend,
--         next_deadline,
--         final_deadline,
--         prev_deadlines_met,
--         deliver_frequency,
--         team_capability,
--         documentation_level
--     FROM Project
--     WHERE projectID = p_project_id;
    
--     SELECT
--         percievedTaskDifficulty,
--         projectSatisfaction,
--         teamMotivation,
--         teamHappiness,
--         turnover
--     INTO
--         perceived_task_difficulty,
--         project_satisfaction,
--         team_motivation,
--         team_happiness,
--         turnover
--     FROM TeamMetrics
--     WHERE projectID = p_project_id;
    
--     SELECT
--         OAuthSubID,
--         currentProductSatisfaction,
--         scopeSatisfaction,
--         numberOfMeetings
--     INTO
--         oauth_sub_id,
--         current_product_satisfaction,
--         scope_satisfaction,
--         number_of_meetings
--     FROM ClientMetrics
--     WHERE projectID = p_project_id;
    
--     SELECT
--         managerCapability,
--         previousProjectSuccess
--     INTO
--         manager_capability,
--         previous_project_success
--     FROM TeamManager
--     WHERE teamManagerID = (SELECT teamManagerID FROM Project WHERE projectID = p_project_id);
    
--     SELECT
--         numberOfLangs,
--         reuseLevel,
--         interfacingLevel,
--         testQuality,
--         codeErrorDensity
--     INTO
--         number_of_langs,
--         reuse_level,
--         interfacing_level,
--         test_quality,
--         code_error_density
--     FROM ProjectCode
--     WHERE projectID = p_project_id;
    
--     SELECT
--         stability,
--         complexity,
--         clarity,
--         dependence,
--         realisticSchedule,
--         clearObjectives,
--         flexiDevEnviro
--     INTO
--         stability,
--         complexity,
--         clarity,
--         dependence,
--         realisticSchedule,
--         clearObjectives,
--         flexiDevEnviro
--     FROM ProjectRequirements
--     WHERE projectID = p_project_id;
       
