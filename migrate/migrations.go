package migrate

import (
	"fmt"
)

func Run() {
	arr, err := GetExistingMigrations()
	if err != nil {
		fmt.Println("Migration stopped. Error.", err.Error())
	}

	fmt.Println(arr)
}
