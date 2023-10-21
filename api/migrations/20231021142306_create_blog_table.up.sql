CREATE TABLE blog_posts (
    id serial PRIMARY KEY,
    title TEXT NULL,
    description TEXT NULL,
    content TEXT NULL,
    banner_url TEXT NULL,
    banner_alt TEXT NULL,
    timestamp BIGINT NULL
);