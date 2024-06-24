.PHONY: all build run stop down

all: build run

build:
	@echo "Building Docker containers..."
	docker-compose build

run:
	@echo "Running Docker containers..."
	docker-compose up -d

stop:
	@echo "Stopping Docker containers..."
	docker-compose stop

down:
	@echo "Removing Docker containers..."
	docker-compose down

logs:
	@echo "Displaying logs..."
	docker-compose logs -f

clean:
	@echo "Cleaning up Docker containers and volumes..."
	docker-compose down -v
