package main

import (
	"encoding/json"
	"flag"
	"os"
	"time"

	"github.com/spacejunkrpg/character"
)

func main() {
	charCmd := flag.NewFlagSet("character", flag.ExitOnError)
	charSeed := charCmd.Int64("seed", 0, "an integer to regenerate a character from")

	if len(os.Args) < 2 {
		println("Expected a subcommand.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "character":
		charCmd.Parse(os.Args[2:])
		var seed int64
		if *charSeed == 0 {
			seed = time.Now().UnixNano()
		} else {
			seed = *charSeed
		}

		c := character.GenerateCharacter(seed)
		j, _ := json.Marshal(c)
		println(string(j))

	}
}
