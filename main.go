package main

import (
	"jwt-authentication-golang/database"
)

func main() {
	// Initialize Database
	database.Connect("root:root@tcp(localhost:3306)/jwt_auth_golang?parseTime=true")
	database.Migrate()
}
