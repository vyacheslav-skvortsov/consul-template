package main // import "github.com/hashicorp/consul-template"

import "os"

func main() {
	cli := NewCLI(os.Stdout, os.Stderr)
	os.Exit(cli.Run(os.Args))

	//DEBUG
	// flags := []string{"consul-template", "-template", "./_test/test.tpl:./_test/test", "-once"}
	// cli := NewCLIDebug(os.Stderr)
	// os.Exit(cli.Run(flags))
}
