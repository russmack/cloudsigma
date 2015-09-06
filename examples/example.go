package main

import (
	"bufio"
	"fmt"
	"github.com/russmack/cloudsigma"
	"os"
	"strconv"
	"strings"
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
		//"DownloadImage": getImage,
		"ApiUrls":                     getApiUrls,
		"CloudStatus":                 getCloudStatus,
		"CloudStatusXML":              getCloudStatusXml,
		"Locations":                   getLocations,
		"Capbilites":                  getCapabilities,
		"CurrentUsage":                getCurrentUsage,
		"BurstUsage":                  getBurstUsage,
		"DailyBurstUsage":             getDailyBurstUsage,
		"Usage":                       getUsage,
		"Servers":                     getServers,
		"Drives":                      getDrives,
		"Snapshots":                   getSnapshots,
		"Vlans":                       getVlans,
		"Ips":                         getIps,
		"Acls":                        getAcls,
		"ServerCreate":                getServerCreate,
		"ServerDelete":                getServerDelete,
		"ServerShutdown":              getServerShutdown,
		"ServerStop":                  getServerStop,
		"ServerStart":                 getServerStart,
		"DriveCreate":                 getDriveCreate,
		"DriveDelete":                 getDriveDelete,
		"Profile":                     getProfile,
		"Balance":                     getBalance,
		"Pricing":                     getPricing,
		"Discounts":                   getDiscounts,
		"Transactions":                getTransactions,
		"Licenses":                    getLicenses,
		"Subscriptions":               getSubscriptions,
		"NotificationContacts":        getNotificationContacts,
		"GetNotificationPreferences":  getNotificationPreferences,
		"EditNotificationPreferences": editNotificationPreference,
	}

	for k, v := range calls {
		fmt.Println("\nRequesting " + k + "...")
		resp := v(location, username, password)
		fmt.Println("Response from " + k + ":\n")
		fmt.Println(string(resp) + "\n")
	}
}

// confirmProceed provides an opportunity for the user to avoid write operations.
func confirmProceed() bool {
	fmt.Println("Are you sure you want to do this? ( y/n )")
	in := bufio.NewReader(os.Stdin)
	answer, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading confirmation.", err)
		os.Exit(1)
	}
	answer = strings.Trim(answer, "\t \r\n")
	if answer == "y" {
		return true
	} else {
		return false
	}
}

