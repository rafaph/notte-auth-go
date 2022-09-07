include ../../Makefile
DEV_DIR=../..
SERVICE=auth

.PHONY: build
build:
	$(call compose,build --no-cache && $(call down))

.PHONY: up
up:
	$(call compose,up auth && $(call down))

.PHONY: shell
shell:
	$(call run,bash)
