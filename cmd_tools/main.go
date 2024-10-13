package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Yamon955/Learn/cmd_tools/cmd"
	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	trpc "trpc.group/trpc-go/trpc-go"
)

func main() {
	_ = trpc.NewServer()
	if len(os.Args) < 2 {
		exitOnError(1)
	}
	c, ok := cmd.CMDs[os.Args[1]]
	if !ok {
		exitOnError(1)
	}
	if err := handle(c, os.Args[2:]); err != nil {
		os.Exit(2)
	}

}

func exitOnError(code int) {
	fmt.Println("Usage:", os.Args[0], cmd.CMDs.GetKeys())
	os.Exit(code)
}

func handle(c cmd.Command, args []string) error {
	if err := parseArgs(c, args); err != nil {
		log.Printf("parse args error: %v\n", err)
		return err
	}
	ret := c.Process()
	fmt.Printf("%s\n", ret)
	return nil
}

func parseArgs(c cmd.Command, args []string) (err error) {
	cmd, needArgs := c.Get()
	defer func() {
		if err != nil {
			cmd.Usage()
		}
	}()
	if !needArgs {
		return nil
	}
	if len(args) == 0 {
		return fmt.Errorf("empty args")
	}
	return cmd.Parse(args)
}
