# Go CI/CD
Repo to practice CI/CD setup for go

## Overview
Nice setup for any go repo. Goal is to minimize the number of repetative work done during development.

## Building / Testing
Done with an action to setup go, build test and upload benchmarks.

## Benchmark
Converts go test coverage into `lcov` format and displays it via github-pages.

https://iamnotaturtle.github.io/go-ci/dev/bench/

## Linting
Handled by [golangci-lint](https://golangci-lint.run/) which has multiple linters and uses `govet`.

## Auto formatting
Handled by a stricter formatter [gofumpt](https://github.com/mvdan/gofumpt) which formats code on pull requests.

## Releases
Releases work in two ways: by auto releasing the latest build and by creating a release if a tag is made.

## TODO
* [ ] deploy apps for each PR and on master
* [ ] error tracking (ex: sentry)
* [ ] logs
* [ ] security scanning (could be handled by dependabot)
