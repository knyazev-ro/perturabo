package sysmigrations

import (
	"perturabo/create"
	"perturabo/registry"
)

var migrationNameCreatePerturaboMigrationsTable_0000 = "0000_create_perturabo_migrations_table"

var UpCreatePerturaboMigrationsTable_0000 = registry.Register(
	registry.Action.Up,
	migrationNameCreatePerturaboMigrationsTable_0000,
	func() any {
		return &create.Table{
			Name: "perturabo_migrations",
			Body: []*create.Column{
				create.NewId(),
				//
			},
		}
	},
)

var DownCreatePerturaboMigrationsTable_0000 = registry.Register(
	registry.Action.Down,
	migrationNameCreatePerturaboMigrationsTable_0000,
	func() any {
		return &create.Table{
			Name: "perturabo_migrations",
			Drop: true,
		}
	},
)
