// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nosarthur/dtree/db"
	"github.com/nosarthur/dtree/git"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [repo path(s)]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Add git repo paths",
	Run: func(cmd *cobra.Command, args []string) {

		var (
			name, commitMsg string
			err             error
			newRepos        []db.Repo
		)

		// load repos from DB
		repos, err := db.ReadRepos()
		if err != nil {
			fmt.Printf("failed to read repos from DB: %v", err)
			os.Exit(1)
		}
		existing := map[string]bool{}
		for _, repo := range repos {
			existing[*repo.Path] = true
		}

		for _, arg := range args {
			path, err := filepath.Abs(arg)
			if err != nil {
				fmt.Printf("cannot get absolute path for %s", arg)
				continue
			}
			if !git.IsRepo(path) {
				fmt.Printf("%s is not a git repo!\n", path)
				continue
			}
			// add new path to DB
			if !existing[path] {
				commitMsg, err = git.GetCommitMsg(path)
				if err != nil {
					fmt.Println(err)
					continue
				}
				name = filepath.Base(path)
				fmt.Printf("Repo %s added\n", name)
				record := db.Repo{
					Name: &name,
					Path: &path,
					Msg:  &commitMsg,
				}
				newRepos = append(newRepos, record)
			}
		}

		db.CreateRepos(newRepos)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
