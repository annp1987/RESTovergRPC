package main

import (
	"context"
	"os"
	"sync"
	"github.com/annp1987/RESTovergRPC/server"
)

var (
	host     = os.Getenv("DB_USERS_HOST")
	port     = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_NAME")
	password = os.Getenv("DB_USERS_PASSWORD")
)

func Init() {
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "1234"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "postgres"
	}
	if password == "" {
		password = "postgres"
	}
}
func main() {
	ctx := context.Background()
	// for test locally with go run main.go
	Init()
	dburl := map[string]string {
		"Host": host,
		"Port": port,
		"User": user,
		"Type": dbname,
		"Password": password,
	}
	go server.StartGRPC(ctx, dburl)

	go server.StartHTTP()

	// Block forever
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()

}
