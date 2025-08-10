package registry

type MigrationAction struct {
	Up   string
	Down string
}

var Action = &MigrationAction{
	Up:   "up",
	Down: "down",
}

var Up = map[string]func() any{}
var Down = map[string]func() any{}

func Register(action string, name string, f func() any) func() any {
	if action == Action.Up {
		Up[name] = f
	} else {
		Down[name] = f
	}
	return f
}
