### 出力先のディレクトリ
BINDIR=./bin
BIN_INTIME=$(BINDIR)/intime
### ビルド対象となるmainの格納場所
CMD_INTIME=/workspaces/cmd/intime/main.go
### ビルドコマンド
GO_BUILD=go build -o
GO_BUILD_LINUX=GOOS=linux GOARCH=amd64 go build -o
GO_BUILD_WINDOWS=GOOS=windows GOARCH=amd64 go build -o
### swagInit用
SWAG_INIT=swag init --parseDependency --parseInternal -g /cmd/intime/main.go
### ターゲットの定義
.PHONY:clean test build build-linux build-windows run setup swagInit
### セットアップで何かあれば下記についてきして
setup:build
### ALL
all:clean test build
### クリーン linuxだと動く
clean:
	rm $(BINDIR)/*
### テスト
test:
	go test -v
### build
build:build-linux
### linux用
build-linux:
	$(GO_BUILD_LINUX) $(BIN_INTIME) $(CMD_INTIME)
### windows用
####build-windows:
####	$(GO_BUILD_WINDOWS) $(BIN_INTIME).exe $(CMD_INTIME)
build-windows:
	$(GO_BUILD) $(BIN_INTIME).exe $(CMD_INTIME)
### 実行用
run:
	./bin/intime
### swagInit用 移動用はこちら	$ (cd /etc/apache2 && ls -l)
swagInit:
	$(SWAG_INIT)