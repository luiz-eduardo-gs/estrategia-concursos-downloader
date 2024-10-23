build:
	go build -o ./bin/downloader ./cmd/downloader

env:
	cp .env.example .env