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
func main() {
	ctx := context.Background()
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
