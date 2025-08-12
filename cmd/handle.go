package cmd

import "flag"

func Handle(args []string) {

	Init()

	command := args[1]
	if command == "help" {
		// GetHelp()
		return
	}

	flagSet := flag.NewFlagSet("args", flag.ContinueOnError)
	flagSet.Parse(args[2:])
	argsArr := flagSet.Args()

	switch command {
	case "create:migration":
		CreateMigration(argsArr)
	case "alter:migration":
		AlterMigration(argsArr)
	case "migrate:run":
		Migrate(argsArr)
	case "migrate:rollback":
		Rollback(argsArr)
	default:
		println("Error. Unknown command ", command)
		return
	}

}
