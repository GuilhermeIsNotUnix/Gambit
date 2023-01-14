package main

import (
	"log"
	"os/exec"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Given the time in seconds, it will schedule the computer to shutdown in the especified time
func shutdown(timeInSeconds int) {
	cmd := exec.Command("shutdown", "-s", "-t", strconv.Itoa(timeInSeconds))
	out, err := cmd.Output()

	if err != nil {
		log.Println("Erro, não foi possivel: ", err)
	}

	log.Println("Output: ", string(out))
}

// Given the time in seconds, it will schedule the computer to shutdown in the especified time
func abortShutdown() {
	cmd := exec.Command("shutdown", "-a")
	out, err := cmd.Output()

	if err != nil {
		log.Println("Erro, não foi possivel: ", err)
	}

	log.Println("Output: ", string(out))
}

// Format the scheduled time string in a way that it returns the hour or the minutes. If hour_return is true, it will return hour, else return minutes.
func formatTimeString(scheduled_time string, hour_return bool) int {
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

// Given an hour and minutes the function constructs a formal standard date type, if you given hour is less the actual system hour, it means it will be that hour from the next day, or else its the same day
func constructDate(givenHour int, givenMinute int) time.Time {
	if givenHour < time.Now().Hour() {
		date := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, givenHour, givenMinute, 0, 0, time.UTC)
		return date
	}

	date := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), givenHour, givenMinute, 0, 0, time.UTC)
	return date
}

// Compare two dates and return the time difference between them in seconds
func compareDate(date1 time.Time, date2 time.Time) int {
	dif := date2.Sub(date1)
	time_in_seconds := dif.Seconds()

	return int(time_in_seconds)
}

// Validate input veryfying if its in the format hh:mm in the range 00:00 to 23:59
func validateTimeInput(input string) string {
	var scheduled_time string
	var invalid = false

	//fmt.Println("Digite uma hora para ser agendado o desligamento no formato (hh:mm, exemplo: 22:31)")
	//fmt.Scanln(&scheduledTime)

	if formatTimeString(scheduled_time, true) < 0 || formatTimeString(scheduled_time, true) > 23 {
		invalid = true

		if formatTimeString(scheduled_time, true) >= 0 || formatTimeString(scheduled_time, true) <= 23 {
			invalid = true

			if formatTimeString(scheduled_time, false) < 0 || formatTimeString(scheduled_time, false) > 59 {
				invalid = true
			}
		}

		if invalid != false {
			log.Println("Hora invalida, digite a hora novamente no formato hh:mm, lembrando que horas vão de 00 a 23 e minutos de 0 a 59")
			scheduled_time = ""

			return scheduled_time
		}
	}

	scheduled_time = input
	return scheduled_time
}

func getFutureHour(scheduledTime string) int {
	hour := formatTimeString(scheduledTime, true)
	return hour
}

func getFutureMinute(scheduledTime string) int {
	minutes := formatTimeString(scheduledTime, false)
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

		scheduled_time := validateTimeInput(input.Text)
		log.Println(scheduled_time)

		if scheduled_time != "" {
			future_date := constructDate(formatTimeString(scheduled_time, true), formatTimeString(scheduled_time, false))
			//future_date := constructDate(02, 07)
			timeInSeconds := compareDate(time.Now(), future_date)

			shutdown(timeInSeconds)
		}

	}), widget.NewButton("Cancelar agendamentos", func() {
		abortShutdown()
	}))

	window.SetContent(content)
	window.ShowAndRun()
}