func getApiUrls(location string, username string, password string) []byte {
	// Create an ApiUrls.
	o := cloudsigma.NewApiUrls()
	args := o.NewList()
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

// Cloud

func getCloudStatus(location string, username string, password string) []byte {
	// Create a CloudStatus.
	o := cloudsigma.NewCloudStatus()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func getCloudStatusXml(location string, username string, password string) []byte {
	// Create a CloudStatus.
	o := cloudsigma.NewCloudStatus()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location
	args.Format = "xml"

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func getLocations(location string, username string, password string) []byte {
	// Create a Locations.
	o := cloudsigma.NewLocations()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func getCapabilities(location string, username string, password string) []byte {
	// Create a Capabilities.
	o := cloudsigma.NewCapabilities()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte(err.Error())
	}
	return resp
}

// Usage

func getCurrentUsage(location string, username string, password string) []byte {
	// Create a CurrentUsage.
	o := cloudsigma.NewCurrentUsage()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getBurstUsage(location string, username string, password string) []byte {
	// Create a BurstUsage.
	o := cloudsigma.NewBurstUsage()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getDailyBurstUsage(location string, username string, password string) []byte {
	// Create a DailyBurstUsage.
	o := cloudsigma.NewDailyBurstUsage()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func getUsage(location string, username string, password string) []byte {
	// Create a Usage.
	o := cloudsigma.NewUsage()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

// Infrastructure

func getServers(location string, username string, password string) []byte {
	// Create a Servers.
	o := cloudsigma.NewServers()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func getDrives(location string, username string, password string) []byte {
	// Create a Drives.
	o := cloudsigma.NewDrives()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getDriveCreate(location string, username string, password string) []byte {
	if !confirmProceed() {
		return []byte("Skipping.")
	}

	newDrives := []cloudsigma.DriveRequest{
		cloudsigma.DriveRequest{"disk", "Example Drive", 6174015488},
	}
	// Create a Drives.
	o := cloudsigma.NewDrives()
	args := o.NewCreate(newDrives)
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getDriveDelete(location string, username string, password string) []byte {
	if !confirmProceed() {
		return []byte("Skipping.")
	}

	// Create a Drives.
	o := cloudsigma.NewDrives()
	args := o.NewDelete("")
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getSnapshots(location string, username string, password string) []byte {
	// Create a Snapshots.
	o := cloudsigma.NewSnapshots()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getImage(location string, username string, password string, driveUuid string) []byte {
	if !confirmProceed() {
		return []byte("Skipping.")
	}

	// Create an Images.
	o := cloudsigma.NewImages()
	args := o.NewDownload(driveUuid)
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	//resp, err := client.Call(nil, args)
	resp, err := client.Download(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getVlans(location string, username string, password string) []byte {
	// Create a Vlans.
	o := cloudsigma.NewVlans()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getIps(location string, username string, password string) []byte {
	// Create an Ips.
	o := cloudsigma.NewIps()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getAcls(location string, username string, password string) []byte {
	// Create an Acls.
	o := cloudsigma.NewAcls()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getServerCreate(location string, username string, password string) []byte {
	if !confirmProceed() {
		return []byte("Skipping.")
	}

	newServers := []cloudsigma.ServerRequest{
		cloudsigma.ServerRequest{"Example Server 0", 1000, 536870912, "vncP455word"},
	}
	// Add a few more.
	for i := 1; i <= 3; i++ {
		newServers = append(newServers,
			cloudsigma.ServerRequest{
				"Example Server " + strconv.Itoa(i), 1000, 536870912, "vncP455word"})
	}
	// Create a Servers.
	o := cloudsigma.NewServers()
	args := o.NewCreate(newServers)
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getServerDelete(location string, username string, password string) []byte {
	if !confirmProceed() {
		return []byte("Skipping.")
	}

	// Create a Servers.
	o := cloudsigma.NewServers()
	args := o.NewDelete("")
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getServerStart(location string, username string, password string) []byte {
	if !confirmProceed() {
		return []byte("Skipping.")
	}

	// Create a Servers.
	o := cloudsigma.NewServers()
	args := o.NewStart("")
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getServerStop(location string, username string, password string) []byte {
	if !confirmProceed() {
		return []byte("Skipping.")
	}

	// Create a Servers.
	o := cloudsigma.NewServers()
	args := o.NewStop("")
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getServerShutdown(location string, username string, password string) []byte {
	if !confirmProceed() {
		return []byte("Skipping.")
	}

	// Create a Servers.
	o := cloudsigma.NewServers()
	args := o.NewShutdown("")
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

// Billing

func getProfile(location string, username string, password string) []byte {
	// Create a Profile.
	o := cloudsigma.NewProfile()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getBalance(location string, username string, password string) []byte {
	// Create a Balance.
	o := cloudsigma.NewBalance()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getPricing(location string, username string, password string) []byte {
	// Create a Pricing.
	o := cloudsigma.NewPricing()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getDiscounts(location string, username string, password string) []byte {
	// Create a Discounts.
	o := cloudsigma.NewDiscounts()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getTransactions(location string, username string, password string) []byte {
	// Create a Transactions.
	o := cloudsigma.NewTransactions()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getSubscriptions(location string, username string, password string) []byte {
	// Create a Subscriptions.
	o := cloudsigma.NewSubscriptions()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}
func getLicenses(location string, username string, password string) []byte {
	// Create a Licenses.
	o := cloudsigma.NewLicenses()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

// Notifications

func getNotificationContacts(location string, username string, password string) []byte {
	// Create a NotificationContacts
	o := cloudsigma.NewNotificationContacts()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func getNotificationPreferences(location string, username string, password string) []byte {
	// Create a NotificationPreferences
	o := cloudsigma.NewNotificationPreferences()
	args := o.NewList()
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error calling client.", err)
		return []byte{}
	}
	return resp
}

func editNotificationPreference(location string, username string, password string) []byte {
	if !confirmProceed() {
		return []byte("Skipping.")
	}

	// Create the object to post as body.
	p := cloudsigma.Preference{
		Contact: "/api/2.0/notification_contacts/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/",
		Medium:  "email",
		Type:    "trial_expire",
		Value:   false,
	}

	// Create a NotificationPreferences
	o := cloudsigma.NewNotificationPreferences()
	args := o.NewEdit(p)
	args.Username = username
	args.Password = password
	args.Location = location

	// Create a client.
	client := &cloudsigma.Client{}
	resp, err := client.Call(nil, args)
	if err != nil {
		fmt.Println("Error, but only because this demo has an invalid contact id.", err)
		return []byte{}
	}
	return resp
}
