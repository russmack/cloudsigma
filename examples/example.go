package main

import (
	//"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"os"
	"private/russmack/cloudsigma"
	//"strings"
)

type Config struct {
	Username string
	Password string
}

func main() {
	cfg, _ := getLogin()
	fmt.Println("Using login: ", cfg)
	username := cfg.Username
	password := cfg.Password

	//resp := cloudsigma.GetApiUrls()
	//resp := cloudsigma.GetCloudStatus()
	//resp := cloudsigma.GetUser(username, password)
	//resp := cloudsigma.GetUsage()
	//resp := cloudsigma.GetCurrentUsage()
	//resp := cloudsigma.GetServers()
	//resp := cloudsigma.GetProfile(username, password)
	//resp := cloudsigma.GetNotificationContacts()
	//resp := cloudsigma.GetDrives()

	// Get balance.
	resp := getBalance(username, password)
	fmt.Println("Response:")
	fmt.Println(string(resp))

	// Get all notification preferences.
	resp = getNotificationPreferences(username, password)
	fmt.Println("Response:")
	fmt.Println(string(resp))

	// Set a notification preference.
	resp = setNoticationPreference(username, password)
	fmt.Println("Response:")
	fmt.Println(string(resp))
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

func setNoticationPreference(username string, password string) []byte {
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
