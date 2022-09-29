package main

import (
	"fmt"
	"time"
)

func main() {
	// Duration, there are Second, Minute, and Hour
	tenSeconds := 10 * time.Second
	duration, err := time.ParseDuration("1h1m1s")
	if err != nil {
		panic(err)
	}
	fmt.Println(tenSeconds, duration)
	// Sleep
	time.Sleep(time.Second)
	fmt.Println("after one second with sleep")
	// After, functionally the same as sleep but
	// usually used in a select case
	<-time.After(time.Second)
	fmt.Println("after one second with after")
	// Tick, we don't use this to prevent user error
	// because it would not be collected by the GC
	// Ticker
	done := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()
	work(done)
	// Location
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	fmt.Println(loc)
	// Time
	t := time.Now()
	fmt.Println(t)
	t, err = time.Parse(time.RFC3339, "2022-09-19T05:41:45.399Z")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	t = time.Date(2022, time.September, 19, 12, 42, 20, 0, loc)
	fmt.Println(t)
	// Add
	added := time.Now().Add(time.Hour)
	fmt.Println(added)
	// Sub
	previousTime := time.Now().Add(-1 * time.Hour)
	subbed := added.Sub(previousTime)
	fmt.Println(subbed)
	// Since
	s := time.Since(previousTime)
	fmt.Println(s)
	// Until
	s = time.Until(added)
	fmt.Println(s)
}

func work(done <-chan bool) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}
