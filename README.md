# Go CI/CD
Repo to practice CI/CD setup for go

## Overview
Nice setup for any go repo. We have these actions so far:
* linting via [golangci-lint](https://golangci-lint.run/)
* auto formatting via [gofumpt](https://github.com/mvdan/gofumpt)
* building/testing
* benchmarking

## Linting
Handled by `golangci-lint` which has multiple linters and uses `govet`.

## Auto formatting
Handled by a stricter formatter `gofumpt` which formats code on pull requests.

## Benchmark
Converts go test coverage into `lcov` format and displays it via github-pages.

https://iamnotaturtle.github.io/go-ci/dev/bench/

## TODO
* [ ] releases
* [ ] deploy apps for each PR and on master
* [ ] error tracking (ex: sentry)
* [ ] logs
