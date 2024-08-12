-- Cassandra limitation to run one command per file.
-- Secondary indexes pulls data from multiple partitions, hence using ORDER BY is not allowed
CREATE INDEX messageStatus ON message (messageStatus);