# EmpaticaServer

Restful Golang Server which manage Articles.

The Articles are saved in a Mysql Relational Database.

`GET http://<endpoint>/articles` - List all the Articles.

`GET http://<endpoint>/articles/<id>` - List the Article with the provided id.

`POST http://<endpoint>/articles` - Create an Article using the provided Payload. 

`Example:
{
	"id": "5",
    "title": "My Title", 
    "description": "My description", 
    "content": "My content"
}`

`PUT http://<endpoint>/articles/<id>` --> Update an Article using the provided Payload.  Create the Article if the id doesn't exist

`Example:
{
    "id": "5",
    "title": "My Title", 
    "description": "My description", 
    "content": "My content"
}`

`DELETE http://<endpoint>/articles/<id>` -  Delete the Article with the provided id

## Database Migration

The project uses the module [golang-migrate](http://github.com/golang-migrate/migrate) to manage every changes to the database. The migration files are store in the path db/migrations


## Run locally with a Local Mysql Database
Clone the project and run it using an existent Local Mysql Database 

`export DB_ENDPOINT=localhost:3306 && export DB_USERNAME=<username> && export DB_PASSWORD=<password> && export DB_NAME=<database_name> && go run *.go`

The server will be available to the address `http:/localhost/`

Open in a browser and get the Articles stored in the database:

`http://localhost/articles`





