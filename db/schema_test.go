package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setup(t *testing.T) (string, func()) {

	const testdb = "test.db"
	dir, err := ioutil.TempDir("", "")
	require.Nil(t, err, "fail to create temp directory")
	teardown := func() {
		os.RemoveAll(dir)
	}
	return filepath.Join(dir, testdb), teardown
}

func TestMustInit(t *testing.T) {
	testdb, teardown := setup(t)
	defer teardown()

	fmt.Println(testdb)
	MustInit(testdb)
	require.FileExists(t, testdb, "test DB does not exist")
	fi, err := os.Stat(testdb)
	assert.Nil(t, err, "fail to stat test DB")
	modTime := fi.ModTime()

	// do not create DB file again
	MustInit(testdb)
	require.FileExists(t, testdb, "test DB does not exist")
	fi, err = os.Stat(testdb)
	assert.Nil(t, err, "fail to stat test DB")
	assert.Equal(t, modTime, fi.ModTime(), "test DB overwritten")
}
