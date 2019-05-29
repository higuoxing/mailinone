package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vgxbj/mailinone/pkg/config"
)

var (
	configFilePath *string
)

func init() {
	configFilePath = flag.String("c", "", "Path to configuration file")

	flag.Parse()
}

func main() {
	var (
		configs *config.Configs
		err     error
	)

	configs, err = config.ReadConfigFromFile(*configFilePath)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	err = configs.Verify()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Println(configs)
}
