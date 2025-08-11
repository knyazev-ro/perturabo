package commands

import "perturabo/migrate"

func Migrate(args []string) {
	migrate.Up()
}
