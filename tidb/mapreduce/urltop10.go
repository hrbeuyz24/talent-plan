package main

import (
	"strconv"
	"strings"
)

// URLTop10 .
func URLTop10(nWorkers int) RoundsArgs {
	// YOUR CODE HERE :)
	// And don't forget to document your idea.
	var args RoundsArgs

	args = append(args, RoundArgs{
		MapFunc:    ExampleURLCountMap,
		ReduceFunc: ExampleURLCountReduce,
		NReduce:    nWorkers,
	})

	args = append(args, RoundArgs{
		MapFunc: URLTop10Map,
		//MapFunc:    ExampleURLCountMap,
		ReduceFunc: ExampleURLTop10Reduce,
		NReduce:    1,
	})
	return args
}

func URLTop10Map(filename string, contents string) []KeyValue {
	lines := strings.Split(contents, "\n")
	cnts := make(map[string]int, len(lines))
	for _, url := range lines {
		tmp := strings.Split(url, " ")
		if len(tmp) != 2 {
			continue
		}
		cnt, _ := strconv.Atoi(tmp[1])
		cnts[tmp[0]] = cnt
	}

	kvs := make([]KeyValue, 0, 10)
	urls, num := TopN(cnts, 10)
	for index, url := range urls {
		kvs = append(kvs, KeyValue{
			Key:   "",
			Value: url + " " + strconv.Itoa(num[index]),
		})
	}
	return kvs
}
