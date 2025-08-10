package migrations

import (
	"perturabo/alter"
	"perturabo/create"
	"perturabo/registry"
)

var alterBlogsTable_0005 = "0005_alter_blogs_table" // Name of migration MUST be the same as the filename without .go extenstion

var UpAlterBlogsTable_0005 = registry.Register(
	registry.Action.Up,
	alterBlogsTable_0005,
	func() any {
		return &alter.Table{
			Name: "blogs",
			Body: []*alter.Column{
                //alter.Col() ... Fileds place here
			},
		}
	},
)

var DownAlterBlogsTable_0005 = registry.Register(
	registry.Action.Down,
	alterBlogsTable_0005,
	func() any {
		return &alter.Table{
			Name: "blogs",
			Body: []*alter.Column{
                //alter.Col() ... Fileds place here
			},
		}
	},
)
