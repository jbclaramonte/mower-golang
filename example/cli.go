package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jbclaramonte/mower-golang/core/helper"
	"github.com/jbclaramonte/mower-golang/core/service"
	"github.com/jbclaramonte/mower-golang/example/cli/repository"
	cliService "github.com/jbclaramonte/mower-golang/example/cli/service"
)

func main() {
	file := flag.String("file", "", "relative path to the file with the commands for the mower")

	flag.Parse()

	log.Printf("file: %v \n", *file)

	repository := repository.NewLocalFileStorage()
	content, err := repository.GetFileContent(*file)
	if err != nil {
		panic(err)
	}
	fmt.Printf(*content)

	garden, mower, commands := cliService.TranslateContentToCommands(*content)
	for _, command := range commands {
		*mower = service.ApplyCommand(*garden, *mower, string(command))
	}

	fmt.Printf("final result: %v", helper.PrettyJson(mower))

}
