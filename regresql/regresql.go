package regresql

import (
	"database/sql"
	"fmt"

	// "github.com/mndrix/tap-go"
	_ "github.com/lib/pq"
)

func Init(dir string, pguri string) {
	fmt.Println("Init: init -C %s", dir)

	fmt.Printf("Trying to connect to %s\n", pguri)
	_, err := sql.Open("postgres", pguri)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Connected")
	}
}

func Update(dir string) {
	fmt.Println("Update: update -C %s", dir)
}

func Test(dir string) {
	fmt.Println("Test: test -C %s", dir)
}

func List(dir string) {
	fmt.Println("List: list -C %s", dir)

	suite := Walk(dir)
	suite.Println()

	return
}
