// Utilization monitor of your development machine

// Build a command line application within a single main.go file

// Criteria - collect the CPU utilization of your machine
// collect RAM utilization of your machine
// Collect backing storage utilization of your machine

package main

import (
	"fmt"
	"runtime"
)

func GoRoot() string {
	goroot := runtime.GOROOT()
	fmt.Print(goroot)
	return goroot
}

func main() {

	GoRoot()
	// // cpu, _ := cpu.Get()
	// memory, _ := memory.Get()

	// fmt.Println("Memmory used: %d", memory.Used)
	// // fmt.Println("CPU used: %d", cpu.Used)

}
