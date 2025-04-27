## Задание

1. Для каждого сервиса из Домашнего задания 2 выбрать подход API (gRPC, REST, JSON-RPC, GraphQL) и аргументировать почему выбрали его
2. Описать для каждого сервиса в _Swagger/OpenAPI/GraphQL/Protobuf schema_ ваш API
3. Реализовать сервер обслуживающий API (REST/JSON-RPC/GraphQL/gRPC server). Обработка входящих запросов, валидация запросов (без бизнес логики) и отдача заглушки в виде ответа.
4. ⭐ Создать REST/JSON-RPC/GraphQL/gRPC клиентов (интегрировать общение сервисов между собой)
5. Добавить grpc-gateway и генерацию OpenAPI
6. ⭐ Добавить linter и форматирование _proto_ файлов
7. ⭐ Прикрутить в сервисе с grpc-gateway SwaggerUI сервер, с помощью которого можно выполнять запросы на grpc-gateway proxy.

## **Полезные ссылки**

- [swagger editor](https://editor.swagger.io/)
- [swagger-api/swagger-codegen](https://github.com/swagger-api/swagger-codegen)
- [go-swagger](https://github.com/go-swagger/go-swagger)
- [swaggo/swag](https://github.com/swaggo/swag)
- [Стандартный макет Go проекта](https://github.com/golang-standards/project-layout/blob/master/README_ru.md)
- [go HTTP сервер](https://pkg.go.dev/net/http#hdr-Servers)
- [echo framework](https://echo-labstack-com.translate.goog/docs/quick-start?_x_tr_sl=en&_x_tr_tl=ru&_x_tr_hl=ru&_x_tr_pto=sc&_x_tr_hist=true)
- [graphql-go](https://www.howtographql.com/graphql-go/1-getting-started/)
- [gqlgen](https://gqlgen.com/getting-started/)
- [https://habr.com/ru/companies/ruvds/articles/444346/](https://habr.com/ru/companies/ruvds/articles/444346/)
- [gRPC](https://grpc.io/)
- [protoc](https://grpc.io/docs/protoc-installation/)
- [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- [protovalidate](https://github.com/bufbuild/protovalidate)
- [https://gist.github.com/StevenACoffman/fe5f7774c750926210b642a0997479b0](https://gist.github.com/StevenACoffman/fe5f7774c750926210b642a0997479b0)
- [https://medium.com/@pointgoal/grpc-how-to-add-swagger-ui-on-grpc-466e5fd71097](https://medium.com/@pointgoal/grpc-how-to-add-swagger-ui-on-grpc-466e5fd71097)
- [https://github.com/moguchev/gofunc_autumn_2024](https://github.com/moguchev/gofunc_autumn_2024)
