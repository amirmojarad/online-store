package main

import (
	"context"
	"fmt"
	"log"
	"online-supermarket/controllers/db"
	"online-supermarket/controllers/ent"
	"online-supermarket/views/api"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// TODO creating api in views
// TODO add edges
// TODO auth
// TODO adding gRPC
// TODO Create a service for shared media

func main() {

	dbConf := db.GetDatabaseConfig()
	log.Println(dbConf.Username, dbConf.Dbname, dbConf.Password, "asd")
	client, err := ent.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbConf.Username, dbConf.Dbname, dbConf.Password))

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()
	client = client.Debug()

	log.Println("Database Connected!")

	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Second)

	defer cancel()

	// Dump migration changes to stdout.
	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	api.RunAPI(ctx, client)
}
