# Gambit

Gambit is a shutdown scheduler utility software for Windows made in Go programming language with the Fyne GUI library. Gambit is licensed under the Apache-2.0 license, please read it carefully.
Get it here (or on the releases tab): https://guilhermeisnotunix.github.io/software/gambit

## Windows SmartScreen or Antivirus preventing use

Unfortunately, from the tests I did, some antiviruses remove the Gambit .exe or in other cases, when running, Windows SmartScreen prevents its use with a security message. Note that there is nothing malicious about Gambit, the source is there and you can check it out or compile it yourself. The issue is that Gambit is a program without a digital signature and it is not in the Microsoft store either and because of that antiviruses or Windows SmartScreen think it's malicious without actually checking that it is not something malicious. And the criticism I leave here is that it is too expensive in my country to pay a monthly digital subscription to maintain a non-profit program. So Windows SmartScreen labeling anything malware is no coincidence, good job Microsoft, every day more reasons to make me prefer Linux...

# Use

The use is simple for now. Given an input of time in the format HH:MM, it will schedule a shutdown to this given time automatically calculating difference in seconds between two dates (scheduled date and now date).

# Compiling

To compile basically you need to have all the imports, but fyne have its requirements to work also, so you will need the **Go tools** (at least version 1.12), a **C compiler** (GCC for example, to connect with system graphics drivers) and a **system graphics driver**, make sure you have them first on your system, then execute:

$ go get fyne.io/fyne/v2@latest

$ go install fyne.io/fyne/v2/cmd/fyne@latest

$ cd Gambit

$ go mod tidy

$ go build -ldflags -H=windowsgui
