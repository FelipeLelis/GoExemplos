package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("esperado 'foo' ou 'bar' subcomandos")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcomando 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcomando 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("esperado 'foo' ou 'bar' subcomandos")
		os.Exit(1)
	}
}