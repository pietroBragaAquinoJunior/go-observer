package internal

type Observer interface {
	Update(string)
	GetId() string
}
