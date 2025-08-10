package migrations

import (
	"perturabo/create"
	"perturabo/registry"
)

var CreateUserTable_0001 = registry.Register(
	"0001_create_user_table",
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
