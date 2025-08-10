package migrations

import (
	"perturabo/create"
)

func CreateUserTable_0001() *create.Table {
	return &create.Table{
		Name: "users",
		Body: []*create.Column{
			create.NewId(),
			create.NewString("name", 255).SetNullable(),
			create.NewBigInteger("value"),
		},
	}
}
