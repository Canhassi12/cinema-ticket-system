package main

import (
	"context"
	"log"
	"ticket-system/src/config"
	"ticket-system/src/database"

	"github.com/scylladb/gocqlx/v2/migrate"
	"github.com/spf13/pflag"
)

var verbose = pflag.Bool("verbose", false, "output more info")

func main() {
	pflag.Parse()

	log.Println("Bootstrap database...")

	if *verbose {
		log.Printf("Configuration = %+v\n", config.Config())
	}

	createKeyspace()
	migrateKeyspace()
	printKeyspaceMetadata()
}

func createKeyspace() {
	ses, err := config.Session()
	if err != nil {
		log.Fatalln("session: ", err)
	}
	defer ses.Close()

	if err := ses.Query(database.KeyspaceCQL).Exec(); err != nil {
		log.Fatalln("ensure keyspace exists: ", err)
	}
}

func migrateKeyspace() {
	ses, err := config.Keyspace()
	if err != nil {
		log.Fatalln("session: ", err)
	}
	defer ses.Close()

	if err := migrate.Migrate(context.Background(), ses, "database/ddl"); err != nil {
		log.Fatalln("migrate: ", err)
	}
}

func printKeyspaceMetadata() {
	ses, err := config.Keyspace()
	if err != nil {
		log.Fatalln("session: ", err)
	}
	defer ses.Close()

	m, err := ses.KeyspaceMetadata(database.KeySpace)
	if err != nil {
		log.Fatalln("keyspace metadata: ", err)
	}

	log.Printf("Keyspace metadata = %+v\n", *m)
}
