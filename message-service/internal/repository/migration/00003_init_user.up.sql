CREATE TABLE IF NOT EXISTS user (
    userId text,
    channelId text,
    createdAt timestamp,
    PRIMARY KEY (userId, channelId)
);