package main

import (
	DBFactory "github.com/projects/vue-go-1st/api/db"
	serverFactory "github.com/projects/vue-go-1st/api/server"
	"os"
	"strconv"
)

var port int
var dbHost string
var dbPort string
var dbUser string
var dbPassword string
var dbDatabase string

func init() {
	rawPort := os.Getenv("PORT")

	if len(rawPort) > 0 {

		var err error

		port, err = strconv.Atoi(rawPort)

		if err != nil {
			panic(err)
		}

	} else {
		port = 1000

		dbHost = "localhost"      /*os.Getenv("DB_HOST")*/
		dbPort = "3306"           /*os.Getenv("DB_PORT")*/
		dbUser = "root"           /*os.Getenv("DB_USER")*/
		dbPassword = ""           /*os.Getenv("DB_PASSWORD")*/
		dbDatabase = "vue-go-1st" /*os.Getenv("DB_DATABASE")*/
	}
}
func main() {
	db, err := DBFactory.Connect(dbHost, dbPort, dbUser, dbPassword, dbDatabase)
	if err != nil {
		panic(err)
	}

	server := serverFactory.NewServer(port, db)
	server.Start()
}
