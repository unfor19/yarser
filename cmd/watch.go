/*
Copyright © 2021 Meir Gabay <unfor19@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

// Global variables
var CounterWriteEvents = 0

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch file for changes",
	Long: `Watch file for changes and activate a command.
For example:

yarser watch "$YARSER_WATCH_FILE_PATH"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		CustomWatcher(args[0], "", readFile)
	},
}

func readFile(filePath string, dstFilePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func CustomWatcher(watchFilePath string, dstFilePath string, fn callback) {
	logger.Info("Watching for changes in ", watchFilePath, " ...")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					// Write events are doubled, no idea why, so once an event occurs we want to stop the following one
					if CounterWriteEvents == 0 {
						if err := fn(watchFilePath, dstFilePath); err != nil {
							logger.Error("Error invoking callback", watchFilePath, dstFilePath)
							log.Fatalln(err)
						}
						CounterWriteEvents += 1
					} else {
						CounterWriteEvents = 0
					}

				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logger.Error("Unknown event error")
				log.Fatalln(err)
			}
		}
	}()

	if err := watcher.Add(watchFilePath); err != nil {
		logger.Error("Failed to watch file")
		log.Fatalln(err)
	}

	<-done
}

func init() {
	rootCmd.AddCommand(watchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// watchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// watchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
