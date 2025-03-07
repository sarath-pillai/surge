# Surge : Minimal HTTP Performance Testing Tool

Surge is a quick way to benchmark your HTTP endpoints. It uses goroutines to initate concurrent requests for a set duration and reports back details like response time(Highest, Lowest), and status codes. Its a pretty light weight and minimal one that performs a quick concurrent test against a given endpoint. 

## Basic Usage

```bash
sh-3.2$ go run main.go -u https://www.example.com
Running Performance test against https://www.example.com with concurrency of 1
Execution duration complete
Lowest Response Time: 0.00s
Highest Response Time: 0.04s
Responses Other than 200: 0
0.04s      1256 200
0.02s      1256 200
0.01s      1256 200
0.00s      1256 200
```

By default the test is executed for 5 seconds. 

## Run concurrent tests for a specified duration

```bash
sh-3.2$ go run main.go -u https://www.example.com -c 100 -d 3
Running Performance test against https://www.example.com with concurrency of 100, for a duration of 3 seconds
Execution duration complete
Lowest Response Time: 0.01s
Highest Response Time: 0.07s
Responses Other than 200: 0
0.06s      1256 200
0.06s      1256 200
0.06s      1256 200
0.06s      1256 200
```

## Command line parameters

| Parameter     | Description                        |
|---------------|-----------------------------------|
| `-u`          | URL of the endpoint to test       |
| `-c`          | Number of concurrent requests     |
| `-d`          | Duration of the test in seconds   |
| `-m`          | HTTP method (GET or POST)         |
| `-b`          | HTTP body (for POST requests). Its a full file path with body data     |
| `-d`          | Duration for how long the test should run (default 5)     |
| `--header`          | Header Name and Value in the format headerName:headerValue. You can pass multiple using --header multiple times     |
| `-a`          | Basic HTTP Auth. "username:password"   |


```bash
Usage of Surge:
 surge -u 'https://www.example.com'
  -a string
        Basic authentication in the format username:password
  -b string
        HTTP body. This is a file containing the data that needs to be sent
  -c int
        How many Concurrent requests should be sent (default 1)
  -ct string
        contentType
  -d int
        Duration for how long the test should run (default 5)
  -header value
        Header Name and Value
  -m string
        HTTP method. Currently GET & POST is supported (default "GET")
  -u string
        The endpoint IP, URL against which the test needs to be performed
```