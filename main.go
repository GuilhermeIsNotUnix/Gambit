package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

// Format the scheduled time string in a way that it returns the hour or the minutes. If hourReturn is true, it will return hour, else return minutes.
func formatTimeString(scheduledTime string, hourReturn bool) int {
	if hourReturn == true {
		hour, err := strconv.Atoi(string([]rune(scheduledTime)[0:2]))
		if err != nil {
			fmt.Println("Erro: ", err)
		}

		return hour
	}

	minutes, err := strconv.Atoi(string([]rune(scheduledTime)[3:5]))
	if err != nil {
		fmt.Println("Erro: ", err)
	}

	return minutes
}

// Input times in the format hh:mm in the range 00:00 to 23:59
func getTime() string {
	var scheduledTime string

	fmt.Println("Digite uma hora para ser agendado o desligamento no formato (hh:mm, exemplo: 22:31)")
	fmt.Scanln(&scheduledTime)

	for formatTimeString(scheduledTime, true) < 0 || formatTimeString(scheduledTime, true) > 23 {
		fmt.Println("Hora invalida, digite a hora novamente no formato hh:mm, lembrando que horas vão de 00 a 23 e minutos de 0 a 59")
		fmt.Scanln(&scheduledTime)

		if formatTimeString(scheduledTime, true) >= 0 || formatTimeString(scheduledTime, true) <= 23 {
			for formatTimeString(scheduledTime, false) < 0 || formatTimeString(scheduledTime, false) > 59 {
				fmt.Println("Hora invalida, digite a hora novamente no formato hh:mm, lembrando que horas vão de 00 a 23 e minutos de 0 a 59")
				fmt.Scanln(&scheduledTime)
			}
		}
	}

	return scheduledTime
}

func getFutureHour(scheduledTime string) int {
	hour := formatTimeString(scheduledTime, true)
	return hour
}

func getFutureMinute(scheduledTime string) int {
	minutes := formatTimeString(scheduledTime, false)
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

// Given the time in seconds, it will schedule the computer to shutdown in the especified time
func shutdown(timeInSeconds int) {
	cmd := exec.Command("shutdown", "-s", "-t", strconv.Itoa(timeInSeconds))
	out, err := cmd.Output()

	if err != nil {
		fmt.Println("Erro, não foi possivel: ", err)
	}

	fmt.Println("Output: ", string(out))
}

// Given the time in seconds, it will schedule the computer to shutdown in the especified time
func abortShutdown() {
	cmd := exec.Command("shutdown", "-a")
	out, err := cmd.Output()

	if err != nil {
		fmt.Println("Erro, não foi possivel: ", err)
	}

	fmt.Println("Output: ", string(out))
}

func main() {
	scheduledTime := getTime()

	remainingHour := RemainingHour(scheduledTime)
	remainingMinutes := RemainingMinutes(scheduledTime)

	timeInSeconds := convertTimeToSeconds(remainingHour, remainingMinutes)

	var esc string
	fmt.Println("[PRESSIONE ENTER PARA CONCLUIR]")
	fmt.Scanln(&esc)
	shutdown(timeInSeconds)
}
