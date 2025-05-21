package database

import (
	"context"
	"github.com/ClickHouse/ch-go"
	"github.com/ClickHouse/ch-go/chpool"
	"log"
	"void-studio.net/fiesta/config"
)

var ctx = context.Background()
var pool *chpool.Pool

func init() {
	var err error = nil
	pool, err = chpool.Dial(ctx, chpool.Options{
		ClientOptions: ch.Options{
			Address:  config.Config.Database.Address,
			Database: config.Config.Database.Username,
			User:     config.Config.Database.Username,
			Password: config.Config.Database.Password,
		},
	})

	if err != nil {
		log.Fatalf("failed to connect to database %v", err)
	}

	ExecuteQuery(ch.Query{
		Body: "CREATE TABLE IF NOT EXISTS chat (player String, message String, realm String, location Bool, command Bool, private Bool, cords String, server String, time DateTime) ENGINE = MergeTree() ORDER BY time;",
	})

	ExecuteQuery(ch.Query{
		Body: "CREATE TABLE IF NOT EXISTS items (player String, item String, amount UInt8, action Bool, realm String, cords String, server String, time DateTime) ENGINE = MergeTree() ORDER BY time;",
	})

	ExecuteQuery(ch.Query{
		Body: "CREATE TABLE IF NOT EXISTS deaths (player String, inventory String, reason String, realm String, cords String, server String, time DateTime) ENGINE = MergeTree() ORDER BY time;",
	})

	ExecuteQuery(ch.Query{
		Body: "CREATE TABLE IF NOT EXISTS movement (player String, from String, to String, server String, time DateTime) ENGINE = MergeTree() ORDER BY time;",
	})

	ExecuteQuery(ch.Query{
		Body: "CREATE TABLE IF NOT EXISTS logged (player String, action Bool, realm String, cords String, server String, time DateTime) ENGINE = MergeTree() ORDER BY time;",
	})
}
