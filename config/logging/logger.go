package logging


import (
	"log/slog"
	"os"
)

func init() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	logger := slog.New(logHandler).With(
		"application", "go-try",
		"env", "local",
	)
	slog.SetDefault(logger)
}


func GetLogger(name string) *slog.Logger {
	return slog.Default().With("logger", name)
}