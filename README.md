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
The database migration files are in `/migrations` so feel free to simply source them directly. Alternatively, you can apply them using `migrate` by running:
```bash
$ export POSTGRESQL_URL="postgres://$PG_USER:$PG_PASS@localhost:5432/$PG_DB?sslmode=disable"
$ migrate -database ${POSTGRESQL_URL} -path /migrations up
```
_**NOTE:** Remember to replace the `$PG*` variables with their actual values_

### Curl sing up
```
curl -X POST http://localhost:8080/user/sing-up -H "Content-type: application/json" \
     -d '{ "name": "Anton", "username": "tony", "password": "123"}'
```

After sing up we get a token in the logs. This is done for ease of testing. And add it to the request header

### Curl example add company
```
curl --location --request POST 'localhost:8080/company/add' \
--header 'Authorization: Bearer YOUR_TOKEN' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "OOO_KOLOBOK",
    "description": "MAKE KOLOBOK",
    "amount_of_employees": 34,
    "registered": true,
    "type": "Cooperative"
}''
```