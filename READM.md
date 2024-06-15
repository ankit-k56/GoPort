# Go Port

[![Tool Category](https://badgen.net/badge/Tool/Port%20Scanner/black)](https://github.com/ankit-k56/GoPort)
[![APP Version](https://badgen.net/badge/Version/v1.1/red)](https://github.com/ankit-k56/GoPort)
[![Go Version](https://badgen.net/badge/Go/1.21.3/blue)](https://golang.org/doc/go1.21.3)

GoPort is a streamlined and efficient port scanning tool developed in GoLang. Engineered for speed and performance, this CLI tool can scan both TCP and UDP ports on a host. By leveraging GoLang's concurrency features, GoPort delivers fast and straightforward port scanning, making it both user-friendly and highly effective. Additionally, GoPort outputs the results to a separate file to prevent terminal clutter.

## Installation

To install GoPort, navigate to the root project directory and run the following commands:

```bash
  go build
  go install
```

## Commands

**goport**

The main command for GoPort.

**goport ping**

Used to perform port scanning.

## Flags

Specify the host to scan:

```bash
  goport ping -a google.com
```

Specify a single port to scan:

```bash
  goport ping -p 8080
```

Specify a range of ports to scan:

```bash
  goport ping -p 8080-8090
```

Specify multiple ports to scan:

```bash
  goport ping -p 8080,8090
```

Scan UDP ports by adding the -u flag, by default it scans tcp ports:

```bash
  goport ping -p 8080,8090 -h google.com -u
```

## Todo's

- [ ] Implement additional protocols for scanning.
- [ ] Add support for IPv6 addresses.
- [ ] Improve error handling and logging.
- [ ] Create a detailed user guide with examples.
- [ ] Develop a web-based interface for GoPort.
- [ ] Implement stealth scan.
- [ ] Adding support for storing output in user defined file.
