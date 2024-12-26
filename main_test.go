package main

import (
	"testing"
	"time"
)

func Test_fastScan(t *testing.T) {
	type args struct {
		start   int
		end     int
		addr    string
		timeout time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "scanme success",
			args: args{
				addr:    "scanme.nmap.org",
				start:   0,
				end:     100,
				timeout: 1 * time.Second,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fastScan(tt.args.start, tt.args.end, tt.args.addr, tt.args.timeout)
		})
	}
}

func Test_slowScan(t *testing.T) {
	type args struct {
		start   int
		end     int
		addr    string
		timeout time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "scanme success",
			args: args{
				addr:    "scanme.nmap.org",
				start:   0,
				end:     100,
				timeout: 1 * time.Second,
			},
		},
		{
			name: "scanme port below zero",
			args: args{
				addr:    "scanme.nmap.org",
				start:   -1,
				end:     100,
				timeout: 1 * time.Second,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slowScan(tt.args.start, tt.args.end, tt.args.addr, tt.args.timeout)
		})
	}
}

//func Test_worker(t *testing.T) {
//	type args struct {
//		ports   chan int
//		results chan int
//		addr    string
//		timeout time.Duration
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		{
//			name: "scanme success",
//			args: args{
//				addr:    "scanme.nmap.org",
//				ports:   make(chan int, 100),
//				results: make(chan int),
//				timeout: 1 * time.Second,
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			worker(tt.args.ports, tt.args.results, tt.args.addr, tt.args.timeout)
//		})
//	}
//}

func Test_workersScan(t *testing.T) {
	type args struct {
		start   int
		end     int
		address string
		timeout time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "scanme success",
			args: args{
				address: "scanme.nmap.org",
				start:   0,
				end:     100,
				timeout: 1 * time.Second,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workersScan(tt.args.start, tt.args.end, tt.args.address, tt.args.timeout)
		})
	}
}
