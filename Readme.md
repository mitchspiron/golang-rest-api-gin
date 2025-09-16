# Command for migrate

# Required: Installed migrate CLI

https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md

> > > migrate create -ext sql -dir ./cmd/migrate/migrations -seq create_tablename_table

# Command to run migration and push it in database

> > > go run ./cmd/migrate/main.go up
