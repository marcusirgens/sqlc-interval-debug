//go:build withdb

package main

import (
	"context"
	"database/sql"
	"flag"
	"math"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"pkg.iterate.no/pgutil"

	querytest "github.com/marcusirgens/sqlc-interval-debug/go"
)

var dbUrl = flag.String("database-url", "", "Database URL")

func TestQueries(t *testing.T) {
	t.Logf("connecting to %s", *dbUrl)
	db, err := sql.Open("pgx", *dbUrl)
	if err != nil {
		t.Fatalf("connecting to database: %v", err)
	}
	ctx := context.Background()
	if ddl, hasDeadline := t.Deadline(); hasDeadline {
		// sub a second to give the pinger some time to cancel
		c, ccl := context.WithDeadline(ctx, ddl.Add(-time.Second))
		ctx = c
		defer ccl()
	}
	if err := pgutil.Wait(ctx, db); err != nil {
		t.Fatalf("connecting to database: %v", err)
	}

	q := querytest.New(db)

	t.Run("Insert various Foos", func(t *testing.T) {
		durs := []time.Duration{
			time.Millisecond * 1,
			time.Millisecond * 15,
			time.Second * 1,
			time.Second * 15,
			time.Minute * 1,
			time.Minute * 15,
			time.Hour * 1,
			time.Hour * 15,
		}
		for _, d := range durs {
			err := q.Insert(ctx, querytest.InsertParams{
				Bar:      true,
				Interval: int64(d),
			})
			if err != nil {
				t.Errorf("inserting Foo: %v", err)
			}
		}
	})

	t.Run("List all Foos", func(t *testing.T) {
		fs, err := q.Get(ctx, math.MaxInt32)
		if err != nil {
			t.Errorf("fetching foos: %v", err)
		}
		_ = fs
	})
}