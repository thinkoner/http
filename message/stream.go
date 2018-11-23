package message

type Stream interface {
	ToString() string

	Close()

	GetSize() uint64

	Tell() int

	Eof() bool

	IsSeekable() bool

	Seek(offset int, whence int)

	Rewind()

	IsWritable() bool

	Write(str string) int

	IsReadable() bool

	Read(length int) string

	GetContents() string

	GetMetadata(key string)
}
