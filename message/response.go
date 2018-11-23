package message

type Response interface {
	Message
	GetStatusCode() int
	WithStatus(code int, reasonPhrase string)
	GetReasonPhrase() string
}
