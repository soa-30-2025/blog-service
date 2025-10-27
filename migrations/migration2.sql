CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS likes (
    blog_id UUID REFERENCES blogs(id) ON DELETE CASCADE,
    user_id VARCHAR(50) NOT NULL,
    PRIMARY KEY (blog_id, user_id)
);
