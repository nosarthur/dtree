package db

import "fmt"

// CreateRepos writes to DB
func (h Handle) CreateRepos(rs []Repo) {
	for _, repo := range rs {
		_, err := h.NamedExec(insertRepo, repo)
		if err != nil {
			fmt.Printf("fail to add repo %s to DB: %v\n", *repo.Name, err)
		}
	}
}

// ReadRepos returns all repos in DB
func (h Handle) ReadRepos() ([]Repo, error) {
	repos := []Repo{}
	err := h.Select(&repos, "SELECT * FROM repo ORDER BY name ASC")
	if err != nil {
		return nil, fmt.Errorf("fail to read repos from DB: %v", err)
	}
	return repos, nil
}

// DeleteRepos deletes the selected repos
func (h Handle) DeleteRepos(names []string) {
	for _, name := range names {
		_, err := h.Exec("DELETE FROM repo WHERE name=$1", name)
		if err != nil {
			fmt.Printf("fail to delete repo %s: %v", name, err)
			continue
		}
	}
}
