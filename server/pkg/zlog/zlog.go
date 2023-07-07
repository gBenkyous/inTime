package zlog

import (
	config "intimeServer/config"
	"io"
	"os"

	"github.com/goark/errs"
	"github.com/rs/zerolog"
)

func Trace(msg string) {
	// ファイル出力先指定
	zlog := zerolog.Nop()
	file, err := os.OpenFile(config.LOG_FILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		zlog = zerolog.New(os.Stdout)
	} else {
		zlog = zerolog.New(io.MultiWriter(
			file,
			zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false},
		))
	}
	// タイムスタンプ設定
	zlog = zlog.Level(zerolog.DebugLevel).With().Timestamp().Caller().Logger()
	if err != nil {
		zlog.Error().Interface("error", errs.Wrap(err, errs.WithContext("file", file))).Str("file", config.LOG_FILE_NAME).Msg("error in opening logfile")
	}
	zlog.Trace().Msg(msg)
}

func Debug(msg string) {
	// ファイル出力先指定
	zlog := zerolog.Nop()
	file, err := os.OpenFile(config.LOG_FILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		zlog = zerolog.New(os.Stdout)
	} else {
		zlog = zerolog.New(io.MultiWriter(
			file,
			zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false},
		))
	}
	// タイムスタンプ設定
	zlog = zlog.Level(zerolog.DebugLevel).With().Timestamp().Caller().Logger()
	if err != nil {
		zlog.Error().Interface("error", errs.Wrap(err, errs.WithContext("file", file))).Str("file", config.LOG_FILE_NAME).Msg("error in opening logfile")
	}
	zlog.Debug().Msg(msg)
}

func Info(msg string) {
	// ファイル出力先指定
	zlog := zerolog.Nop()
	file, err := os.OpenFile(config.LOG_FILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		zlog = zerolog.New(os.Stdout)
	} else {
		zlog = zerolog.New(io.MultiWriter(
			file,
			zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false},
		))
	}
	// タイムスタンプ設定
	zlog = zlog.Level(zerolog.DebugLevel).With().Timestamp().Caller().Logger()
	if err != nil {
		zlog.Error().Interface("error", errs.Wrap(err, errs.WithContext("file", file))).Str("file", config.LOG_FILE_NAME).Msg("error in opening logfile")
	}
	zlog.Info().Msg(msg)
}

func Warn(msg string) {
	// ファイル出力先指定
	zlog := zerolog.Nop()
	file, err := os.OpenFile(config.LOG_FILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		zlog = zerolog.New(os.Stdout)
	} else {
		zlog = zerolog.New(io.MultiWriter(
			file,
			zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false},
		))
	}
	// タイムスタンプ設定
	zlog = zlog.Level(zerolog.DebugLevel).With().Timestamp().Caller().Logger()
	if err != nil {
		zlog.Error().Interface("error", errs.Wrap(err, errs.WithContext("file", file))).Str("file", config.LOG_FILE_NAME).Msg("error in opening logfile")
	}
	zlog.Warn().Msg(msg)
}

func Error(msg string) {
	// ファイル出力先指定
	zlog := zerolog.Nop()
	file, err := os.OpenFile(config.LOG_FILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		zlog = zerolog.New(os.Stdout)
	} else {
		zlog = zerolog.New(io.MultiWriter(
			file,
			zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false},
		))
	}
	// タイムスタンプ設定
	zlog = zlog.Level(zerolog.DebugLevel).With().Timestamp().Caller().Logger()
	if err != nil {
		zlog.Error().Interface("error", errs.Wrap(err, errs.WithContext("file", file))).Str("file", config.LOG_FILE_NAME).Msg("error in opening logfile")
	}
	zlog.Error().Stack().Err(err).Msg(msg)
}

func Fatal(msg string) {
	// ファイル出力先指定
	zlog := zerolog.Nop()
	file, err := os.OpenFile(config.LOG_FILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		zlog = zerolog.New(os.Stdout)
	} else {
		zlog = zerolog.New(io.MultiWriter(
			file,
			zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false},
		))
	}
	// タイムスタンプ設定
	zlog = zlog.Level(zerolog.DebugLevel).With().Timestamp().Caller().Logger()
	if err != nil {
		zlog.Error().Interface("error", errs.Wrap(err, errs.WithContext("file", file))).Str("file", config.LOG_FILE_NAME).Msg("error in opening logfile")
	}
	zlog.Fatal().Msg(msg)
}

func Panic(msg string) {
	// ファイル出力先指定
	zlog := zerolog.Nop()
	file, err := os.OpenFile(config.LOG_FILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		zlog = zerolog.New(os.Stdout)
	} else {
		zlog = zerolog.New(io.MultiWriter(
			file,
			zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false},
		))
	}
	// タイムスタンプ設定
	zlog = zlog.Level(zerolog.DebugLevel).With().Timestamp().Caller().Logger()
	if err != nil {
		zlog.Error().Interface("error", errs.Wrap(err, errs.WithContext("file", file))).Str("file", config.LOG_FILE_NAME).Msg("error in opening logfile")
	}
	zlog.Panic().Msg(msg)
}
