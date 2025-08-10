package migrations

import (
	"perturabo/alter"
	"perturabo/create"
	"perturabo/registry"
)

var migrationNameAlterUserTable_0002 = "0002_alter_user_table"

var UpAlterUserTable_0002 = registry.Register(
	registry.Action.Up,
	migrationNameAlterUserTable_0002,
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

var DownAlterUserTable_0002 = registry.Register(
	registry.Action.Down,
	migrationNameAlterUserTable_0002,
	func() any {
		return &alter.Table{
			Name: "users",
			Body: []*alter.Column{
				alter.Col("name").Type("VARCHAR(255)"),
				alter.Col("first_name").Type("VARCHAR(255)"),
				alter.Col("amount").SetDefault("255"),
				alter.Col("amount").DropDefault(),
				alter.Col("last_name").Drop(),
				alter.Col("author_id").DropNotNull(),
				alter.Col("author_id").SetNotNull(),
				alter.Col("full_name").Rename("name"),
				alter.Col("json_field").RenameTypeUsing("TEXT"),
				alter.Col(nil).Add(create.NewBigInteger("author_id").ToSQL()),
			},
		}
	},
)
