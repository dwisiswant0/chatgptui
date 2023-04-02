package main

import (
	"fmt"
	"os"

	"github.com/dwisiswant0/chatgptui/common"
)

func showBanner() {
	fmt.Fprintf(os.Stderr, header, common.Version)
}

func showUsage() {
	main := os.Args[0]
	fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", main)
	fmt.Fprint(os.Stderr, options)
	fmt.Fprintf(os.Stderr, examples, main, main)
}

func showVersion() {
	fmt.Fprintf(os.Stderr, "ChatGPTUI %s\n", common.Version)
	os.Exit(2)
}

func listAllModels() {
	for _, model := range common.OpenaiModels {
		fmt.Println(model)
	}
	os.Exit(0)
}

func removeConfig() {
	_ = os.Remove(common.GetConfigPath())
	os.Exit(0)
}
