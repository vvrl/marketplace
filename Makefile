docker:
	docker-compose -f docker-compose.yml up -d

.DEFAULT_GOAL := docker


stop: 
	docker-compose stop

db:
	docker exec -it market-psql psql market_user -d market_db