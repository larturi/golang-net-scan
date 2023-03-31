package portscan

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "scanme.nmap.org", "URL to scan")

func MainPortscan() {

	// Ej: go run port.go --site=scanme.nmap.org

	flag.Parse()

	var wg sync.WaitGroup

	for i := 0; i < 65535; i++ {
		wg.Add(1)

		go func(port int) {
			defer wg.Done()

			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))

			if err != nil {
				return
			}

			conn.Close()

			fmt.Printf("Port %d is open\n", port)
		}(i)

	}

	wg.Wait()

	fmt.Println("Scan completed")
}
