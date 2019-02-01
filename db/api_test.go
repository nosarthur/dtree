package db_test

import (
	"testing"

	"github.com/nosarthur/dtree/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepoAPI(t *testing.T) {
	testdb, teardown := setup(t)
	defer teardown()

	db.MustInit(testdb)
	// read
	repos, err := db.ReadRepos()
	require.Nil(t, err, "fail to read repos")
	assert.Empty(t, repos, "DB is not empty")

	// create
	name := "test"
	path := "test path"
	msg := "test msg"
	r := db.Repo{&name, &path, &msg}
	db.CreateRepos([]db.Repo{r})
	repos, err = db.ReadRepos()
	assert.Equal(t, 1, len(repos))
	assert.Equal(t, repos[0], r)

	// delete
	db.DeleteRepos([]string{name})
	repos, err = db.ReadRepos()
	assert.Empty(t, repos, "DB is not empty")
}
