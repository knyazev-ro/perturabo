package migrations

import (
	"perturabo/alter"
	"perturabo/create"
	"perturabo/registry"
)

var AlterUserTable_0002 = registry.Register(
	"0002_alter_user_table",
	func() any {
		return &alter.Table{
			Name: "users",
			Body: []*alter.Column{
				alter.Col("name").Type("VARCHAR(253)"),
				alter.Col("first_name").Type("VARCHAR(128)"),
				alter.Col("amount").DropDefault(),
				alter.Col("amount").SetDefault("255"),
				alter.Col(nil).Add(create.NewString("last_name", 255).ToSQL()),
				alter.Col("author_id").SetNotNull(),
				alter.Col("author_id").DropNotNull(),
				alter.Col("name").Rename("full_name"),
				alter.Col("json_field").RenameTypeUsing("JSONB"),
				alter.Col("auhtor_id").Drop(),
			},
		}
	},
)
