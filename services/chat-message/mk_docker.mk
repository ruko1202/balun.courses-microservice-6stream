include mk_vars.mk

build-image:
	$(info build image ${SERVICE_NAME})
	docker build -f=./.build/Dockerfile --tag=${SERVICE_NAME} --build-arg=SERVICE_NAME=${SERVICE_NAME} ./
	make docker-prune
	docker images | grep ${SERVICE_NAME}

docker-run: docker-stop
	$(info run container ${SERVICE_NAME})
	docker run -d -p 8080:8080 --name ${SERVICE_NAME} ${SERVICE_NAME}
	docker ps -a

docker-stop:
	docker stop ${SERVICE_NAME} || true
	docker rm ${SERVICE_NAME} || true
	make docker-prune

docker-prune:
	$(info docker prune)
	docker container prune -f
	docker image prune -f
