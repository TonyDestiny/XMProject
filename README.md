### Requirements
* Docker and Go
* [golang-migrate/migrate](https://github.com/golang-migrate/migrate)

### Usage
Clone the repository with:
```bash
git clone gitlab.com/TonyDestiny/XMProject
```

Copy the `env.example` file to a `.env` file.
```bash
$ cp .env.example .env
```

Build and start the services with:
```bash
$ docker-compose up --build
```
The database migration files are in `db/migrations` so feel free to simply source them directly. Alternatively, you can apply them using `migrate` by running:
```bash
$ export POSTGRESQL_URL="postgres://$PG_USER:$PG_PASS@localhost:5432/$PG_DB?sslmode=disable"
$ migrate -database ${POSTGRESQL_URL} -path db/migrations up
```
_**NOTE:** Remember to replace the `$PG*` variables with their actual values_


### Run postgres
```
docker run --name=postgres -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
```
### apply migrate
```
migrate -path ./migration -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
```

### Curl sing up
```
curl -X POST http://localhost:3333/user/sing-up -H "Content-type: application/json" \
     -d '{ "name": "Anton", "username": "tony", "password": "123"}'
```

### Curl example add company
```
curl -X POST http://localhost:3333/company/add -H "Content-type: application/json" \
     -d '{
            "id": "245feffd-36ac-4eb7-afbd-f67ba40b4439",
            "name": "OOO_KOLOBOK",
            "description": "MAKE KOLOBOK",
            "amount_of_employees": 34,
            "registered": true,
            "type": "Cooperative"
}'
```