# dev_setup_sample_ts_go

GoとTSでプロダクトを作る場合のサンプル

## Setup

### air（hot reload）

```
cd backend
go install github.com/air-verse/air@latest
air init
air
```

## Linter

### golangci-lint

* https://gist.github.com/maratori/47a4d00457a92aa426dbd48a18776322#file-golangci-yml より、`/backend/.golangci.yml`を作成

```sh
cd backend
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b ./bin v2.6.2
golangci-lint run
```
