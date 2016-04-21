# go-cli-base
A simple Go cli which you can use as a base for creating your own. Decoupled from any external framework and easy to support new flags.

## Running

Set your `$GOPATH` to the root of the **go-cli-base** directory (if your directory is under `~` then you should set it as `~/go-cli-base/`) by running `export GOPATH=~/go-cli-base/`

Set your `$GOBIN` to the `bin/` directory under **go-cli-base** by running `export GOBIN=$GOPATH/bin`

Install the cli by running `go install main/main.go`. You should now have an executable called **main** in your `go-cli-base/bin` directory.

Now, simply run `main -h` from the `go-cli-base/bin` directory.

## Setup

Built with **Go** version **1.5.3**, so I recommend using this or a newer version of Go

## Testing

After doing the setup, run `go test ../tests`

## Adding new flags

To add a new flag, simple add a new entry to the `AcceptedFlags` map in the `flags.go` file. Create a function to execute when that flag is present, and voilà!



