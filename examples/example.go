package main

import (
	"encoding/json"
	"fmt"
	"github.com/russmack/cloudsigma"
	"io/ioutil"
)

type Config struct {
	Username string
	Password string
}

func main() {
	cfg, _ := getLogin()
	fmt.Println("Using login: ", cfg.Username)
	username := cfg.Username
	password := cfg.Password

	calls := map[string]func(string, string) []byte{
		"ApiUrls":                    getApiUrls,
		"CloudStatus":                getCloudStatus,
		"CurrentUsage":               getCurrentUsage,
		"Usage":                      getUsage,
		"Servers":                    getServers,
		"Drives":                     getDrives,
		"Balance":                    getBalance,
		"Subscriptions":              getSubscriptions,
		"NotificationContacts":       getNotificationContacts,
		"GetNotificationPreferences": getNotificationPreferences,
		"SetNotificationPreferences": setNotificationPreference,
	}

	for k, v := range calls {
		resp := v(username, password)
		fmt.Println("Response (" + k + ") :")
		fmt.Println(string(resp) + "\n")
	}
}

func getLogin() (Config, error) {
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	cfg := Config{}
	err = json.Unmarshal(f, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}

func getApiUrls(username string, password string) []byte {
	// Create an ApiUrls.
	o := cloudsigma.NewApiUrls()
	args := o.NewGet()

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.")
		panic(err)
	}
	return resp
}

func getCloudStatus(username string, password string) []byte {
	// Create a CloudStatus.
	o := cloudsigma.NewCloudStatus()
	args := o.NewGet()
	args.Username = username
	args.Password = password

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.")
		panic(err)
	}
	return resp
}

func getCurrentUsage(username string, password string) []byte {
	// Create a CurrentUsage.
	o := cloudsigma.NewCurrentUsage()
	args := o.NewGet()
	args.Username = username
	args.Password = password

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.")
		panic(err)
	}
	return resp
}

func getUsage(username string, password string) []byte {
	// Create a Usage.
	o := cloudsigma.NewUsage()
	args := o.NewGet()
	args.Username = username
	args.Password = password

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.")
		panic(err)
	}
	return resp
}

func getServers(username string, password string) []byte {
	// Create a Servers.
	o := cloudsigma.NewServers()
	args := o.NewGet()
	args.Username = username
	args.Password = password

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.")
		panic(err)
	}
	return resp
}

func getDrives(username string, password string) []byte {
	// Create a Drives.
	o := cloudsigma.NewDrives()
	args := o.NewGet()
	args.Username = username
	args.Password = password

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.")
		panic(err)
	}
	return resp
}

func getBalance(username string, password string) []byte {
	// Create a Balance.
	o := cloudsigma.NewBalance()
	args := o.NewGet()
	args.Username = username
	args.Password = password

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.")
		panic(err)
	}
	return resp
}

func getSubscriptions(username string, password string) []byte {
	// Create a Subscriptions.
	o := cloudsigma.NewSubscriptions()
	args := o.NewGet()
	args.Username = username
	args.Password = password

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.")
		panic(err)
	}
	return resp
}

func getNotificationContacts(username string, password string) []byte {
	// Create a NotificationContacts
	o := cloudsigma.NewNotificationContacts()
	args := o.NewGet()
	args.Username = username
	args.Password = password

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.")
		panic(err)
	}
	return resp
}

func getNotificationPreferences(username string, password string) []byte {
	// Create a NotificationPreferences
	o := cloudsigma.NewNotificationPreferences()
	args := o.NewGet()
	args.Username = username
	args.Password = password

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.")
		panic(err)
	}
	return resp
}

func setNotificationPreference(username string, password string) []byte {
	// Create the object to post as body.
	p := cloudsigma.Preference{
		Contact: "/api/2.0/notification_contacts/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/",
		Medium:  "email",
		Type:    "trial_expire",
		Value:   false,
	}

	// Create a NotificationPreferences
	o := cloudsigma.NewNotificationPreferences()
	args := o.NewSet(p)
	args.Username = username
	args.Password = password

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.")
		panic(err)
	}
	return resp
}
