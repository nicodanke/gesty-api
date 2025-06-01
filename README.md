# API Gasty

Microservice architectured API made in GO. Use gRPC to communicate between micro services and gRPC Gateway to allow HTTP1.1 requests (like a regular RestAPI).

## Database

We use Postgres as database, each microservice is in charge of its own database. For local development we run the databases inside a docker compose file, this file runs the necessary databases to run the application.

To generate database documentation we use [dbml/cli](https://dbml.dbdiagram.io/cli). This tool can be install with the following command: `sudo npm install -g @dbml/cli`. Previously you have to install [Node](https://nodejs.org/en/download) in your computer.

Once installed, we can add the database documentation inside `doc/db.dbml` and visualize it in this [web page](https://databasediagram.com/app). After the doc/db.dbml is complete, then you can run make `make db_schema` inside the service to generate a file called `doc/schema.sql`. This file contains SQL DDL that will be used in the migrations.


## Migrations

In order to run migrations you need to install golang-migration CLI ([docs](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)). If you are using Mac, you can run: `brew install golang-migrate` to install it.

Migrations are run or created inside each service. The services are inside services folder.


### Create new migration

In order to create a new migration you can run in the root of the service: `migrate create -ext sql -dir db/migrations <migration_name>`.

### Run migration

To run a migration in each service you have a Makefile that has the following commands:

* `make migrate-up`: run all migrations
* `make migrate-up1`: run the last migration
* `make migrate-down`: revert all migrations
* `make migrate-down1`: revert the last migration


## SQLC (Query and model generation)

We use SQLC in order to create queries and map database tables with models. In order to install SQLC you have to run: `brew install sqlc`

### Add new queries

In order to add a new query you just need to add the query inside `db/query` in a `.sql` file. To add the query you just have to follow the [SQLC documentation](https://docs.sqlc.dev/en/latest/howto/select.html#). After adding the necessary queries, you just have to run

`make sqlc-generate`

to generate the SQL queries and models. The models and queries are going to be generated inside `db/sqlc` folder.


## Installation

For the installation and run of the application you can do it manually

```bash
pip install foobar
```

## Usage

```python
import foobar

# returns 'words'
foobar.pluralize('word')

# returns 'geese'
foobar.pluralize('goose')

# returns 'phenomenon'
foobar.singularize('phenomena')
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)