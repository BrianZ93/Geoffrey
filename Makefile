run_local:
	docker compose -f docker-compose.yaml build
	docker compose -f docker-compose.yaml up

stop_local:
	docker compose -f docker-compose.yaml down