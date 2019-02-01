[![Build Status](https://travis-ci.org/nosarthur/dtree.svg?branch=master)](https://travis-ci.org/nosarthur/dtree)
[![codecov](https://codecov.io/gh/nosarthur/dtree/branch/master/graph/badge.svg)](https://codecov.io/gh/nosarthur/dtree)
# dtree
Dependence checker for python repositories

## sub-commands

* `add <repo-path(s)>`
* `check <repo:file>`
* `ls [repo-name]`
* `rm <repo-name(s)>`
* `update [repo-name(s)]`

where `<>` denotes required arguments and `[]` optional arguments.

## contributing

* report/fix bugs
* suggest/implement features
* review/create pull requests

For PR, make sure that

1. the code is linted, e.g.,
    * `go fmt`
    * `go vet`
    * golint: `go get -u github.com/golang/lint/golint`
    * goimports: `go get -u golang.org/x/tools/cmd/goimports`
    * errcheck: `go get -u github.com/kisielk/errcheck`

   There are editor bindings for all of them.
1. the commit message is concise and clear

## todo

* implment `dtree rm <repo-name(s)>`: delete repo in DB
* replace `exec.Command("git", ...)` by native implementations
    * [go-git](https://github.com/src-d/go-git)? also see [here](https://git-scm.com/book/en/v2/Appendix-B%3A-Embedding-Git-in-your-Applications-go-git)
    * [git2go](https://github.com/libgit2/git2go)?
* implement `dtree check [repo:file]`, using [tree package]() for visualization
* implement `dtree ls [repo-name]` to show repo statistics, e.g., number of python files
* line contributions from different authors on a file (maybe on the fly while waiting DB response for import dependencies)
* add more tests
* refactor the db code to allow parallel test
