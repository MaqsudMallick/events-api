# create container postgres
docker run --name my-postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
psql -U postgres
create database eventdb;

# dev environment for the program
wgo run main.go
