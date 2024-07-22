CREATE TABLE IF NOT EXISTS channel (
    channelId text,
    userId text,
    createdAt timestamp,
    PRIMARY KEY (channelId, userId)
);