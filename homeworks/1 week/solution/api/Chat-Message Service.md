
```yaml
name: Chat/Message Service
responsibility: Сообщения, чаты

data: Chats, Messages
api:
    - POST /message
    - GET /messages
publishes_events:
    - MessageSent
consumes_events:
    - FriendAccepted
interacts_with:
    - Friend Service
    - Notification Service
```


