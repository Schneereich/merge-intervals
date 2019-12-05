package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Schneereich/merge-intervals/interval"
)

func main() {
	if len(os.Args) < 2 {
		intervals := []interval.Interval{
			{15, 20},
			{2, 6},
			{17, 24},
		}
		log.Println("No intervals given with the execution of the program. Using default values:", intervals)
		result := interval.Merge(intervals)
		fmt.Println(result)
		return
	}

	// read commandline input and concatenate
	args := strings.TrimSpace(strings.Join(os.Args[1:],""))
	log.Println("The following intervals are read from the input:", args)

	// parse interval pairs of the commandline input
	object, _ := regexp.Compile(`(\[(\d+),(\d+)\])`)
	matches := object.FindAllStringSubmatch(args, -1)
	var intervals []interval.Interval
	for _, match := range matches {
		start, _ := strconv.Atoi(match[2])
		end, _ := strconv.Atoi(match[3])
		intervals = append(intervals, interval.Interval{Start: start, End: end})
	}

	result := interval.Merge(intervals)
	fmt.Println(result)
	return
}
