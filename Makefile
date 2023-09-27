.PHONY: up-docker build-docker build-up

build-docker:
	docker-compose build

build-up:
	docker-compose up --build

down:
	docker-compose down

up:
	docker-compose up 

mado: 
	sudo docker-compose up --build


# before running don't forgot that we need to stop the PostgreSQL
# sudo service postgresql stop