install:
	# Golang CI Linter
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.33.0
	golangci-lint --version

	# Wails
	sudo pacman -S gcc pkgconf webkit2gtk gtk3
