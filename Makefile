start:
	@docker-compose up -d && docker-compose logs -f

stop:
	@docker-compose down

logs:
	@docker-compose logs -f --tail=100

restart:
	@docker-compose restart

.PHONY: start stop logs