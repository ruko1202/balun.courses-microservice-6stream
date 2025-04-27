
```yaml
name: User/Profile Service
responsibility: Профиль, никнейм, аватар, about

data: Profiles
api:
    - GET /profile
    - PUT /profile
    - GET /users?search=
publishes_events:
    - ProfileUpdated
consumes_events:
    - UserRegistered
interacts_with:
    - Auth Service
    - Friend Service
    - Media Service
```


