package registry

var Funcs = map[string]func() any{}

func Register(name string, f func() any) func() any {
	Funcs[name] = f
	return f
}
