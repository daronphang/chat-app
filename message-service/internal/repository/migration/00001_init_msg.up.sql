-- Cassandra limitation to run one command per file.
CREATE TABLE IF NOT EXISTS message (
    messageId bigint,
    channelId text,
    senderId text,
    messageType text,
    content text,
    messageStatus text,
    createdAt timestamp,
    PRIMARY KEY (channelId, messageId)
) WITH CLUSTERING ORDER BY (messageId DESC);