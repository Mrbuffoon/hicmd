// cfgtest.go
package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	help    bool
	version bool
	list    *bool

	tag   string
	name  string
	fuzzy string
	edit  string
)

var usage_help string = `
Usage: hicmd <options> [parm]

Options:
-h  show help and exit
-v  show version and exit
-l  list all commond and exit
-t  quary commond according tag
-n  quary commond according name
-f  fuzzy quary commond according string
-e  edit commond according name
`

func init() {
	flag.BoolVar(&help, "h", false, "show help and exit")
	flag.BoolVar(&version, "v", false, "show version and exit")
	list = flag.Bool("l", false, "list all commond and exit")

	flag.StringVar(&tag, "t", "", "quary commond according tag")
	flag.StringVar(&name, "n", "", "quary commond according name")
	flag.StringVar(&fuzzy, "f", "", "fuzzy quary commond according string")
	flag.StringVar(&edit, "e", "", "edit commond according name")

	flag.Usage = usage
}

func usage() {
	fmt.Println(usage_help)
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}
	if version {
		fmt.Println("hicmd version: 1.0.0")
		os.Exit(0)
	}
	if *list {
		fmt.Println("list all commond")
		os.Exit(0)
	}

	if tag != "" {
		fmt.Println("Tag is " + tag)
		os.Exit(0)
	}
	if name != "" {
		fmt.Println("Name is " + name)
		os.Exit(0)
	}
	if edit != "" {
		fmt.Println("Edit commond is " + edit)
		os.Exit(0)
	}
	if fuzzy != "" {
		fmt.Println("Fuzzy is " + fuzzy)
		os.Exit(0)
	}

}
