package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vgxbj/mailinone/internal/config"
	"github.com/vgxbj/mailinone/internal/mail"
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

	acc := configs.Accounts[0]
	c, err := mail.NewClient(&acc)
	if err != nil {
		log.Fatal(err)
	}

	imapc := c.IMAPClient()
	err = imapc.Login()
	if err != nil {
		log.Fatal(err)
	}

	defer imapc.Logout()

	mbs, err := imapc.GetMailboxes()
	if err != nil {
		log.Fatal(err)
	}

	for _, mb := range mbs {
		fmt.Println(mb.Name)
	}
}
