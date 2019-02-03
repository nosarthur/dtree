package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // DB initializer
)

// Handle exposes the DB APIs
type Handle struct {
	*sqlx.DB
}

// There is redundancy in `name` and `path`. But it may be convenient.
var (
	schema = `
CREATE TABLE repo (
    name VARCHAR(31) NOT NULL,
    path VARCHAR(255) PRIMARY KEY,
    msg  VARCHAR(255) NOT NULL
);

CREATE TABLE file (
	id 	 INTEGER PRIMARY KEY,
    name VARCHAR(31) NOT NULL,
    path VARCHAR(255) NOT NULL,
	repo_name VARCHAR(31) NOT NULL,
	FOREIGN KEY(repo_name) REFERENCES repo(name) ON DELETE CASCADE
);

CREATE TABLE import (
	importer INTEGER NOT NULL,
	importee INTEGER NOT NULL,
	FOREIGN KEY(importer) REFERENCES file(id) ON DELETE CASCADE,
	FOREIGN KEY(importee) REFERENCES file(id) ON DELETE CASCADE,
	UNIQUE (importer, importee)
);
`

	insertRepo = `
INSERT INTO repo(name, path, msg)
VALUES(:name, :path, :msg)`
)

// Repo serializes the rows in repo table
type Repo struct {
	Name *string
	Path *string
	Msg  *string
}

func mustConnect(path string) Handle {
	// Sqlite has foreign key turned off by default
	src := fmt.Sprintf("file:%s?foreign_keys=on", path)
	db := sqlx.MustOpen("sqlite3", src)

	return Handle{db}
}

// MustInit create the DB file at path if it is not present, initializes it, and
// returns a DB handle. If the DB file already exists, simply return a handle.
func MustInit(path string) Handle {

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		_, err = os.Create(path)
		if err != nil {
			fmt.Printf("failed to create DB %s: %v", path, err)
			os.Exit(1)
		}
		// create tables
		h := mustConnect(path)
		h.MustExec(schema)
		fmt.Printf("DB file created at %s", path)
		return h
	} else if err != nil {
		fmt.Printf("failed to initialize DB at %s: %v\n", path, err)
		os.Exit(1)
	}
	// connect to existing DB
	return mustConnect(path)
}
