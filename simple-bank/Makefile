postgres:
	docker run --name some-postgres-name -p 5432:5432 -e POSTGRES_USER=root_user -e POSTGRES_PASSWORD=mysecretpassword -d postgres:12-alpine
createdb:
	docker exec -it some-postgres-name createdb --usename=root_user --owner=root_user simple_bank
dropdb:
	docker exec -it some-postgres-name dropdb --usename=root_user --owner=root_user simple_bank
.PHONY: postgres createdb dropdb