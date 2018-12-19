package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mslusarczyk/median/config"
	"github.com/mslusarczyk/median/slidingwindow"
)

func main() {
	params, err := config.ParseParams()
	if err != nil {
		fmt.Printf("Parameters parsing failed: %s", err)
		return
	}
	tw := slidingwindow.NewSlidingWindow(params.WindowSize)

	in, err := os.Open(params.InputFile)
	if err != nil {
		log.Fatalf("Cannot open `int` file: [%s], err: %s", params.InputFile, err)
	}
	defer in.Close()

	out, err := os.Create(params.OutputFile)
	if err != nil {
		log.Fatalf("Cannot create `out` file: [%s], err: %s", params.OutputFile, err)
	}
	defer out.Close()

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		metricValue := scanner.Text()
		metricValue = strings.TrimSpace(metricValue)
		i, err := strconv.Atoi(metricValue)
		if err != nil {
			log.Fatalf("Cannot parse value [%s], err: %s", metricValue, err)
		}

		tw.AddDelay(i)

		_, err = fmt.Fprintf(out, "%d\r\n", tw.GetMedian())
		if err != nil {
			log.Fatalf("Cannot write to `out` file: [%s], err: %s", params.OutputFile, err)
		}
	}
}
