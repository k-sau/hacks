package main

import (
	"context"
	"fmt"
	"os"
	"time"
	"flag"

	"github.com/jackc/pgx"
)

func main() {
	domain := flag.String("d", "", "Domain name")
	flag.Parse()
	if *domain == "" {
		fmt.Fprintf(os.Stderr, "Domain flag is not defined\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	conn, err := pgx.Connect(context.Background(), os.Getenv("FINDOMAIN_DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	
	domains, err := conn.Query(context.Background(), "SELECT * FROM subdomains WHERE name LIKE '%" + *domain + "%'")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	for domains.Next() {
		var id int32
		var name string
		var timestamp time.Time

		err := domains.Scan(&id, &name, &timestamp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to read data: %v\n", err)
		}
		fmt.Printf("%s\n", name)
	}
}