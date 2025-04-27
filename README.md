# balun.courses-microservice-6stream
- lessons - https://github.com/moguchev/microservices_like_in_bigtech_6

# Make new service
1. Run
```shell
./create_service.sh {you project name}
```
2. Add service to compose `services/compose.yaml`
3. Add service to nginx configuration `services/nginx/nginx.conf`

# Workflow

### Run all services
```shell
cd services
make compose-up
```

### Stop all services
```shell
cd services
make compose-stop
```

### Chesk services
```shell
cd services
make check_service
```