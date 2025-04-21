
```yaml
name: Notification Service
responsibility: Уведомления (push, email)

data: Notifications
api:
    - POST /notify
publishes_events:
    - NotificationSent
consumes_events:
    - MessageSent
    - FriendRequestSent
    - FriendAccepted
interacts_with:
    - User/Profile Service
    - Chat/Message Service
    - Friend Service
```


