package database

const KeySpace = "go"

const KeyspaceCQL = "CREATE KEYSPACE IF NOT EXISTS go WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'}  AND durable_writes = true;"
