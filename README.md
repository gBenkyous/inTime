# inTime
# dev test write
基本的にdevブランチから切ってdevブランチにぶち込んで下さい。

devブランチをマスターブランチにぶち込むときには皆が承認後マージするようにしてください。

# 実行コマンド

各種サービスの起動 

```
# 開発環境
docker-compose --profile dev up -d --build

# 本番環境
docker-compose --profile prod up -d --build
```

Server の起動

```
docker-compose exec server-dev make build run
```