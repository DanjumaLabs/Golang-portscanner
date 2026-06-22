# Golang-portscanner


A lightweight, concurrent port scanner written in Go that can scan multiple ports on a target host with stealth capabilities.

## Features

- **Concurrent Scanning**: Uses goroutines for fast parallel port scanning
- **Flexible Port Input**: Supports individual ports, comma-separated lists, or port ranges
  - Examples: `22`, `22,80,443`, `1-1024`
- **Banner Grabbing**: Detects service types (SSH, HTTP) from server responses
- **Stealth Mode**: Includes random delays between port connections to avoid triggering IDS systems
- **TCP Connection-based**: Uses TCP connect() method for reliable port detection
- **Real-time Output**: Shows open ports immediately as they're discovered

## Requirements

- Go 1.11 or higher

## Installation

```bash
go build -o portscanner

USAGE
./portscanner

Then follow the prompts
Enter target (ip or domain): 192.168.1.1
Enter port 1-1024 or 22,40,443: 1-100

Example
target: example.com
ports: 80

Output
go port_scanner starting...
Enter target (ip or domain): example.com
target set to: example.com
Enter port 1-1024 or 22,40,443: 80,443
parsed ports: [80 443]

port 80 is OPEN: HTTP|HTTP/1.1 200 OK...
port 443 is OPEN: HTTP|HTTP/1.1 200 OK...

--- scan summary ---
open ports: 2
closed ports: 0

For any errors,contributions or question
ashtech01@protonmail.com

Legal Notice
This tool is intended for authorized network testing and security research only. Unauthorized access to computer systems is illegal. Always obtain proper authorization before scanning networks.
