package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sort"
	"sync"
	"time"
)

func main() {
	var start, end int
	var address string
	var timeout time.Duration
	flag.IntVar(&start, "start", 1, "Start from port number greater or equals than 1")
	flag.IntVar(&end, "end", 1024, "End with port number")
	flag.StringVar(&address, "addr", "scanme.nmap.org", "Address to scan")
	flag.DurationVar(&timeout, "t", 5*time.Second, "Timeout in seconds")
	flag.Parse()

	//slowScan(start, end, address, timeout)
	//fastScan(start, end, address, timeout)
	workersScan(start, end, address, timeout)
}

func slowScan(start, end int, addr string, timeout time.Duration) {
	for i := start; i <= end; i++ {
		address := fmt.Sprintf("%s:%d", addr, i)
		conn, err := net.DialTimeout("tcp", address, timeout)
		if err != nil {
			// порт закрыт или отфильтрован
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}

func fastScan(start, end int, addr string, timeout time.Duration) {
	var wg sync.WaitGroup
	for i := start; i <= end; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", addr, j)
			conn, err := net.DialTimeout("tcp", address, timeout)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait()
}

func workersScan(start, end int, address string, timeout time.Duration) {
	log.Printf("Start scanning - %s", address)
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, address, timeout)
	}
	log.Printf("%d workers started", cap(ports))
	if start < 1 {
		start = 1
	}
	go func() {
		for i := start; i <= end; i++ {
			ports <- i
		}
	}()
	for i := start - 1; i < end; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)

	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}

func worker(ports, results chan int, addr string, timeout time.Duration) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", addr, p)
		conn, err := net.DialTimeout("tcp", address, timeout)
		//log.Printf("Scanning %s:%d", addr, p)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}
