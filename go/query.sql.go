// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package querytest

import (
	"context"
)

const get = `-- name: Get :many
SELECT bar, "interval" FROM foo LIMIT $1
`

func (q *Queries) Get(ctx context.Context, limit int32) ([]Foo, error) {
	rows, err := q.db.QueryContext(ctx, get, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Foo
	for rows.Next() {
		var i Foo
		if err := rows.Scan(&i.Bar, &i.Interval); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insert = `-- name: Insert :exec
INSERT INTO foo (bar, interval) VALUES ($1, $2)
`

type InsertParams struct {
	Bar      bool
	Interval int64
}

func (q *Queries) Insert(ctx context.Context, arg InsertParams) error {
	_, err := q.db.ExecContext(ctx, insert, arg.Bar, arg.Interval)
	return err
}
