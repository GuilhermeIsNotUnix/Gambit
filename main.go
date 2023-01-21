package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Given the time in seconds, it will schedule the computer to Shutdown in the especified time
func Shutdown(timeInSeconds int, window fyne.Window) {
	cmd := exec.Command("shutdown", "-s", "-t", strconv.Itoa(timeInSeconds))
	out, err := cmd.Output()

	if err != nil {
		if strings.Contains(err.Error(), "1190") {
			dialog.ShowInformation(string(out), "Erro: O sistema já está com um agendamento em andamento\nCancele e faça um novo\n", window)
		}
	} else {
		dialog.ShowInformation("Gambit", "Você agendou o horario", window)
		log.Println("Output: ", string(out))
	}
}

// Given the time in seconds, it will schedule the computer to shutdown in the especified time
func AbortShutdown(window fyne.Window) {
	cmd := exec.Command("shutdown", "-a")
	out, err := cmd.Output()

	if err != nil {
		dialog.ShowInformation(string(out), "Erro: Não existem agendamentos feitos\n", window)
	} else {
		dialog.ShowInformation("Gambit", "Você cancelou os horario agendados", window)
		log.Println("Output: ", string(out))
	}
}

// The machine will shutdown immediatly
func ShutdownNow(window fyne.Window) {
	cmd := exec.Command("shutdown", "-s")
	out, err := cmd.Output()

	if err != nil {
		dialog.ShowInformation(string(out), "Erro: Não foi possivel desligar o sistema\n", window)
	} else {
		dialog.ShowInformation("Gambit", "Desligando...", window)
		log.Println("Output: ", string(out))
	}
}

// Format the scheduled time string in a way that it returns the hour or the minutes. If hour_return is true, it will return hour, else return minutes. It first checks if the unformatted string is not empty and if the lenght is 5, just then it catchs the hour or minute as integer.
func FormatTimeString(scheduled_time string, hour_return bool) int {
	if scheduled_time != "" && len(scheduled_time) == 5 {
		if hour_return == true {
			hour, err := strconv.Atoi(scheduled_time[0:2])
			if err != nil {
				log.Println("Erro no primeiro ATOI(): ", err)
			}

			return hour
		}

		minutes, err := strconv.Atoi(scheduled_time[3:5])
		if err != nil {
			log.Println("Erro no segundo ATOI(): ", err)
		}

		return minutes
	}

	return -1
}

// Given an hour and minutes the function constructs a formal standard date type, if you given hour is less the actual system hour, it means it will be that hour from the next day, or else its the same day
func ConstructDate(givenHour int, givenMinute int) time.Time {
	if givenHour < time.Now().Hour() {
		date := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, givenHour, givenMinute, 0, 0, time.UTC)
		return date
	}

	date := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), givenHour, givenMinute, 0, 0, time.UTC)
	return date
}

// Compare a given date with time.Now() and return the time difference between them in seconds
func CompareDate(given_date time.Time) int {
	dif_date := time.Date(given_date.Year()-time.Now().Year(), given_date.Month()-time.Now().Month(), given_date.Day()-time.Now().Day(), given_date.Hour()-time.Now().Hour(), given_date.Minute()-time.Now().Minute(), given_date.Second()-time.Now().Second(), given_date.Nanosecond()-time.Now().Nanosecond(), time.UTC)

	//dif := time.Until(given_date).Abs().Seconds()
	dif_time_in_seconds := int(dif_date.Hour()*60*60 + dif_date.Minute()*60 + dif_date.Second())
	log.Println("dif_time: ", dif_time_in_seconds)

	return dif_time_in_seconds
}

// Validate input veryfying if its in the format hh:mm in the range 00:00 to 23:59
func ValidateTimeInput(input string, window fyne.Window) string {
	var scheduled_time string
	var invalid = false

	//fmt.Println("Digite uma hora para ser agendado o desligamento no formato (hh:mm, exemplo: 22:31)")
	//fmt.Scanln(&scheduledTime)

	if FormatTimeString(input, true) < 0 || FormatTimeString(input, true) > 23 {
		invalid = true

		if FormatTimeString(input, true) >= 0 || FormatTimeString(input, true) <= 23 {
			invalid = true

			if FormatTimeString(input, false) < 0 || FormatTimeString(input, false) > 59 {
				invalid = true
			}
		}

		if invalid != false {
			dialog.ShowInformation("Gambit", "Hora inválida, digite a hora novamente no formato hh:mm, lembrando que horas vão de 00 a 23 e minutos de 0 a 59", window)
			scheduled_time = ""

			return scheduled_time
		}
	}

	scheduled_time = input
	return scheduled_time
}

func GetFutureHour(scheduledTime string) int {
	hour := FormatTimeString(scheduledTime, true)
	return hour
}

func GetFutureMinute(scheduledTime string) int {
	minutes := FormatTimeString(scheduledTime, false)
	return minutes
}

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Gambit v1.0")
	window.Resize(fyne.NewSize(1000, 470))

	input := widget.NewEntry()
	input.SetPlaceHolder("(hh:mm, exemplo: 22:31)")

	content := container.NewVBox(input, widget.NewButton("Agendar", func() {
		log.Println("Content was:", input.Text)

		scheduled_time := ValidateTimeInput(input.Text, window)

		if scheduled_time != "" {
			scheduled_date := ConstructDate(FormatTimeString(scheduled_time, true), FormatTimeString(scheduled_time, false))
			time_in_seconds := CompareDate(scheduled_date)

			Shutdown(time_in_seconds, window)
		}

	}), widget.NewButton("Cancelar agendamentos", func() {
		AbortShutdown(window)
	}),

		widget.NewButton("Desligar agora", func() {
			ShutdownNow(window)
		}))

	window.SetContent(content)
	window.ShowAndRun()
}
