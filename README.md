<h1 align="center">Welcome to dev_setup_sample_ts_go 👋</h1>
<p>
  <a href="https://twitter.com/taichi\_kosgiii" target="_blank">
    <img alt="Twitter: taichi\_kosgiii" src="https://img.shields.io/twitter/follow/taichi\_kosgiii.svg?style=social" />
  </a>
</p>

> TypeScriptとGoで新規プロダクトを作成する時のテンプレ

## Author

👤 **Taichi Kosugi**

* Twitter: [@taichi\_kosgiii](https://twitter.com/taichi\_kosgiii)
* Github: [@taichi-k](https://github.com/taichi-k)

## Show your support

Give a ⭐️ if this project helped you!

## Setup

### go

Goをインストールしてください。

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



***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
