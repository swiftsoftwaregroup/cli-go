# cli-go
Template for Command Line Interface (CLI) tool in Go

## Development

### Setup for macOS

Install `goenv`:

```bash
brew install goenv
```

Install Go:

```bash
# install go 1.22.5
goenv install 1.22.5

# to install latest Go version
goenv install latest
```

### Work on macOS

Configure project:

```bash
source configure.sh
```

Open the project in Visual Studio Code:

```bash
code .
```

### Build

```bash
go build ./...
```

### Run

```bash
echo "John" > name.txt

./cli-go greet name.txt
./cli-go greet --language es name.txt
./cli-go greet -l bg name.txt
```

Output:

```bash
Hello, John!
Hola, John!
Здравей, John!
```

### Test

```bash
go test -v
```

### Generate Docs

```bash
godoc
```

Browse docs:

```bash
# this package
open http://localhost:6060/pkg/github.com/swiftsoftwaregroup/cli-go

# all packages
open http://localhost:6060/pkg/
```

### How to create a new project

```bash
# create module
go mod init github.com/<username>/cli-go

# add packages
# see: https://cobra.dev
go get github.com/spf13/cobra

touch main.go
```

Tools:

```bash
go install golang.org/x/tools/cmd/godoc@latest
```

