package migrations

import (
	"perturabo/common"
	"perturabo/create"
	"perturabo/registry"
)

var createGerardMigrationsTable_0000 = "0000_create_gerard_migrations_table" // Name of migration MUST be the same as the filename without .go extenstion

var UpCreateGerardMigrationsTable_0000 = registry.Register(
	registry.Action.Up,
	createGerardMigrationsTable_0000,
	func() any {
		return &create.Table{
			Name: "gerard_migrations",
			Body: []*create.Column{
				create.NewId(),
				create.NewString("name", 255),
				create.NewBigInteger("wave_id"),

				create.NewTimestamp("updated_at").Default(common.Now()),
				create.NewTimestamp("created_at").Default(common.Now()),
			},
		}
	},
)

var DownCreateGerardMigrationsTable_0000 = registry.Register(
	registry.Action.Down,
	createGerardMigrationsTable_0000,
	func() any {
		return &create.Table{
			Name: "gerard_migrations",
			Drop: true,
		}
	},
)
