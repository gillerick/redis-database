package main

import (
	"bufio"
	"fmt"
	"os"
	"redis-database/src/db"
	"strings"
)

const prompt = ">> "
const commands = `GET key: Prints value for a given key and NULL if value not in database
	 SET key value: Sets database key to the specified value. Returns 'OK' if operation successful
     DEL key: Deletes a key-value pair from the database
     COUNT value: Prints the number of keys that have the given value assigned to them. Prints 0 if value not assigned anywhere
	 EXISTS key: Prints 1 if value with key exists, otherwise prints 0
     INCR key [increment]: Increments a value by increment (if provided), otherwise by 1
     DECR key [decrement]: Decrements a value by decrement (if provided), otherwise by 1
     GETSET key value: Sets the string value of a key and return its old value                                 
	 `

func main() {
	store := make(db.Store)
	reversedStore := make(db.ReversedStore)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(prompt)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := strings.TrimSpace(scanner.Text())
		output, err := db.EvaluateCommand(line, &store, &reversedStore)

		//Handle any error that might occur
		if err != nil {
			fmt.Println(err.Error())
		}

		if output == "HELP" {
			fmt.Println(commands)
		}

		fmt.Println(output)
	}

}
