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
	"strings"

	"github.com/nosarthur/dtree/git"
	"github.com/nosarthur/tree"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls [repo-name]",
	Args:  cobra.MaximumNArgs(1),
	Short: "Show information of all repos, or statistics of a single repo.",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			reposToDelete, hashAndMsg []string
		)

		dbHandle := getDBHandle()
		repos, err := dbHandle.ReadRepos()
		if err != nil {
			fmt.Printf("failed to read repos from DB: %v", err)
			os.Exit(1)
		}
		for _, repo := range repos {
			if !git.IsRepo(*repo.Path) {
				reposToDelete = append(reposToDelete, *repo.Path)
				continue
			}
			hashAndMsg = strings.SplitAfterN(*repo.Msg, " ", 2)
			fmt.Printf("%-18s %s%s", *repo.Name, tree.Colorize(hashAndMsg[0], tree.Blue), hashAndMsg[1])
		}
		if len(reposToDelete) > 0 {
			dbHandle.DeleteRepos(reposToDelete)
		}

	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
