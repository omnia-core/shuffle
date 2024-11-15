APP = shuffle
BRANCH = $(shell git branch --show-current)
HASH = $(shell git show --pretty=format:%h --no-patch)
TAG = $(shell git for-each-ref refs/tags --sort=-taggerdate --format='%(refname:short)' --count=1)
GIT_INFO = $(BRANCH)-$(HASH)
COMMIT = $(HASH)

.PHONY: info
info:
	@echo '$(GIT_INFO)'
	@echo '$(BRANCH)'
	@echo '$(COMMIT)'

.PHONY: build
build: info
	@go build -v -a -o bin/shuffle cmd/shuffle/main.go
