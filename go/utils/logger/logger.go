package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var currentLogFile *os.File

func getLogFilePath() string {
	// 現在の日付でファイル名を生成
	now := time.Now()
	fileName := fmt.Sprintf("%d-%02d-%02d.log", now.Year(), now.Month(), now.Day())
	return filepath.Join("logs", fileName)
}

func openLogFile() (*os.File, error) {
	// logsディレクトリが存在しない場合は作成
	if err := os.MkdirAll("logs", 0755); err != nil {
		return nil, fmt.Errorf("ログディレクトリの作成に失敗しました: %v", err)
	}

	logPath := getLogFilePath()
	return os.OpenFile(
		logPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
}

func Init() {
	var err error
	currentLogFile, err = openLogFile()
	if err != nil {
		log.Fatal().Err(err).Msg("ログファイルを開けませんでした")
	}

	// 複数の出力先を設定（コンソールとファイル）
	multi := zerolog.MultiLevelWriter(os.Stdout, currentLogFile)

	// グローバルロガーの設定
	log.Logger = zerolog.New(multi).
		Output(zerolog.ConsoleWriter{
			Out:        multi,
			TimeFormat: time.RFC3339,
			NoColor:    false,
			FormatMessage: func(i interface{}) string {
				if str, ok := i.(string); ok {
					var prettyJSON bytes.Buffer
					if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err == nil {
						return "\n" + prettyJSON.String()
					}
				}
				return fmt.Sprintf("%v", i)
			},
		}).
		With().
		Timestamp().
		Caller().
		Logger()

	// タイムスタンプのフォーマットを設定
	zerolog.TimeFieldFormat = time.RFC3339

	// 日付が変わった時にログファイルを切り替えるゴルーチンを開始
	go rotateLogFile()
}

func rotateLogFile() {
	for {
		now := time.Now()
		// 次の日の0時0分0秒までの時間を計算
		next := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
		duration := next.Sub(now)

		// 次の日までスリープ
		time.Sleep(duration)

		// 古いファイルを閉じる
		if currentLogFile != nil {
			currentLogFile.Close()
		}

		// 新しいファイルを開く
		var err error
		currentLogFile, err = openLogFile()
		if err != nil {
			log.Error().Err(err).Msg("新しいログファイルの作成に失敗しました")
			continue
		}

		// ロガーを再設定
		multi := zerolog.MultiLevelWriter(os.Stdout, currentLogFile)
		log.Logger = zerolog.New(multi).
			Output(zerolog.ConsoleWriter{
				Out:        multi,
				TimeFormat: time.RFC3339,
				NoColor:    false,
				FormatMessage: func(i interface{}) string {
					if str, ok := i.(string); ok {
						var prettyJSON bytes.Buffer
						if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err == nil {
							return "\n" + prettyJSON.String()
						}
					}
					return fmt.Sprintf("%v", i)
				},
			}).
			With().
			Timestamp().
			Caller().
			Logger()
	}
}

// エラーログを記録する関数
func LogError(err error, message string) {
	log.Error().
		Err(err).
		Msg(message)
}

// 警告ログを記録する関数
func LogWarn(message string) {
	log.Warn().
		Msg(message)
}

// 情報ログを記録する関数
func LogInfo(message string) {
	log.Info().
		Msg(message)
}

// 情報ログを記録する関数
func LogInfoObject(object interface{}) {
	prettyJSON, err := json.MarshalIndent(object, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("オブジェクトのJSON変換に失敗しました")
		return
	}
	log.Info().
		RawJSON("object", prettyJSON).
		Msg("info")
}

// FormatGinLog Ginのログをフォーマットする関数
func FormatGinLog(param gin.LogFormatterParams) string {
	// エラーが発生した場合は詳細情報を含める
	var errorDetail string
	if param.ErrorMessage != "" {
		errorDetail = fmt.Sprintf(" | Error: %s", param.ErrorMessage)
	}

	// リクエストパラメータをJSONに変換
	params := map[string]interface{}{
		"client_ip":  param.ClientIP,
		"method":     param.Method,
		"path":       param.Path,
		"protocol":   param.Request.Proto,
		"status":     param.StatusCode,
		"latency":    param.Latency.String(),
		"user_agent": param.Request.UserAgent(),
		"error":      param.ErrorMessage,
		"time":       param.TimeStamp.Format("2006-01-02 15:04:05"),
		"headers":    param.Request.Header,
	}

	// クエリパラメータがある場合は追加
	if len(param.Request.URL.RawQuery) > 0 {
		params["query"] = param.Request.URL.RawQuery
	}

	// JSONに変換
	jsonBytes, err := json.MarshalIndent(params, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal request params")
		return fmt.Sprintf("[GIN] %v | %3d | %13v | %15s | %-7s %s%s\n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			errorDetail,
		)
	}

	return fmt.Sprintf("[GIN] %s\n", string(jsonBytes))
}
