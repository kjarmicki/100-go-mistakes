package main

import "fmt"

const (
	StatusSuccess = "success"
)

func notify(status string) {
	fmt.Printf("notice: %s\n", status)
}

func doStuffAndNotify() string {
	var status string
	defer notify(status) // this will bind the current status argument - an empty string, and execute later
	status = StatusSuccess
	return status
}

func doStuffAndNotifyClosure() string {
	var status string
	defer func() {
		notify(status) // status referenced from the outside rather than passed as an argument
	}()
	status = StatusSuccess
	return status
}

func DeferArgumentEvaluation() {
	doStuffAndNotify()        // will print an empty string
	doStuffAndNotifyClosure() // will print success
}
