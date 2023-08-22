package adapter

type IAdapter interface {
	Init() error
}

func RunAdapter(adapter IAdapter) {
	adapter.Init()
}
