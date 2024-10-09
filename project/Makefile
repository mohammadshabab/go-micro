SHELL=cmd.exe
FRONT_END_BINARY=frontApp.exe
BROKER_BINARY=brokerApp

# up starts all containers in the background without forcing build
up:
	@echo Starting docker images
	docker-compose up -d 
	@echo Docker images started!

# up_build stops docker compose if running, build and start docker compose
up_build: build_broker
	@echo Stopping docker images 
	docker-compose down
	@echo Building and starting docker images 
	docker-compose up --build -d
	@echo Docker images build and started!

# stop docker compose
down:
	@echo Stoping docker compose
	docker-compose down
	@echo Down Finished!

# builds the broker binary as linux executable 
build_broker:
	@echo Building broker binary 
	chdir ..\broker-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${BROKER_BINARY} ./cmd/api
	@echo Build Finished!

# builds the frontend binary
build_front:
	@echo Building frontend binary 
	chdir ..\front-end && set CGO_ENABLED=0&& set GOOS=windows&& go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo Build Finished!

# starts the front end
start: build_front
	@echo Starting front end
	chdir ..\front-end && start /B ${FRONT_END_BINARY} > nul 2>&1 &

# stop the frontend 
stop:
	@echo stopping front end
	@taskkill /IM "${FRONT_END_BINARY}" /F
	@echo Stopped the front end!	



