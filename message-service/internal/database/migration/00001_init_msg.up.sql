BEGIN;
CREATE TABLE IF NOT EXISTS message(
    messageId bigint,
    previousMessageId bigint,
    channelId text,
    senderId text,
    type text,
    content text,
    createdAt date,
    PRIMARY KEY (channelId, messageId)
) WITH CLUSTERING ORDER BY (messageId DESC);

CREATE TABLE IF NOT EXISTS channel(
    channelId text PRIMARY KEY,
    userId text,
    createdAt date
);

CREATE TABLE IF NOT EXISTS user(
    userId text PRIMARY KEY,
    associatedUserId text,
    channelId text,
    createdAt date
);
COMMIT;