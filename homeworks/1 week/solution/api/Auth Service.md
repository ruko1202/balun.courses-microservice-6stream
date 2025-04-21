```yaml
name: Auth Service
responsibility: Регистрация, логин, OAuth, токены

data: Users, Credentials
api:
    - POST /register
    - POST /login
    - POST /oauth
publishes_events:
    - UserRegistered #после успешной регистрации
    - UserLoggedIn
consumes_events: []
interacts_with:
    - User/Profile Service
```