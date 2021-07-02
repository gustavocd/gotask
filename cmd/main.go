package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gustavocd/gotask/internal/device"
	"github.com/gustavocd/gotask/internal/file"
)

var (
	buildTime string
	version   string
)

func main() {
	if err := run(); err != nil {
		log.Printf("startup %v", err)
		os.Exit(1)
	}
}

func run() error {
	displayVersion := flag.Bool("version", false, "Display version and exit")

	flag.Parse()

	if *displayVersion {
		fmt.Printf("Version:\t%s\n", version)
		fmt.Printf("Build time:\t%s\n", buildTime)
		os.Exit(0)
	}

	inputs, err := file.ReadDeviceInputs("./challenge.in")
	if err != nil {
		return err
	}

	var combinations [][][]int
	for _, input := range inputs {
		c := device.ProduceCombinations(
			input.Capacity,
			input.Tasks[0],
			input.Tasks[1],
		)
		combinations = append(combinations, c)
	}

	f := device.FormatCombinations(combinations)
	err = os.WriteFile("./challenge.out", []byte(f), 0777)
	if err != nil {
		return err
	}

	fmt.Println("Combinations created successfully!")

	return nil
}
