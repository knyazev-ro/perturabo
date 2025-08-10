package migrations

import (
	"perturabo/create"
	"perturabo/registry"
)

var migrationNameCreateUserTable_0001 = "0001_create_user_table"

var UpCreateUserTable_0001 = registry.Register(
	registry.Action.Up,
	migrationNameCreateUserTable_0001,
	func() any {
		return &create.Table{
			Name: "users",
			Body: []*create.Column{
				create.NewId(),
				create.NewString("name", 255).SetNullable(),
				create.NewBigInteger("value"),
			},
		}
	},
)

var DownCreateUserTable_0001 = registry.Register(
	registry.Action.Down,
	migrationNameCreateUserTable_0001,
	func() any {
		return &create.Table{
			Name: "users",
			Drop: true,
		}
	},
)
