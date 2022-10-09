### 出力先のディレクトリ
BINDIR=./bin
BIN_INTIME=$(BINDIR)/intime
### ビルド対象となるmainの格納場所
CMD_INTIME=./cmd/intime/main.go
### ビルドコマンド
GO_BUILD=go build -o
GO_BUILD_LINUX=GOOS=linux GOARCH=amd64 go build -o
GO_BUILD_WINDOWS=GOOS=windows GOARCH=amd64 go build -o
### ターゲットの定義
.PHONY:clean test build build-linux build-windows run
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

run:
	./bin/intime