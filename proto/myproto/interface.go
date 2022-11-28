package myproto

type MyMsg interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}
