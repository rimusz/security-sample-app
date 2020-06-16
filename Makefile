GIT_HASH=$(shell git rev-parse HEAD)
BUILD_DATE=$(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

.PHONY: image
image:
	@echo "++ Building security-sample-app docker image..."
	docker build -t rimusz/security-sample-app:${GIT_HASH} .

.PHONY: psuh
push:
	@echo "++ Pushing security-sample-app docker image..."
	docker push rimusz/security-sample-app:${GIT_HASH}
