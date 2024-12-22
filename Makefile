start:
	@docker-compose up -d && docker-compose logs -f

purge:
	@docker-compose down --remove-orphans

build:
	@docker-compose build

rebuild:
	@docker-compose build --no-cache

stop:
	@docker-compose down

logs:
	@docker-compose logs -f --tail=100

restart:
	@docker-compose restart
	@docker-compose logs -f

postgres:
	@docker exec -it emil-portfolio-postgres-1 psql -U postgres -W emil-portfolio

.PHONY: start stop logs restart rebuild purge build
