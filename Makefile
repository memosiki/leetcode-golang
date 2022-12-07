.PHONY: all githooks pretty gen-readme check-readme
all:

githooks:
	cp hooks/* .git/hooks

pretty:
	go fmt

gen-readme readme:
	go run generateReadme.go > README.md

check-readme:
	go run generateReadme.go --check-only
