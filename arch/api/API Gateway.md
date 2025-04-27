
```yaml
name: API Gateway
responsibility: Единая точка входа для клиентов

data: []
api:
    - /api/*
publishes_events: []
consumes_events: []
interacts_with:
    - Auth Service
    - User/Profile Service
    - Friend Service
    - Chat/Message Service
    - Notification Service
    - Media Service

```



