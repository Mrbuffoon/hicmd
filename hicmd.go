package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

const global_version string = "1.0.0"

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

type cmd struct {
	name        string
	tag         string
	file_path   string
	description string
}

var cmd_array []cmd

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

func create_cmd_array() {
	var cmd_temp cmd
	var c_name, c_tag, c_description, c_filepath string
	fi, err := os.Open("/usr/local/hicmd/full_cmd.conf")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Sscan(string(a), &c_name, &c_tag, &c_description, &c_filepath)
		cmd_temp.name = c_name
		cmd_temp.tag = c_tag
		cmd_temp.description = c_description
		cmd_temp.file_path = c_filepath
		cmd_array = append(cmd_array, cmd_temp)
	}
}

func main() {
	create_cmd_array()
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}
	if version {
		fmt.Println("hicmd version: ", global_version)
		os.Exit(0)
	}
	if *list {
		//fmt.Println("list all commond")
		for _, temp_cmd := range cmd_array {
			fmt.Println(temp_cmd.name, temp_cmd.tag, temp_cmd.description)
		}
		os.Exit(0)
	}

	if tag != "" {
		//fmt.Println("Tag is " + tag)
		i := 0
		for _, temp_cmd := range cmd_array {
			if temp_cmd.tag == tag {
				fmt.Println(temp_cmd.name, temp_cmd.tag, temp_cmd.description)
				i++
			}
		}
		if i == 0 {
			fmt.Println("have no Tag ", tag)
		}
		os.Exit(0)
	}
	if name != "" {
		//fmt.Println("Name is " + name)
		i := 0
		for _, temp_cmd := range cmd_array {
			if temp_cmd.name == name {
				fmt.Println(temp_cmd.name, temp_cmd.tag, temp_cmd.description)
				i++
			}
		}
		if i == 0 {
			fmt.Println("have no Name ", name)
		}
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
