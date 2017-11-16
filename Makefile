.PHONY: image

IMAGE_NAME ?= codeclimate/codeclimate-govet

image:
	docker build --tag "$(IMAGE_NAME)" .
