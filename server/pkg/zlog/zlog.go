package zlog

import (
	"fmt"
	"intimeServer/config"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/goark/errs"
	"github.com/rs/zerolog"
)

var logger *zerolog.Logger
var once sync.Once

func Trace(msg string) {
	zlog := setupLogger()
	zlog.Trace().Msg(msg)
}

func Debug(msg string) {
	zlog := setupLogger()
	zlog.Debug().Msg(msg)
}

func Info(msg string) {
	zlog := setupLogger()
	zlog.Info().Msg(msg)
}

func Warn(msg string) {
	zlog := setupLogger()
	zlog.Warn().Msg(msg)
}

func Error(msg string) {
	zlog := setupLogger()
	zlog.Error().Msg(msg)
}

func Fatal(msg string) {
	zlog := setupLogger()
	zlog.Fatal().Msg(msg)
}

func Panic(msg string) {
	zlog := setupLogger()
	zlog.Panic().Msg(msg)
}

func setupLogger() *zerolog.Logger {
	// 今日の日付を取得
	today := time.Now().Format("2006-01-02")
	// ファイル名に今日の日付を追加
	fileName := config.LOG_FILE_NAME + "_" + today
	// ファイルが存在するかどうかを確認
	_, err := os.Stat(fileName)

	// ファイル名に含まれる日付と現在の日付が同じか確認
	if os.IsNotExist(err) || !isSameDay(fileName) {
		// 一度だけ実行する
		once.Do(func() {
			fmt.Println("setupLogger 関数が一度だけ実行されました")

			dir := filepath.Dir(fileName)
			if err := os.MkdirAll(dir, 0755); err != nil {
				zlog := zerolog.New(os.Stdout)
				zlog.Error().Interface("error", errs.Wrap(err, errs.WithContext("dir", dir))).Msg("ログディレクトリの作成エラー")
			}

			file, err := os.Create(fileName)
			if err != nil {
				zlog := zerolog.New(os.Stdout)
				zlog.Error().Interface("error", errs.Wrap(err, errs.WithContext("file", fileName))).Msg("ログファイルの作成エラー")
			} else {
				file.Close()
			}

			openFile, openErr := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
			if openErr != nil {
				zlog := zerolog.New(os.Stdout)
				zlog.Error().Interface("error", errs.Wrap(err, errs.WithContext("file", fileName))).Msg("ログファイルのオープンエラー")
			}

			// タイムスタンプ設定
			zlog := zerolog.New(io.MultiWriter(
				openFile,
				zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false},
			)).Level(zerolog.DebugLevel).With().Timestamp().Caller().Logger()

			// 設定された状態のzerolog.Loggerを返却
			logger = &zlog
		})
	}

	return logger
}

// ファイル名に含まれる日付と現在の日付が同じ日かどうかを判定
func isSameDay(fileName string) bool {
	// ファイル名から日付を抽出
	parts := strings.Split(fileName, "_")
	if len(parts) != 2 {
		return false
	}
	// ファイル名の日付と現在の日付を比較
	return parts[1] == time.Now().Format("2006-01-02")
}
