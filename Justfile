PROJECT_NAME := "AlphaNum"
PROJECT_CLI := "alphanum"

alias arc := archive

set dotenv-load := false


@_default: _term-wipe
	just --list


# Archive GoReleaser dist
archive: _term-wipe
	#!/bin/sh
	tag="$(git tag --points-at main)"
	app="{{PROJECT_NAME}}"
	arc="${app}_${tag}"

	echo "app = '${app}'"
	echo "tag = '${tag}'"
	echo "arc = '${arc}'"
	if [ ! -e _dist ]; then
		mkdir _dist
	fi
	if [ -e dist ]; then
		echo "Move dist -> _dist/${arc}"
		# mv dist "_dist/${arc}"

		# echo "cd distro"
		# cd distro
		
		printf "pwd = "
		pwd
		
		ls -Alh
	else
		echo "dist directory not found for archiving"
	fi


# Build app
build: _term-wipe
	GOOS=darwin GOARCH=amd64 go build -o bin/macos/{{PROJECT_CLI}} ./cmd/{{PROJECT_CLI}}/{{PROJECT_CLI}}.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux/{{PROJECT_CLI}} ./cmd/{{PROJECT_CLI}}/{{PROJECT_CLI}}.go
	GOOS=windows GOARCH=amd64 go build -o bin/windows/{{PROJECT_CLI}}.exe ./cmd/{{PROJECT_CLI}}/{{PROJECT_CLI}}.go
	@# if [ '{{os()}}' = 'windows' ]; then
	@# 	go build -o bin/windows/{{PROJECT_CLI}}.exe ./cmd/{{PROJECT_CLI}}/{{PROJECT_CLI}}.go
	@# else
	@# 	go build -o bin/{{os()}}/{{PROJECT_CLI}} ./cmd/{{PROJECT_CLI}}/main.go
	ls -al bin/*/*


# Clean up this place!
clean: _term-wipe
	rm c.out


# Build distro
distro: _term-wipe
	#!/bin/sh
	# goreleaser
	just archive


# Build and install the app
install: _term-wipe
	#!/bin/sh
	cd cmd/{{PROJECT_CLI}}
	go install


# Run code
run +args='': _term-wipe
	go run ./cmd/{{PROJECT_CLI}}/{{PROJECT_CLI}}.go {{args}}


# Run a test
@test cmd="package": _term-wipe
	just test-{{cmd}}

# Run Go Test Coverage
@test-coverage:
	echo "You need to run:"
	echo "go test -coverprofile=c.out"
	echo "go tool cover -func=c.out"

# Run Go Unit Tests
test-package:
	go test


_term-wipe:
	#!/usr/bin/env bash
	set -exo pipefail
	if [[ ${#VISUAL_STUDIO_CODE} -gt 0 ]]; then
		clear
	elif [[ ${KITTY_WINDOW_ID} -gt 0 ]] || [[ ${#TMUX} -gt 0 ]] || [[ "${TERM_PROGRAM}" = 'vscode' ]]; then
		printf '\033c'
	elif [[ "${TERM_PROGRAM}" = 'Apple_Terminal' ]] || [[ "${TERM_PROGRAM}" = 'iTerm.app' ]]; then
		osascript -e 'tell application "System Events" to keystroke "k" using command down'
	elif [[ -x "$(which tput)" ]]; then
		tput reset
	elif [[ -x "$(which tcap)" ]]; then
		tcap rs
	elif [[ -x "$(which reset)" ]]; then
		reset
	else
		clear
	fi

