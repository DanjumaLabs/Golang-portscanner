package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("go port_scanner starting...")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter target (ip or domain):")
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	fmt.Println("target set to:", target)

	fmt.Print("Enter port 1-1024 or 22,40,443: ")
	port_input, _ := reader.ReadString('\n')
	port_input = strings.TrimSpace(port_input)

	ports := parseports(port_input)
	fmt.Println("parsed ports:", ports)

	opencount := 0
	closedcount := 0

	var mu sync.Mutex

	var wg sync.WaitGroup

	for _, port := range ports {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()

			stealthDelay()

			open, banner := scanport(target, p)
			mu.Lock()
			if open {
				fmt.Printf("port %d is OPEN: %s\n", p, banner)
			} else {
				closedcount++
			}
			mu.Unlock()
		}(port)

	}
	wg.Wait()
	fmt.Println("\n--- scan summary ---")
	fmt.Printf("open ports: %d\n", opencount)
	fmt.Printf("closed ports: %d\n", closedcount)
}
