## Задание

1. Попробовать применить Clean Architecture + Hexagon (Ports & Adapters):
    1. продумать доменные модели (entities)
    2. разделить на уровни:
        1. бизнес логика (usecases)
        2. транспортный слой (delivery)
        3. слой доступа к данным (repository, services, adapters)
2. Реализовать сервисы по чистой архитектуре:
    1. delivery
    2. usecases
    3. (имплементацию репозиториев можно не делать)
3. ⭐ Внедрить DI контейнеры (`github.com/uber-go/fx, github.com/uber-go/dig, github.com/google/wire, ...`)

## **Полезные ссылки**

- [https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [https://amitshekhar.me/blog/go-backend-clean-architecture](https://outcomeschool.com/blog/go-backend-clean-architecture)
- [https://habr.com/ru/companies/oleg-bunin/articles/516500/](https://habr.com/ru/companies/oleg-bunin/articles/516500/)
- [https://github.com/evrone/go-clean-template/tree/master](https://github.com/evrone/go-clean-template/tree/master)
- [https://github.com/bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)
- [Внедрение зависимостей в GO](https://habr.com/ru/articles/541676/)
