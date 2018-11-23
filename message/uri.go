package message

type Uri interface {
	GetScheme() string

	GetAuthority() string

	GetUserInfo() string

	GetHost() string

	GetPort() int

	GetPath() string

	GetQuery() string

	GetFragment() string

	WithScheme(scheme string)

	WithUserInfo(user string, password string)

	WithHost(host string) string

	WithPort(port int)

	WithPath(path string)

	WithQuery(query string)

	WithFragment(fragment string)

	ToString() string
}
