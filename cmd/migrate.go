package cmd

import "perturabo/migrate"

func Migrate(args []string) {
	migrate.Up()
}
