
```yaml
name: Friend Service
responsibility: Дружба, заявки, подтверждения

data: FriendRequests, Friends
api:
    - POST /request
    - POST /accept
    - POST /decline
    - DELETE /friend
    - GET /friends
publishes_events:
    - FriendRequestSent
    - FriendAccepted
    - FriendRemoved
consumes_events:
    - UserRegistered
    - ProfileUpdated
interacts_with:
    - User/Profile Service
    - Chat/Message Service
```


