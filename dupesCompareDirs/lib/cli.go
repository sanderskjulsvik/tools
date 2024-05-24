package dupescomparedirs

import (
	"flag"
	"fmt"
	"os"
)

// HandleCliInput will run flag.Parse,
// this means that if you want to get more cli input you would need to add them before this call

type CliInput struct {
	OutputJson bool
	Dir1       string
	Dir2       string
}

func HandleCliInput() CliInput {
	outputJson := flag.Bool("json", false, "If set to true Output as json")
	dir1 := flag.String("dir1", "", "Path to 1st dir")
	dir2 := flag.String("dir2", "", "Path to 2nd dir")
	flag.Parse()

	// Check if directory paths are provided
	if *dir1 == "" || *dir2 == "" {
		fmt.Println("Please provide directory paths to compare")
		os.Exit(1)
	}

	return CliInput{
		OutputJson: *outputJson,
		Dir1:       *dir1,
		Dir2:       *dir2,
	}

}
