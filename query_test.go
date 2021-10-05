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
	_ "github.com/lib/pq"

	querytest "github.com/marcusirgens/sqlc-interval-debug/go"
	"pkg.iterate.no/pgutil"
)

var dbUrl = flag.String("database-url", "", "Database URL")

func TestQueries(t *testing.T) {
	var drivers = []string{"postgres", "pgx"}
	for _, driver := range drivers {
		t.Run(driver, func(t *testing.T) {
			testQueries(t, driver)
		})
	}
}

// testQueries is TestQueries wth different drivers.
func testQueries(t *testing.T, driver string) {
	// Bootstrapping
	t.Logf("connecting to %s", *dbUrl)
	db, err := sql.Open(driver, *dbUrl)
	if err != nil {
		t.Fatalf("connecting to database: %v", err)
	}

	// create a context with a deadline if the test has a deadline.
	ctx := context.Background()
	if ddl, hasDeadline := t.Deadline(); hasDeadline {
		// sub a second to give the pinger some time to cancel
		c, ccl := context.WithDeadline(ctx, ddl.Add(-time.Second))
		ctx = c
		defer ccl()
	}

	// wait for the database to wake up.
	if err := pgutil.Wait(ctx, db); err != nil {
		t.Fatalf("connecting to database: %v", err)
	}

	// Actual testing starts here.
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
			// Create new rows in the Foo table using sqlc.
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
		// List all the rows in the Foo table using sqlc.
		fs, err := q.Get(ctx, math.MaxInt32)
		if err != nil {
			t.Errorf("fetching foos: %v", err)
		}
		_ = fs
	})
}
