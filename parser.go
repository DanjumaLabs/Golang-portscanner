package main

import (
	"strconv"
	"strings"
)

func parseports(input string) []int {
	var ports []int

	if strings.Contains(input, "-") {
		rangeparts := strings.Split(input, "-")
		start, err1 := strconv.Atoi(rangeparts[0])
		end, err2 := strconv.Atoi(rangeparts[1])
		if err1 != nil || err2 != nil {
			return ports
		}
		for i := start; i <= end; i++ {
			ports = append(ports, i)
		}

	}
	parts := strings.Split(input, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		port, err := strconv.Atoi(part)

		if err != nil {
			continue
		}
		ports = append(ports, port)
	}

	return ports
}
