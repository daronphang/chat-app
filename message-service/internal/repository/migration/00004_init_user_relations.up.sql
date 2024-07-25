CREATE TABLE IF NOT EXISTS user_relation (
    userId text,
    relationId text,
    createdAt timestamp,
    PRIMARY KEY (userId, relationId)
);