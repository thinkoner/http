package message

type Request interface {
	Message

	GetRequestTarget() string

	WithRequestTarget(interface{})

	GetMethod() string

	WithMethod(message string)

	GetUri() Uri

	WithUri(uri Uri, preserveHost bool)
}
