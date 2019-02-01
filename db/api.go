package db

import "fmt"

// CreateRepos writes to DB
func CreateRepos(rs []Repo) {
	for _, repo := range rs {
		_, err := conn.NamedExec(insertRepo, repo)
		if err != nil {
			fmt.Printf("fail to add repo %s to DB: %v\n", *repo.Name, err)
		}
	}
}

// ReadRepos returns all repos in DB
func ReadRepos() ([]Repo, error) {
	repos := []Repo{}
	err := conn.Select(&repos, "SELECT * FROM repo ORDER BY name ASC")
	if err != nil {
		return nil, fmt.Errorf("fail to read repos from DB: %v", err)
	}
	return repos, nil
}

// DeleteRepos deletes the selected repos
func DeleteRepos(paths []string) {
	for _, path := range paths {
		_, err := conn.Exec("DELETE FROM repo WHERE path=$1", path)
		if err != nil {
			fmt.Printf("fail to delete repo %s: %v", path, err)
			continue
		}
	}
}
