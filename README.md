# Create CLI command with Go

The purpose of this application is to create a command for cli for use from your local machine. The application will build 
a Go file and then place the built file in your /usr/local/bin.

### Prerequisites

Golang:
```
Installation instructions: https://golang.org/dl/
```

## Usage

How to install:

```
git clone https://github.com/brianjohnson2/create-cli-from-go.git
cd create-cli-from-go/
go run main.go createcommand
```

Confirm installation was successful:

```
createcommand --help
```

The application will create the command based on the first and only argument that is being based. The file that will be built defaults to main.go, 
but you can specify the file name by passing to -flag.