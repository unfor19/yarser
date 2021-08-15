package main

import (
	"log"
	"os"

	"github.com/acarl005/stripansi"
	"github.com/fsnotify/fsnotify"
	"github.com/mikefarah/yq/v4/cmd"
	"github.com/mikefarah/yq/v4/test"
	"github.com/spf13/cobra"
)

func getRootCommand() *cobra.Command {
	return cmd.New()
}

func CreateEmptyFile(filePath string) os.File {
	outputFile, err := os.Create(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	outputFile.Chmod(0755)

	return *outputFile
}

func RunYq(input string) string {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, input)
	if result.Error != nil {
		log.Println(result.Error)
		return ""
	}
	log.Println("Successfully executed:", "yq", input)
	return stripansi.Strip(result.Output)
}

func ParseYaml(srcFilePath string, dstFilePath string, outputFile os.File, verbose bool) {
	log.Println("Destination file path", dstFilePath)
	explodeResult := RunYq("eval-all" + " explode(.) " + srcFilePath)
	outputFile.Write([]byte(explodeResult))

	delResult := RunYq("eval-all" + " --inplace" + " del(.\".*\") " + dstFilePath)
	if delResult == "" {
		log.Println("Deleted all nodes that start with '.'")
	}
}

func WatchFile(srcFilePath string, dstFilePath string, verbose bool) {
	log.Println("Watching for changes in", srcFilePath, " ...")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					log.Println("here")
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					OutputFile := CreateEmptyFile(dstFilePath)
					defer OutputFile.Close()
					ParseYaml(srcFilePath, dstFilePath, OutputFile, verbose)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}

				log.Fatalln("error:", err)

			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add(srcFilePath); err != nil {
		log.Fatalln("ERROR", err)
	}

	<-done
}

func main() {
	argsWithoutProg := os.Args[1:]
	DefaultSrcFilePath := "tests/github-action/my-action.yml"
	DefaultDstFilePath := "tests/github-action/my-action-new.yml"
	Verbose := false

	if len(argsWithoutProg) < 2 {
		log.Println("Provide SrcFilePath and DstFilePath")
		log.Println("Example: yarser my-action-src.yml my-action.yml")
	}
	SrcFilePath := DefaultSrcFilePath
	if len(argsWithoutProg) > 0 && argsWithoutProg[0] != "" {
		SrcFilePath = argsWithoutProg[0]
	}
	DstFilePath := DefaultDstFilePath
	if len(argsWithoutProg) == 2 && argsWithoutProg[1] != "" {
		DstFilePath = argsWithoutProg[1]
	}
	if len(argsWithoutProg) == 3 && argsWithoutProg[2] == "--once" {
		OutputFile := CreateEmptyFile(DstFilePath)
		defer OutputFile.Close()
		ParseYaml(SrcFilePath, DstFilePath, OutputFile, Verbose)
	} else {
		WatchFile(SrcFilePath, DstFilePath, Verbose)
	}
}
