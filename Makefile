.PHONY: image release

IMAGE_NAME ?= codeclimate/codeclimate-govet
RELEASE_REGISTRY ?= codeclimate
RELEASE_TAG ?= latest

image:
	docker build --tag "$(IMAGE_NAME)" .

release:
	docker tag $(IMAGE_NAME) $(RELEASE_REGISTRY)/codeclimate-govet:$(RELEASE_TAG)
	docker push $(RELEASE_REGISTRY)/codeclimate-govet:$(RELEASE_TAG)
