# inTime
# dev test write
基本的にdevブランチから切ってdevブランチにぶち込んで下さい。

devブランチをマスターブランチにぶち込むときには皆が承認後マージするようにしてください。

# 各種サービスコンテナの起動

```shell
docker-compose up -d --build
```

## Client

コンテナ起動時にアプリも同時に立ち上がります。

下記 URL よりアクセスしてください。

- http://localhost

## Server

ビルド & サーバ起動

```shell
docker-compose exec server make build run
```