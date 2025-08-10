package main

import "perturabo/migrate"

func main() {
	// id := create.NewId().ToSQL()
	// name := create.NewString("name", 255).SetDefault("john").SetNullable().ToSQL()
	// fkey := create.NewForeignId("user_id", "users").SetCascadeOnDelete().SetCascadeOnUpdate().ToSQL()
	// fmt.Println(id)
	// fmt.Println(name)
	// fmt.Println(fkey)

	// alterT := alter.Col().AlterType("name", "VARCHAR(200)").ToSQL()
	// defaultS := alter.Col().SetDefault("number", "265").ToSQL()
	// defaultD := alter.Col().DropDefault("number").ToSQL()
	// NotNullS := alter.Col().SetNotNull("name").ToSQL()
	// NotNullD := alter.Col().DropNotNull("name").ToSQL()
	// ColRename := alter.Col().RenameColumn("name", "title").ToSQL()
	// ColRenameUsing := alter.Col().RenameColumnUsing("name", "integer").ToSQL()
	// fmt.Println(alterT)
	// fmt.Println(defaultS)
	// fmt.Println(defaultD)
	// fmt.Println(NotNullS)
	// fmt.Println(NotNullD)
	// fmt.Println(ColRename)
	// fmt.Println(ColRenameUsing)

	// table := migrations.CreateUserTable_0001()
	// mig := GenerateCreateTableSQL(table)
	// println(mig)

	// alterTable := migrations.AlterUserTable_0002()
	// mig = GenerateAlterTableSQL(alterTable)
	// println(mig)

	// println(DropTableIfExists("users"))

	migrate.Run()
}
