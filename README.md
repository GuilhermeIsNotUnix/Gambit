# Gambit

Gambit is a shutdown scheduler utility software for Windows made in Go programming language with the Fyne GUI library. Gambit is licensed under the Apache-2.0 license, please read it carefully.

# Use

The use is simple for now. Given an input of time in the format HH:MM, it will schedule a shutdown to this given time automatically calculating difference in seconds between two dates (scheduled date and now date).

# Compiling

To compile basically you need to have all the imports, so use go get, and then:

$ go build -ldflags -H=windowsgui
