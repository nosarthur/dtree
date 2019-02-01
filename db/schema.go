package db

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	homedir "github.com/mitchellh/go-homedir"
)

// There is redundancy in `name` and `path`. But it may be convenient.
var (
	conn   *sqlx.DB // global connection singleton
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

func mustConnect(path string) *sqlx.DB {
	if path == "" {
		path = getDBPath()
	}
	// Foreign key is off by default for sqlite
	src := fmt.Sprintf("file:%s?foreign_keys=on", path)
	db := sqlx.MustOpen("sqlite3", src)

	return db
}

// MustInit initializes the DB if it is not present, then sets the global DB connection.
func MustInit(path string) {
	if path == "" {
		path = getDBPath()
	}
	_, err := os.Stat(path)
	if err == nil {
		// connect to existing DB
		conn = mustConnect(path)
	} else if os.IsNotExist(err) {
		_, err = os.Create(path)
		if err != nil {
			fmt.Printf("failed to create DB %s: %v", path, err)
			os.Exit(1)
		}
		// create tables
		conn = mustConnect(path)
		conn.MustExec(schema)
		fmt.Printf("DB file created at %s", path)
	} else {
		fmt.Printf("failed to initialize DB at %s: %v\n", path, err)
		os.Exit(1)
	}
}

// getDBPath returns the location of the DB file
func getDBPath() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Printf("failed to get home directory: %v", err)
		os.Exit(1)
	}
	return filepath.Join(home, "./dtree.db")
}
