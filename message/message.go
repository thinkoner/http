package message

type Message interface {
	GetProtocolVersion() string
	WithProtocolVersion(version string)
	GetHeaders() map[string][]string
	HasHeader(name string) bool
	GetHeader(name string) []string
	GetHeaderLine(name string) string
	WithHeader(name string, value interface{})
	WithAddedHeader(name string, value interface{})
	WithoutHeader(name string)
	GetBody() Stream
	WithBody(body Stream)
}
