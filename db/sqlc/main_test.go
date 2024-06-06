package sqlc

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/akoqassym/simplebank/config"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	var conf config.Config

	conf, err = config.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load the config:", err)
	}

	testDB, err = pgxpool.New(context.Background(), conf.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
