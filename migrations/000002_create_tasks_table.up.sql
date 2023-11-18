CREATE TABLE IF NOT EXISTS tdlist.tasks (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    user_id UUID REFERENCES tdlist.users(id)
);