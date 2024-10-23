# Estratégia Concursos Downloader

This repository contains a script to download PDFs from [Estratégia Concursos](https://www.estrategiaconcursos.com.br)
using their [API](https://api.estrategiaconcursos.com.br).

## Clone the project

```
$ git clone https://github.com/luiz-eduardo-gs/estrategia-concursos-downloader
$ cd estrategia-concursos-downloader
```
## Build the project

```
$ make build 
```

This will create a bin/downloader executable file. If you're on Linux, run it with:

```
$ ./bin/downloader
```

## Run tests

```
$ make test
```

## Dependencies
* [Go](https://go.dev/) >= 1.23.2
* [make](https://www.gnu.org/software/make/) >= 4.3