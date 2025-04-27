### Авторизация пользователя
```plantuml
@startuml
title Авторизация пользователя

actor Client
participant "API Gateway" as APIGW
participant "Auth Service" as Auth

Client -> APIGW : POST /login (email, password)
APIGW -> Auth : POST /login
Auth --> APIGW : JWT токен (и refresh токен)
APIGW --> Client : JWT токен

@enduml

```

### Регистрация нового пользователя
```plantuml
@startuml
title Регистрация нового пользователя

actor Client
participant "API Gateway" as APIGW
participant "Auth Service" as Auth
participant "User/Profile Service" as Profile

Client -> APIGW : POST /register (email, password)
APIGW -> Auth : POST /register
Auth --> APIGW : OK (user_id)
Auth -> Profile : Event: UserRegistered (user_id)
Profile --> Auth : OK (профиль создан)

@enduml

```

### Отправка запроса на установление дружбы
```plantuml
@startuml
title Отправка запроса на установление дружбы

actor Client
participant "API Gateway" as APIGW
participant "Friend Service" as Friend
participant "Notification Service" as Notify

Client -> APIGW : POST /friends/request (user_id)
APIGW -> Friend : POST /request
Friend --> APIGW : OK (заявка создана)
Friend -> Notify : Event: FriendRequestSent
Notify -> Client : Push/Web уведомление получателю заявки

@enduml

```

### Отправка сообщения пользователю (друг)
```plantuml
@startuml
title Отправка сообщения пользователю (друг)

actor Client
participant "API Gateway" as APIGW
participant "Chat/Message Service" as Chat
participant "Friend Service" as Friend
participant "Notification Service" as Notify

Client -> APIGW : POST /messages (текст, получатель)
APIGW -> Chat : POST /messages
Chat -> Friend : Проверка дружбы отправителя и получателя
Friend --> Chat : Дружба подтверждена
Chat --> APIGW : OK (сообщение сохранено)
Chat -> Notify : Event: MessageSent
Notify -> Client : Push/Web уведомление получателю

@enduml

```

### Отправка сообщения пользователю (НЕ друг)
```plantuml
@startuml
title Отправка сообщения пользователю (НЕ друг)

actor Client
participant "API Gateway" as APIGW
participant "Chat/Message Service" as Chat
participant "Friend Service" as Friend

Client -> APIGW : POST /messages (текст, получатель)
APIGW -> Chat : POST /messages
Chat -> Friend : Проверка дружбы отправителя и получателя
Friend --> Chat : Нет дружбы
Chat --> APIGW : Ошибка (нельзя отправить сообщение, не друг)

@enduml

```
