# bcrypt-cli

`bcrypt-cli` is the CLI tool for hashing passwords with bcrypt.

# Install

```bash
go install github.com/ryicoh/bcrypt-cli
```

# Usage

It can be used like `base64` command.

```bash
$ echo "password" | bcrypt-cli
$2a$10$qdKkslZfw1npRKSvsvYTaeWnTWAxYQyk6qYYJrHkvUv.wf87IShRK
```

# Help
```
$ bcrypt-cli -h
Usage of bcrypt-cli:
  -c int
        Hashing cost. Default is 10 (default 10)
  -cost int
        Hashing cost. Default is 10 (default 10)
  -n    Do not print the trailing newline character.
```
