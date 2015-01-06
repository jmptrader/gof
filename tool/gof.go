package main

import (
	"fmt"
	"github.com/apoydence/gof/tool/generate"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func main() {
	if len(os.Args) <= 1 {
		help()
	}

	switch os.Args[1] {
	case "generate":
		generateCmd()
	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		help()
		break
	}
}

func generateCmd() {
	p := "."
	if len(os.Args) > 2 {
		p = os.Args[2]
		info, err := os.Stat(p)
		if err != nil || !info.IsDir() {
			fmt.Printf("Unable to open %s\n", p)
			os.Exit(1)
		}
	}

	walkFunc := func(p string, info os.FileInfo, err error) error {
		if path.Ext(p) == ".gof" {
			return convertGof(p)
		}
		return nil
	}

	if err := filepath.Walk(p, walkFunc); err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
}

func convertGof(p string) error {
	reader, err := os.Open(p)
	if err != nil {
		return err
	}

	goName := p[:len(p)-1]

	fmt.Printf("%s -> %s\n", p, goName)
	writer, err := os.Create(goName)
	if err != nil {
		return err
	}
	err = generate.GofToGo(reader, writer)

	reader.Close()
	writer.Close()

	return err
}

func walkDirs(dirPath string, ch chan string) {
	fis, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return
	}

	for _, fi := range fis {
		joined := path.Join(dirPath, fi.Name())
		if fi.IsDir() {
			walkDirs(joined, ch)
		} else {
			ch <- joined
		}
	}
}

func help() {
	useage := `Usage:

	gof command [arguments]
	
The commands are:

	generate	convert the GoF code to Go code

Use go help [command] for more information about a command.`

	fmt.Println(useage)
	os.Exit(1)
}
