package module

type Module interface {
	Start(addr string)
	Quit()
}
