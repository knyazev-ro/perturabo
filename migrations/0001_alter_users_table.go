package migrations

import (
	"perturabo/alter"
	"perturabo/registry"
)

var alterUsersTable_0001 = "0001_alter_users_table" // Name of migration MUST be the same as the filename without .go extenstion

var UpAlterUsersTable_0001 = registry.Register(
	registry.Action.Up,
	alterUsersTable_0001,
	func() any {
		return &alter.Table{
			Name: "users",
			Body: []*alter.Column{
				// alter.Col("name").SetDefault(alter.),
			},
		}
	},
)

var DownAlterUsersTable_0001 = registry.Register(
	registry.Action.Down,
	alterUsersTable_0001,
	func() any {
		return &alter.Table{
			Name: "users",
			Body: []*alter.Column{
				//alter.Col() ... Fileds place here
			},
		}
	},
)
