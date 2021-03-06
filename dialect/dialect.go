// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package dialect

import (
	"context"
	"database/sql/driver"
	"io/fs"
)

// Dialect names for external usage.
const (
	MySQL    = "mysql"
	SQLite   = "sqlite3"
	Postgres = "postgres"
)

// FileSystem represents a file sytem storage
type FileSystem = fs.FS

// ExecQuerier wraps the standard Exec and Query methods.
type ExecQuerier interface {
	Execer
	Querier
}

// Execer wraps the exec database operations.
type Execer interface {
	// Exec executes a query that doesn't return rows. For example, in SQL, INSERT or UPDATE.
	// It scans the result into the pointer v. In SQL, you it's usually sql.Result.
	Exec(ctx context.Context, query string, args, v interface{}) error
}

// Querier wraps the query database operations.
type Querier interface {
	// Query executes a query that returns rows, typically a SELECT in SQL.
	// It scans the result into the pointer v. In SQL, you it's usually *sql.Rows.
	Query(ctx context.Context, query string, args, v interface{}) error
}

// Driver is the interface that wraps all necessary operations for ent clients.
type Driver interface {
	// ExecQuerier inheritance
	ExecQuerier
	// Dialect returns the dialect name of the driver.
	Dialect() string
	// Tx starts and returns a new transaction.
	// The provided context is used until the transaction is committed or rolled back.
	Tx(context.Context) (Tx, error)
	// Migrate runs the migrations
	Migrate(FileSystem) error
	// Ping sends a ping request
	Ping(context.Context) error
	// Close closes the underlying connection.
	Close() error
}

// Tx wraps the Exec and Query operations in transaction.
type Tx interface {
	// ExecQuerier inheritance
	ExecQuerier
	// actual transaction
	driver.Tx
}

type nopTx struct {
	Driver
}

func (nopTx) Commit() error   { return nil }
func (nopTx) Rollback() error { return nil }

// NopTx returns a Tx with a no-op Commit / Rollback methods wrapping
// the provided Driver d.
func NopTx(d Driver) Tx {
	return nopTx{d}
}
