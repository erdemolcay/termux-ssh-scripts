package termux

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallRequirements() {
	if runningInTermux() {
		fmt.Println("Installing required packages...")
		o, err := exec.Command(
			"pkg",
			"install",
			"termux-api",
			"openssh",
			"-y").Output()
		fmt.Println(string(o))
		if err != nil {
			fmt.Println("Required packages could not be installed.")
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Successfully installed required packages.")
	}
}

func ScheduleJob() {
	if runningInTermux() {
		fmt.Println("Scheduling job...")
		o, err := exec.Command(
			"termux-job-scheduler",
			"-s",
			"/data/data/com.termux/files/usr/bin/termux-ssh-scripts-update",
			"--period-ms=900000",
			"--persisted=true").Output()
		fmt.Println(string(o))
		if err != nil {
			fmt.Println("Job could not be scheduled.")
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Successfully scheduled job.")
	}
}

func runningInTermux() bool {
	if os.Getenv("PREFIX") == "/data/data/com.termux/files/usr" {
		return true
	}
	return false
}
