package sqlc

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

const dbDriver = "postgres"
const dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
