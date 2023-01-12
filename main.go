package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

type TimeData struct {
	Hour    int
	Minutes int
}

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

func getScheduledHour(scheduledTime string) int {
	hour := formatTimeString(scheduledTime, true)

	return hour
}

func getScheduledMinutes(scheduledTime string) int {
	minutes := formatTimeString(scheduledTime, false)

	return minutes
}

func getHour() int {
	return time.Now().Hour()
}

func getMinute() int {
	return time.Now().Minute()
}

func RemainingHour(scheduledTime string) int {
	if getScheduledHour(scheduledTime) < getHour() {
		remainingHour := getScheduledHour(scheduledTime) + 24 - getHour()
		return remainingHour
	}

	remainingHour := getScheduledHour(scheduledTime) - getHour()
	return remainingHour
}

func RemainingMinutes(scheduledTime string) int {
	if getScheduledMinutes(scheduledTime) < getMinute() {
		remainingMinutes := getScheduledMinutes(scheduledTime) + 60 - getMinute()
		return remainingMinutes
	}

	remainingMinutes := getScheduledMinutes(scheduledTime) - getMinute()
	return remainingMinutes
}

// Convert the hours and minutes to seconds
func convertTimeToSeconds(hour int, minutes int) int {
	timeInSeconds := hour*60*60 + minutes*60
	return timeInSeconds
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
