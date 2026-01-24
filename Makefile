up:
	cd deploy && docker compose --env-file ../.env up -d --build

down:
	cd deploy && docker compose --env-file ../.env down

logs:
	cd deploy && docker compose --env-file ../.env logs -f

clean:
	cd deploy && docker compose --env-file ../.env down -v
