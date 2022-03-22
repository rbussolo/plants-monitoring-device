package main

import (
	"fmt"
	"os"
	"strconv"

	d "device/src/device"
)

var users []d.User = []d.User{
	{
		Email:  "rbussolo91@gmail.com",
		ApiKey: "26a9ffa4-c373-4e2e-84a4-24561db63da5",
	},
	{
		Email:  "carlos@gmail.com",
		ApiKey: "afbbfb84-b3e5-462e-a38b-ccaf5bab5440",
	},
	{
		Email:  "test_authenticator@test.com",
		ApiKey: "655951a6-cf19-4196-a0f7-9a071889431c",
	},
}

/*
	Parameters necessary to create a new device
		* ID-DEVICE (string)

	Options parameters
		* --interval 10 				- time, in seconds, to execute call of API sending data (default value 10 seconds)
		* --api localhost:8080 	- location to send data 																(default value localhost:8080)
		* --user email					- email of user to send data 														(default value rbussolo91@gmail.com)
		* --city cuiaba					- name of city																					(default random city)
*/

func main() {
	args := os.Args[1:]

	var device d.Device = *d.NewDevice("")
	var state string = "init"
	var err error

	// Check if all parameters have been informed
	for _, arg := range args {
		if arg == "--interval" || arg == "--api" || arg == "--user" || arg == "--city" {
			state, err = setState(state, arg)

			if err != nil {
				panic(err)
			}
		} else {
			if state == "init" {
				device.Id = arg
			} else if state == "--interval" {
				device.Interval, err = strconv.Atoi(arg)

				if err != nil {
					panic(err)
				}
			} else if state == "--api" {
				device.Api = arg
			} else if state == "--user" {
				device.User = getUser(arg)
			} else if state == "--city" {
				device.City = arg
			}

			state = "init"
		}
	}

	if device.Id == "" {
		panic("error, not ID of device have been informed")
	}

	if device.Interval == 0 {
		panic("error, interval must be more than 0 seconds")
	}

	fmt.Println("Device is running in background")
	fmt.Printf("Configuration: ID=%s | interval=%ds | api=%s | user=%s\n", device.Id, device.Interval, device.Api, device.User.Email)

	// Start to send data
	device.Start()
}

func setState(state, newState string) (string, error) {
	if state == "init" {
		return newState, nil
	}

	return state, fmt.Errorf("error, have been waiting %s arg, but receive %s", state, newState)
}

func getUser(email string) d.User {
	user := d.User{
		Email: email,
	}

	for _, u := range users {
		if u.Email == email {
			return u
		}
	}

	return user
}
