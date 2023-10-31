# go-vestaboard

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/mikehelmick/go-vestaboard?tab=doc)
[![Go](https://github.com/mikehelmick/go-chaff/workflows/Go/badge.svg?event=push)](https://github.com/mikehelmick/go-vestaboard/actions?query=workflow%3AGo)

An unofficial client for the Vestaboard API in go.

# Usage

## Diferent types of clients.
### Readwrite Client
This client is meant if you are just targeting 1 Vestaboard with your application.

```
import github.com/mikehelmick/go-vestabord/v2/clients/api
import github.com/mikehelmick/go-vestaboard/clients/readwrite
client := api.NewClient()
board := readwrite.NewBoard("My first board", "YOUR_READWRITE_API_TOKEN")
client.SendText(context.Background(), board, "Hello world")
```

### Installable Client
This client is meant if you are producing a installable application in the vestaboard ecosystem.
The 
```
import github.com/mikehelmick/go-vestaboard/clients/installables
client := installables.NewClient("YOUR_API_KEY","YOUR_API_SECRET")

subscriptions,err := client.Subscriptions(context.Background())

client.SendText(context.Background(), subscriptions[0], "Hello world")
```

### Local Client
This client is meant if you are using the local send method. This requires additional steps

```
import github.com/mikehelmick/go-vestabord/v2/clients/api
import github.com/mikehelmick/go-vestaboard/clients/localboard
client := localboard.NewClient("http://vestaboard.local:7000")
board := localboard.NewBoard("My first board", "YOUR_LOCAL_API_TOKEN")
client.SendText(context.Background(), board, "Hello world")
```

# Examples

## Send Text

Does what it says - writes 'Hello World' to your vestaboard.

## Send Text Local

Does what it says - writes 'Hello World' to your vestaboard.

## Send Text Readwrite

Does what it says - writes 'Hello World' to your vestaboard.

## Clock

Writes out the current time about every 15 seconds.

## Game of Life

Conway's game of life.

## Subscriptions

Prints out the current subscriptions tied to the Vestaboard to the command-line.

## Test Pattern

Writes a fun test-pattern to the Vestaboard.

## Viewer

Prints results of the Vestaboard 'Viewer' API method to the command-line.