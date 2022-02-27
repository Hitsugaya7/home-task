# Golang Tool

This repository is a golang tool which a tool which makes http requests for provided URLs and prints the address of the request along with the MD5 hash of the response.Utility can handle multiple request at the same time. Number of concurrent requests can be providen by paralel option (by default == 10).

## Prerequisites

Before you start, please make sure you have Go installed in your system. If not, please use the following link to install Golang:
https://golang.org/doc/install

## Getting Started

Clone the git repository in your system and then cd into project root directory

```bash
$ git clone https://github.com/Hitsugaya7/home-task
$ cd home-task
```

Build your tool by executing the following steps
```bash
$ cd tool
$ go build -o myhhtp
```

## Sample Outputs

This tool takes address of the http requests as command line arguments. See below examples
```bash
$ ./myhttp google.com yahoo.com gmail.com facebook.com

```

```bash
$ ./myhttp google.com yahoo.com gmail.com facebook.com
http://facebook.com 0b935fb3f317b074e91dfb076f05a627
http://google.com 7d4205bb30dea4979799361ea7069a24
http://gmail.com 078b66e1e53497cacf18f07e5f555e35
http://yahoo.com 9cf8d5d4eebbb6d0e10c055bf94b82e6

```

This tool also takes "parallel" input parameter(integer value) to execute the http requests in parallel. The default value of "parallel" is 10.
```bash
$ ./myhttp -parallel 3 google.com yahoo.com gmail.com facebook.com
http://facebook.com 0b935fb3f317b074e91dfb076f05a627
http://google.com 7d4205bb30dea4979799361ea7069a24
http://gmail.com 078b66e1e53497cacf18f07e5f555e35
http://yahoo.com 9cf8d5d4eebbb6d0e10c055bf94b82e6

```

In case of failures(say incorrect address), it prints error message and gets exited.
```bash
$ ./myhttp -parallel 3 google.com yahoo.com gmail.com facebook.com hello.ret
Request Error:  Get "http://hello.ret": dial tcp: lookup hello.ret: no such host

```

This repository includes dependencies only from Go standard libraries. It also contains unit test cases which provide the industry standard code coverage and some benchmark tests.

Run your tests by executing the following steps
```bash
$ cd tool
$ go test -bench=. 
```
