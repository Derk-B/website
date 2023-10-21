CREATE TABLE projects (
    id serial PRIMARY KEY,
    title TEXT NULL,
    description TEXT NULL,
    banner_url TEXT NULL,
    banner_alt TEXT NULL,
    project_url TEXT NULL,
    timestamp BIGINT
);