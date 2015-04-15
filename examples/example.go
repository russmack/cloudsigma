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
	fmt.Println("Using login: ", cfg)
	username := cfg.Username
	password := cfg.Password

	// TODO:
	//resp := cloudsigma.GetApiUrls()
	//resp := cloudsigma.GetCloudStatus()
	//resp := cloudsigma.GetUser(username, password)
	//resp := cloudsigma.GetUsage()
	//resp := cloudsigma.GetCurrentUsage()
	//resp := cloudsigma.GetServers()
	//resp := cloudsigma.GetProfile(username, password)
	//resp := cloudsigma.GetNotificationContacts()
	//resp := cloudsigma.GetDrives()

	// Get cloudstatus.
	resp := getCloudStatus(username, password)
	fmt.Println("Response:")
	fmt.Println(string(resp))

	// Get servers.
	resp = getServers(username, password)
	fmt.Println("Response:")
	fmt.Println(string(resp))

	// Get balance.
	resp = getBalance(username, password)
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
