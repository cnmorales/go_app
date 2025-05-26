# Go Sum Application

A simple Go application that can sum two numbers provided as command-line parameters.

## Usage

```bash
# Run the application with two numbers
go run main.go -n1 5.2 -n2 3.8
```

## Build

To build the application:

```bash
go build -o sum
```

Then you can run the compiled binary:

```bash
./sum -n1 5.2 -n2 3.8
```

## Parameters

- `-n1`: First number to add (float)
- `-n2`: Second number to add (float)

## Example

```bash
$ go run main.go -n1 5.2 -n2 3.8
5.20 + 3.80 = 9.00
```
