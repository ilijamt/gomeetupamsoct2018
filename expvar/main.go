package main

import (
	"expvar"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"
)

type TimeVar struct{ v time.Time }

func (o *TimeVar) Set(date time.Time)         { o.v = date }
func (o *TimeVar) Add(duration time.Duration) { o.v = o.v.Add(duration) }
func (o *TimeVar) String() string             { return fmt.Sprintf("\"%s\"", o.v.Format(time.RFC3339)) }

func NewStats(name string) *expvar.Map {
	stats = expvar.NewMap(name)
	stats.Set("counter", new(expvar.Int))
	stats.Set("success_rate", new(expvar.Float))
	stats.Set("pid", new(expvar.Int))

	go func() {
		t := time.NewTicker(time.Millisecond)
		for {
			select {
			case <-t.C:
				stats.Add("counter", rand.Int63n(100))
				stats.Get("success_rate").(*expvar.Float).Set(rand.Float64())
			}
		}
	}()

	return stats
}

var stats *expvar.Map
var lastUpdate *TimeVar

func init() {
	stats = NewStats("stats")
	lastUpdate = &TimeVar{}
	lastUpdate.Set(time.Now())
	stats.Get("pid").(*expvar.Int).Set(int64(os.Getpid()))
	expvar.Publish("last_update", lastUpdate)
	expvar.Publish("goroutines", expvar.Func(func() interface{} {
		return fmt.Sprintf("%d", runtime.NumGoroutine())
	}))
	expvar.Publish("cgocall", expvar.Func(func() interface{} {
		return fmt.Sprintf("%d", runtime.NumCgoCall())
	}))
	expvar.Publish("cpu", expvar.Func(func() interface{} {
		return fmt.Sprintf("%d", runtime.NumCPU())
	}))
}

func randomGoRoutineSpawn() {

	for {
		go func() {
			r := rand.Intn(10)
			time.Sleep(time.Duration(r) * time.Microsecond)
		}()
		r := rand.Intn(3)
		time.Sleep(time.Duration(r) * time.Microsecond)
	}
}

func main() {
	go randomGoRoutineSpawn()
	http.ListenAndServe(":6000", http.DefaultServeMux)

}
