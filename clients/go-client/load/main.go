package load

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"log"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

type ThreadMetric struct {
	perc25 time.Duration
	perc50 time.Duration
	median time.Duration
	perc75 time.Duration
	perc90 time.Duration
	perc95 time.Duration
	perc99 time.Duration
}

var NumberOfThreads int
var NumberOfIterations int
var ThreadMetrics []ThreadMetric

func init()  {
	NumberOfThreads = 1
	NumberOfIterations = 1
	ThreadMetrics = make([]ThreadMetric, NumberOfThreads, NumberOfThreads)
}

func DoLoad(f func() (int64, error))  {
	var wg sync.WaitGroup

	for thread := 0; thread < NumberOfThreads; thread++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, threadNumber int, localThreadMetrics []ThreadMetric, numberOfIterations int) {
			defer wg.Done()

			reqDurations := make([]time.Duration, numberOfIterations, numberOfIterations)
			for i := 0; i< numberOfIterations; i++ {
				start := time.Now()
				res, err := f()
				if err != nil {
					log.Fatalln(err)
				}
				duration := time.Since(start)
				reqDurations[i] = duration
				log.Printf("Thread: %v. Result: %v. It took: %v", threadNumber, res, duration)
			}

			data := stats.LoadRawData(reqDurations)
			median, _ := stats.Median(data)
			perc25, _ := stats.Percentile(data, 25)
			perc50, _ := stats.Percentile(data, 50)
			perc90, _ := stats.Percentile(data, 90)
			perc95, _ := stats.Percentile(data, 95)
			perc99, _ := stats.Percentile(data, 99)
			localThreadMetrics[threadNumber].perc25 = time.Duration(perc25)
			localThreadMetrics[threadNumber].median = time.Duration(median)
			localThreadMetrics[threadNumber].perc50 = time.Duration(perc50)
			localThreadMetrics[threadNumber].perc90 = time.Duration(perc90)
			localThreadMetrics[threadNumber].perc95 = time.Duration(perc95)
			localThreadMetrics[threadNumber].perc99 = time.Duration(perc99)
		}(&wg, thread, ThreadMetrics, NumberOfIterations)
	}
	wg.Wait()
}

func PrintResults()  {
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.AlignRight|tabwriter.Debug)

	fmt.Fprintf(w, "Thread\tIterations\t25\tmedian\t50\t90\t99\t\n")
	for i, v := range ThreadMetrics {
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t\n", i, NumberOfIterations, v.perc25, v.median, v.perc50, v.perc90, v.perc99)
	}
	w.Flush()
}