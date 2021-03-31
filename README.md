# Log parser Go

This program extracts and processes information from log file (Check programming-task/programming-task-example-data.log for example). In particular, the program will return:
* The number of unique IP addresses
* The top 3 most visited URLs
* The top 3 most active IP addresses

## Solution

* Open provided file.
* Read file line by line to extract IP address and URL and update collected information. Assumption: 1 log line contains 1 IP address (IPv4, a numeric address where binary bits are separated by a dot) and 1 URL (an absolute URL starts with http or https or relative URL starts with /). Print an error if one of values is not found.
* Process collected information to answer quesitions and write the result to output.txt.

## How to run

    go run cmd/parsing-logs/main.go [optional filename]

## How to format

    gofmt -w -s . [or optional filename]

## How to run unit tests

    go test ./...

    go clean -testcache (to clean cache)

## How to run unit tests coverage (html)

    go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

## Libraries

* Logging ([log](//golang.org/pkg/log/))
* Printing strings ([fmt](//golang.org/pkg/fmt/))
* Reading files ([bufio](//golang.org/pkg/bufio/))
* Operating system features ([os](//golang.org/pkg/os/))
* Errors ([errors](//golang.org/pkg/errors/))
* Regular expressions ([regexp](//golang.org/pkg/regexp/))
* Strings manipulations ([strings](//golang.org/pkg/strings/))
* Sorting ([sort](//golang.org/pkg/sort/))
* Testing ([testing](//golang.org/pkg/sort/) and [testify](//https://godoc.org/github.com/stretchr/testify/assert"))