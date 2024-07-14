# Chat Application

## Requirements

- Support for 1-on-1 and group chats
- Real-time chat communication
- Support for large scale daily active users (DAU)
- Online indicator
- E2E encryption
- Persistent storage of messages
- Push notification
- Multiple device support
- Chats are sorted by the latest timestamp of last message
- Support for media upload
- Mark unread messages

# Architecture design

## Communication

### Sending and receiving messages

As we need a globally distributed system, users will be connected to different chat servers.

For both sending and receiving messages, websocket is preferred between client and chat server as it is bi-directional and persistent.

### Others

For other features including signup, authentication, profile change, retrieval of older messages, REST over HTTP is preferred.

### Between internal services

For communication between internal services, gRPC is preferred over HTTP.

## Service discovery

The primary role of service discovery is to recommend the best chat server for a client based on the criteria like geographical location, server capacity, etc. **etcd** is a popular open-source solution for service discovery. It registers all the available chat servers and picks the best chat server for a client based on predefined criteria.

<img src="./assets/service-discovery.png">

1. User A tries to log in to the app
2. The load balancer sends the login request to API servers
3. After the backend authenticates the user, service discovery finds the best chat server for User A
4. User A connects to chat server 2 through WebSocket

## Databases

### Generic data

For generic data related to user identity and chat associations, a **relational** database is preferred.

### Message

For core messaging functionality, as it generates a huge volume of traffic, swift read/write operations for large data volumes is required, and **non-relational key-value** data store (Cassandra) is preferred for the following reasons:

- Key-value stores allow easy horizontal scaling
- Key-value stores provide very low latency to access data
- Relational databases do not handle long tail of data well; when the indexes grow large, random access is expensive
- Key-value stores are adopted by other proven reliable chat applications

In Cassandra, records are sharded by the partition keys. On each node, records with the same partition key are sorted by the sort key.

## Schemas

### Message

When a message is received, it will be pushed to a queue. Kafka is chosen as it provides the following guarantees:

- Message order
- Scalability

Key features of a message:

- IDs must be unique
- IDs should be sortable by time (cannot rely on createdAt as two messages can be created at the same time)
- Each message belongs to a channel
- Channel can either refer to a group or 1-on-1 chat
- prevMsgId helps when there is a communication breakdown between chat servers; if a successful message is received but the prevMsgId is mismatched, client will retrieve all the messages from the current prevMsgId

For generating IDs, there are three approaches:

- Auto increment feature in SQL, but NoSQL does not provide such feature
- Use a global 64-bit sequence number
- Use local sequence number generator; this is sufficient as maintaining order within a channel is sufficient

```json
{
  "msgId": 123546341,
  "prevMsgId": 12315141,
  "channelId": "p6o5n4m3l2-k1j0-i9h8-g7f6-e5d4c3b2a1",
  "senderId": "5e4d3c2b-1a0p-9o8n-7m6l-5k4j3i2h1g0f9e8d7",
  "message": "Hello, how are you?",
  "createdAt": "2023-09-17T10:30:00.000"
}
```

### User-and-channel relationship

There are two implementations available.

For option 1, we have one table where the partition key is the channelId and the sort key is the userId. If NoSQL is used, a secondary index is not needed as it does not improve performance for high-cardinality columns in KV store (channelId-userId is unique). Otherwise, you can create secondary index on userId column.

```json
{
  "channelId": "p6o5n4m3l2-k1j0-i9h8-g7f6-e5d4c3b2a1",
  "userId": "5e4d3c2b-1a0p-9o8n-7m6l-5k4j3i2h1g0f9e8d7"
}
```

For option 2, two tables are created:

- First table will use channelId as the partition key and userId as sort key; this is used for broadcasting messages from a channel
- Second table will use userId as the partition key and channelId as sort key; this is used for determining the channels a user belongs to

### User-and-chat-server relationship

Each user can have multiple chat-servers if multiple devices are used. SQL can be used to store this data.

A heartbeat mechanism can be implemented to determine if the client is still connected to the chat server in the event of network disruption.

```json
{
  "chatServerId": "https://chatserver1.com",
  "userId": "5e4d3c2b-1a0p-9o8n-7m6l-5k4j3i2h1g0f9e8d7",
  "lastActiveTimestamp": "2023-09-17T10:30:00.000"
}
```

## Message flow

1. User A sends a chat message to Chat server 1
2. Chat server 1 receives message and generates a message ID from ID generator
3. Chat server 1 sends message to message sync queue and ack message delivery
4. Message server pulls message from queue and stores in key-value store
5. Message is forwarded to user B's message queue
6. If user B is online, Chat server 2 pulls message from queue and sends message to user B via websocket connection
7. If user B is offline, a push notification is sent to user B

## Message synchronization across multiple devices

## Small group chats

## Online presence

Presence servers are responsible for managing online status and communicating with clients through WebSocket.

### User login

After a WebSocket connection is built between the client and the real-time service, user A’s online status and last_active_at timestamp are saved in the KV store. Presence indicator shows the user is online after he/she logs in.

```json
{
  "status": "online",
  "lastActiveTimestamp": "2023-09-17T10:30:00.000"
}
```

### User logout

When a user logs out, it goes through the user logout flow. The online status is changed to offline in the KV store. The presence indicator shows a user is offline.

### User disconnection

When a user disconnects from the internet, the persistent connection between the client and server is lost. A naive way to handle user disconnection is to mark the user as offline and change the status to online when the connection re-establishes. However, this approach has a major flaw. It is common for users to disconnect and reconnect to the internet frequently in a short time.

Instead, we introduce a **heartbeat mechanism** to solve this problem. Periodically, an online client sends a heartbeat event to presence servers. If presence servers receive a heartbeat event within a certain time, a user is considered as online. Otherwise, it is offline. This can be done through CRON job that checks the last active timestamp and current timestamp.

## Services

### API service

- Handles HTTP requests including authentication, signup, user profile change, etc.
- Data is stored in relational store

### Chat service

- Maintains websocket connection with client
- Real-time notifications are pushed and forwarded to websocket

### Service discovery service

- Utilizes etcd as service discovery
- Provides API for adding chat server to service registry
- Maintains

### Notification service

- Provides best-effort delivery
- AWS SNS is used as third-party service

### Presence service

- Maintains all users' online status and last active timestamp in a relational data store
- Broadcasts online status to friends
- Redis as KV data store
