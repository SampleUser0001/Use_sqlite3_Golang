# sqliteにつなぐ

- [sqliteにつなぐ](#sqliteにつなぐ)
  - [init](#init)
  - [実行](#実行)
  - [実行結果](#実行結果)
  - [参考](#参考)

## init

``` bash
cd app
go mod init app
go get github.com/mattn/go-sqlite3
```

## 実行

``` bash
go run app.go
```

## 実行結果

``` txt
Printing all:
Id: 6, Name: init data
Printing all: finish.

Inserted test

Printing all:
Id: 6, Name: init data
Id: 7, Name: test
Printing all: finish.
```

## 参考

- [Go言語でSQLite3を使う](https://zenn.dev/teasy/articles/go-sqlite3-sample)
- [https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go](https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go)
- [【Golang】データベース操作 sqlite3:Qiita](https://qiita.com/geniusmaaakun/items/b3cb44de3b10a526be98)