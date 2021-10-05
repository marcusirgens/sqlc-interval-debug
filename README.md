# sqlc-interval-debug

Contains code to reproduce an issue with `interval` types in [sqlc](https://github.com/kyleconroy/sqlc).

## Sample data used

[`sqlc.json`](./sqlc.json) and [`query.sql`](./query.sql) are copied from [github.com/kyleconroy/sqlc/internal/endtoend/testdata/interval/stdlib](https://github.com/kyleconroy/sqlc/tree/main/internal/endtoend/testdata/interval/stdlib), 
with the only change being the `Insert` query to test row insertion.

[`migrate.sql`](./migrate.sql) has some insert queries to pre-populate the database.

## Running the tests

Please inspect [`Makefile`](./Makefile) before you run anything, of course. There are two useful targets:

### `make test`

Starts a Postgres 9.6 instance listening on port 15432 using Docker and runs tests against it.

### `make teardown`

Removes the database.

# License

Licensed under [`MIT-0`](./LICENSE). 

sqlc is copyright (c) 2019 Kyle Conroy, [and is available under the MIT license](https://github.com/kyleconroy/sqlc/blob/944588ed96592ffee28c3d1d571d11615f9abc5b/LICENSE). 
