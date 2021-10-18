package main

import (
    "api/routes"
    "api/db"
)

func main() {
    db.DBMigrate(db.DBconnect())
    routes.Api()
}