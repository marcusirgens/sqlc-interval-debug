.PHONY: testdatabase test teardown

test: testdatabase
	go test ./... -tags withdb \
		-database-url "postgres://sqlc:sqlc@localhost:15432/sqlc" \
		-timeout 30s

# start a testing database with the migrate.sql as the init script
testdatabase:
	(docker ps --format "{{.Names}}" | grep "interval-failure-test-db" > /dev/null) || \
	docker run \
	--rm --name interval-failure-test-db -p 15432:5432 \
	--detach \
	-v "$$(pwd)/migrate.sql:/docker-entrypoint-initdb.d/00-init.sql" \
	-e "POSTGRES_DB=sqlc" \
	-e "POSTGRES_PASSWORD=sqlc" \
	-e "POSTGRES_USER=sqlc" \
	postgres:9.6-alpine

teardown:
	docker stop interval-failure-test-db
