package internal

type Subject interface {
	Register(Observer)
	DeRegister(Observer)
	NotifyAll()
}
