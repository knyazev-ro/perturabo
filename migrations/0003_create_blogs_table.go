package migrations

import (
	"perturabo/create"
	"perturabo/registry"
)

var createBlogsTable_0003 = "0003_create_blogs_table" // Name of migration MUST be the same as the filename without .go extenstion

var UpCreateBlogsTable_0003 = registry.Register(
	registry.Action.Up,
	createBlogsTable_0003,
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

var DownCreateBlogsTable_0003 = registry.Register(
	registry.Action.Down,
	createBlogsTable_0003,
	func() any {
		return &create.Table{
			Name: "blogs",
			Drop: true,
		}
	},
)
