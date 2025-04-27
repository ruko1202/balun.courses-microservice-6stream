compose-rebuild: compose-stop
	docker-compose up -d --build
	sleep 1
	docker-compose ps -a

compose-up: compose-stop
	docker-compose up -d
	sleep 1
	docker-compose ps -a

compose-stop:
	docker-compose stop -t0
	make docker-prune

docker-prune:
	$(info docker prune)
	docker container prune -f
	docker image prune -f

.curl: url=
.curl:
	@echo "Запрос к $$url"; \
	curl -w "\n%{response_code}\n" "$$url"

.check_service: service=
.check_service:
	@make .curl url="localhost/$$service/version"
	@make .curl url="localhost/$$service/readiness"
	@make .curl url="localhost/$$service/liveness"

check_service:
	@make .check_service service=api-gateway
	@make .check_service service=auth
	@make .check_service service=chat-message
	@make .check_service service=friend
	@make .check_service service=media
	@make .check_service service=notification
	@make .check_service service=user-profile
