package commands

import "perturabo/migrate"

func Rollback(args []string) {
	migrate.Down()
}
