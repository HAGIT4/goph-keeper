run-postgres-docker:
	docker run --name keeper-postgres -p 5432:5432 -e POSTGRES_USER=keeper_user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=keeper -d postgres

run-goose-migrations:
	cd ./migrations
	export GOOSE_DRIVER=postgres
	export GOOSE_DBSTRING="user=keeper_user password=password host=127.0.0.1 dbname=keeper sslmode=disable"
	goose up

generate-pb:
	buf generate