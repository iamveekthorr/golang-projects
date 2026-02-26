package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/alexflint/go-arg"

	"github.com/iamveekthorr/mgrep/worker"
	"github.com/iamveekthorr/mgrep/worklist"
)

func getDirectories(workList *worklist.Worklist, path string) {
	entries, err := os.ReadDir(path)

	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			nextPath := filepath.Join(path, entry.Name())
			getDirectories(workList, nextPath) // Recursively get directories
		} else {
			workList.Add(worklist.NewJob(filepath.Join(path, entry.Name())))
		}
	}
}

var args struct {
	SearchTerm      string `arg:"positional,required"`
	SearchDirectory string `arg:"positional"`
}

func main() {
	arg.MustParse(&args)

	var workersWaitGroup sync.WaitGroup

	workerList := worklist.New(100)

	results := make(chan worker.Result, 100)

	numberOfWorkers := 10

	workersWaitGroup.Add(100)

	go func() {
		defer workersWaitGroup.Done()

		getDirectories(&workerList, args.SearchDirectory)
		workerList.Finalize(numberOfWorkers)
	}()

	for i := 0; i < numberOfWorkers; i++ {
		workersWaitGroup.Add(1)
		go func() {
			defer workersWaitGroup.Done()
			for {
				workEntry := workerList.Next()
				// check if empty
				if workEntry.Path != "" {
					workerResult := worker.FindTextInFile(workEntry.Path, args.SearchTerm)

					if workerResult != nil {
						for _, res := range workerResult.All {
							results <- res
						}
					}
				}
				return // if the path is empty
			}
		}()
	}

	blockWorkerWaitGroup := make(chan struct{})

	go func() {
		workersWaitGroup.Wait()
		close(blockWorkerWaitGroup) // close channel
	}()

	var displayWaitGroup sync.WaitGroup

	displayWaitGroup.Add(1)

	go func() {
		for {
			select {
			case res := <-results:
				fmt.Printf("%v [%v]: %v\n", res.Path, res.LineNumber, res.Line)
			case <-blockWorkerWaitGroup:
				if len(results) == 0 {
					displayWaitGroup.Done()
					return
				}
			}
		}
	}()

	displayWaitGroup.Wait()
}
