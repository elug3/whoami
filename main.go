package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var usageStr = `
Usage: go run main.go [options]
Options:
  -port <port>    Specify the port to listen on (default: 8080)
  -host <host>    Specify the host to listen on (default: localhost)
  -h, --help    Show this help message and exit

`

// printUsage prints the usage information and exits the program.
func printUsage() {
	log.Print(usageStr)
	os.Exit(0)
}

// configureOptions configures the command-line options for the application.
// if the help flag is set, it prints the usage information and exits the program.
func configureOptions(fs *flag.FlagSet, args []string) (map[string]string, error) {
	var (
		port *string
		host *string
		help *bool
	)

	fs.Usage = func() {
		log.Print(usageStr)
	}

	port = fs.String("port", "8080", "Specify the port to listen on")
	host = fs.String("host", "localhost", "Specify the host to listen on")
	help = fs.Bool("h", false, "Show this help message and exit")

	err := fs.Parse(args)
	if err != nil {
		log.Fatalf("Failed to parse options: %v", err)
	}

	if *help {
		printUsage()
	}

	return map[string]string{
		"port": *port,
		"host": *host,
	}, nil
}

// parseRequestDetails extracts and formats
// details from the HTTP request. (url, method, headers, timestamp, ip address)
func parseRequestDetails(r *http.Request) map[string]string {
	timestamp := time.Now().Format(time.RFC3339)
	return map[string]string{
		"URL":       r.URL.String(),
		"Method":    r.Method,
		"Headers":   r.Header.Get("Content-Type"),
		"Timestamp": timestamp,
		"IPAddress": r.RemoteAddr,
	}

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		// response request details
		requestDetails := parseRequestDetails(r)
		for k, v := range requestDetails {
			w.Write([]byte(k + ": " + v + "\n"))
		}
	})
	fs := flag.NewFlagSet("main", flag.ContinueOnError)
	opts, err := configureOptions(fs, os.Args[1:])
	if err != nil {
		log.Fatalf("Failed to configure options: %v", err)
	}

	port := opts["port"]
	host := opts["host"]
	addr := net.JoinHostPort(host, port)

	log.Printf("Server is running on http://%s\n", addr)
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
