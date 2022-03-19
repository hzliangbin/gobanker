package blog

type Config struct {
	LogFile string
	MaxSize int
	MaxBackups int
	MaxAge int
	Compress bool

	LogLevel string
	JsonEncode bool
	StacktraceLevel string
}