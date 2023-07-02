# Gambit

Gambit is a shutdown scheduler utility software for Windows made in Go programming language with the Fyne GUI library. Gambit is licensed under the Apache-2.0 license, please read it carefully.

# Use

The use is simple for now. Given an input of time in the format HH:MM, it will schedule a shutdown to this given time automatically calculating difference in seconds between two dates (scheduled date and now date).

# Compiling

To compile basically you need to have all the imports, but fyne have its requirements to work also, so you will need the **Go tools** (at least version 1.12), a **C compiler** (GCC for example, to connect with system graphics drivers) and a **system graphics driver**, make sure you have them first on your system, then execute:

$ go get fyne.io/fyne/v2@latest
$ go install fyne.io/fyne/v2/cmd/fyne@latest
$ cd Gambit
$ go mod tidy
$ go build -ldflags -H=windowsgui
