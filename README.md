# OAK

[![Documentation][godoc-img]][godoc-url]
![License][license-img]
[![Build Status][travis-img]][travis-url]
[![Coverage][codecov-img]][codecov-url]
[![Go Report Card][report-img]][report-url]

*Golang Database Manager*

[![OAK][oak-img]][oak-url]

## Overview

OAK is a package that facilitates execution of [loukoum][loukoum-url] queries
as well as migrations and scripts generate by [prana][prana-url].

## Installation

```console
$ go get -u github.com/phogolabs/oak
$ go install github.com/phogolabs/oak/cmd/oak
```

## Introduction

Note that OAK is in BETA. We may introduce breaking changes until we reach
v1.0.

Gateway API facilitates object relation mapping and query building by using
[loukoum](loukoum-url) and [sqlx][sqlx-url].

Let's first import all required packages:

```golang
import (
  lk "github.com/ulule/loukoum"
  "github.com/phogolabs/oak"
)
```

and then establish the connection:

```golang
gateway, err := oak.Open("sqlite3", "example.db")
if err != nil {
 return err
}
```

### SQL Queries

All [loukoum][loukoum-url] queries are complaint with `oak.Query` interface:

```golang
// Query returns the underlying query
type Query interface {
	// Query prepares the query
	Query() (string, []Param)
}

// NamedQuery returns the underlying query
type NamedQuery interface {
	// Query prepares the query
	NamedQuery() (string, map[string]Param)
}
```

That allows easy execution of all kind of queries.

Because the package is empowered by [sqlx][sqlx-url]. It can perform field
mapping of Golang structs by reading a `db` field tag. Let's assume that we
have the following struct:

```golang
// Package model contains an object model of database schema 'default'
// Auto-generated at Thu Apr 19 21:36:35 CEST 2018
package model

import null "gopkg.in/volatiletech/null.v6"

// User represents a data base table 'users'
type User struct {
	// ID represents a database column 'id' of type 'INT PRIMARY KEY NOT NULL'
	ID int `db:"id,primary_key,not_null" json:"id" xml:"id" validate:"required"`

	// FirstName represents a database column 'first_name' of type 'TEXT NOT NULL'
	FirstName string `db:"first_name,not_null" json:"first_name" xml:"first_name" validate:"required"`

	// LastName represents a database column 'last_name' of type 'TEXT NULL'
	LastName null.String `db:"last_name,null" json:"last_name" xml:"last_name" validate:"-"`
}
```

#### Insert a new record

```golang

query := lk.Insert("users").
	Set(
		lk.Pair("first_name", "John"),
		lk.Pair("last_name", "Doe"),
	)

if _, err := gateway.Exec(query); err != nil {
  return err
}
```

#### Select all records

```golang
query := lk.Select("id", "first_name", "last_name").From("users")
users := []User{}

if err := gateway.Select(&users, query); err != nil {
  return err
}
```

#### Select a record

```golang
query := lk.Select("id", "first_name", "last_name").
	From("users").
	Where(oak.Condition("first_name").Equal("John"))

user := User{}

if err := gateway.SelectOne(&user, query); err != nil {
  return err
}
```

You can read more details about [loukoum][loukoum-url] on their repository.

### SQL Migrations with Prana

You can execute the migration generated by [Prana][prana-url]. First, you
should load the migration directory by using [Parcello][parcello-url]. You can
load it from embedded resource or from the local directory:

```golang
if err := oak.Migrate(gateway, parcello.Dir("./database/migration")); err != nil {
	return err
}
```

### SQL Scripts and Routines with Prana

OAK provides a way to work with embeddable SQL scripts. It can understand
[SQL Scripts](https://github.com/phogolabs/prana#sql-scripts-and-commands) from
file and execute them as standard SQL queries. Let's assume that we have SQL
query named `show-sqlite-master`.

Let's first load the SQL script from file:

```golang
if err = gateway.LoadRoutinesFromReader(file); err != nil {
	log.WithError(err).Fatal("Failed to load script")
}
```

Then you can execute the desired script by just passing its name:

```golang
routine, err := gateway.Routine("show-sqlite-master")
if err != nil {
  return err
}

_, err = gateway.Exec(routine)
```

Also you can Raw SQL Scripts from your code, you should follow this
example:

```golang
rows, err := gateway.Query(oak.SQL("SELECT * FROM users WHERE id = ?", 5432))
```

If you want to execute named queries, you should use the following code snippet:

```golang
rows, err := gateway.Query(oak.NamedSQL("SELECT * FROM users WHERE id = :id", oak.P{"id": 5432}))
```

### Example

You can check our [Getting Started Example](/example).

For more information, how you can change the default behavior you can read the
help documentation by executing:

## Contributing

We are welcome to any contributions. Just fork the
[project](https://github.com/phogolabs/oak).

*logo made by [Free Pik][logo-author-url]*

[report-img]: https://goreportcard.com/badge/github.com/phogolabs/oak
[report-url]: https://goreportcard.com/report/github.com/phogolabs/oak
[logo-author-url]: https://www.freepik.com/free-photos-vectors/tree
[logo-license]: http://creativecommons.org/licenses/by/3.0/
[oak-url]: https://github.com/phogolabs/oak
[oak-img]: doc/img/logo.png
[codecov-url]: https://codecov.io/gh/phogolabs/oak
[codecov-img]: https://codecov.io/gh/phogolabs/oak/branch/master/graph/badge.svg
[travis-img]: https://travis-ci.org/phogolabs/oak.svg?branch=master
[travis-url]: https://travis-ci.org/phogolabs/oak
[oak-url]: https://github.com/phogolabs/oak
[godoc-url]: https://godoc.org/github.com/phogolabs/oak
[godoc-img]: https://godoc.org/github.com/phogolabs/oak?status.svg
[license-img]: https://img.shields.io/badge/license-MIT-blue.svg
[software-license-url]: LICENSE
[loukoum-url]: https://github.com/ulule/loukoum
[parcello-url]: https://github.com/phogolabs/parcello
[prana-url]: https://github.com/phogolabs/prana
[sqlx-url]: https://github.com/jmoiron/sqlx
