package main

import (
	"bufio"
	"fmt"
	"os"
	"redis-database/src/db"
	"strings"
)

const prompt = ">> "

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

		fmt.Println(output)
	}

}
