
## To build the postgresdb
`docker build -tpostgresdb-c-db ./ `
### And run it
`docker run -d --name my-postgresdb-container -p 5432:5432 my-postgres-db`

### Run gin webservice
`go run main.go`
### Available calls atm
* /albums
* /albums/:id
* /countries
* /cities