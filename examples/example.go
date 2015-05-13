package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"os"
)

func main() {
	config := cloudsigma.NewConfig()
	_, err := config.Load()
	if err != nil {
		fmt.Println("Unable to load config.", err)
		os.Exit(1)
	}
	login := config.Login()
	fmt.Println("\nUsing login: ", login.Username, "\n")
	username := login.Username
	password := login.Password
	location := "zrh"

	calls := map[string]func(string, string, string) []byte{
		"ApiUrls":                    getApiUrls,
		"CloudStatus":                getCloudStatus,
		"CloudStatusXML":             getCloudStatusXml,
		"Locations":                  getLocations,
		"Capbilites":                 getCapabilities,
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
		fmt.Println("\nRequesting " + k + "...")
		resp := v(location, username, password)
		fmt.Println("Response from " + k + ":\n")
		fmt.Println(string(resp) + "\n")
	}
}

func getApiUrls(location string, username string, password string) []byte {
	// Create an ApiUrls.
	o := cloudsigma.NewApiUrls()
	args := o.NewGet()
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func getCloudStatus(location string, username string, password string) []byte {
	// Create a CloudStatus.
	o := cloudsigma.NewCloudStatus()
	args := o.NewGet()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func getCloudStatusXml(location string, username string, password string) []byte {
	// Create a CloudStatus.
	o := cloudsigma.NewCloudStatus()
	args := o.NewGet()
	args.Username = username
	args.Password = password
	args.Location = location
	args.Format = "xml"

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func getLocations(location string, username string, password string) []byte {
	// Create a Locations.
	o := cloudsigma.NewLocations()
	args := o.NewGet()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func getCapabilities(location string, username string, password string) []byte {
	// Create a Capabilities.
	o := cloudsigma.NewCapabilities()
	args := o.NewGet()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte(err.Error())
	}
	return resp
}

func getCurrentUsage(location string, username string, password string) []byte {
	// Create a CurrentUsage.
	o := cloudsigma.NewCurrentUsage()
	args := o.NewGet()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func getUsage(location string, username string, password string) []byte {
	// Create a Usage.
	o := cloudsigma.NewUsage()
	args := o.NewGet()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func getServers(location string, username string, password string) []byte {
	// Create a Servers.
	o := cloudsigma.NewServers()
	args := o.NewGet()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func getDrives(location string, username string, password string) []byte {
	// Create a Drives.
	o := cloudsigma.NewDrives()
	args := o.NewGet()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func getBalance(location string, username string, password string) []byte {
	// Create a Balance.
	o := cloudsigma.NewBalance()
	args := o.NewGet()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func getSubscriptions(location string, username string, password string) []byte {
	// Create a Subscriptions.
	o := cloudsigma.NewSubscriptions()
	args := o.NewGet()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func getNotificationContacts(location string, username string, password string) []byte {
	// Create a NotificationContacts
	o := cloudsigma.NewNotificationContacts()
	args := o.NewGet()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func getNotificationPreferences(location string, username string, password string) []byte {
	// Create a NotificationPreferences
	o := cloudsigma.NewNotificationPreferences()
	args := o.NewGet()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func setNotificationPreference(location string, username string, password string) []byte {
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
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
