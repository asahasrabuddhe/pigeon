package pigeon

type Theme interface {
	Name() string
	HTMLTemplate() string
	PlainTextTemplate() string
}
