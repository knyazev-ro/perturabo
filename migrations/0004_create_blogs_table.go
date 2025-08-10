package migrations

import (
	"perturabo/create"
	"perturabo/registry"
)

var createBlogsTable_0004 = "0004_create_blogs_table" // Name of migration MUST be the same as the filename without .go extenstion

var UpCreateBlogsTable_0004 = registry.Register(
	registry.Action.Up,
	createBlogsTable_0004,
	func() any {
		return &create.Table{
			Name: "blogs",
			Body: []*create.Column{
            	create.NewId(),
                // Fields place here
				create.NewTimestamp("updated_at").SetDefault("now()"),
				create.NewTimestamp("created_at").SetDefault("now()"),
			},
		}
	},
)

var DownCreateBlogsTable_0004 = registry.Register(
	registry.Action.Down,
	createBlogsTable_0004,
	func() any {
		return &create.Table{
			Name: "blogs",
			Drop: true,
		}
	},
)
