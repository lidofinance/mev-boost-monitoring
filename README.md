# go-template
go-skeleton-lido

 ## How to use the template
 
 1. Clone repository
 2. cd root repository
 3. make tools
 4. docker-composer up -d
 5. make migrate
 6. make build
 7. Run service ./bin/service

## How to create migrations?
 ./bin/migrate create -ext=sql -dir=db/migrations <your table name>

  