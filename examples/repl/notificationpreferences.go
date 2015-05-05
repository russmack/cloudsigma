package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/statemachiner"
	"strconv"
)

type CommandGetNotifyPrefs struct {
	responseChan chan string
	promptChan   chan string
	userChan     chan string
}

type CommandSetNotifyPrefs struct {
	Contact      string `json:"contact"`
	Medium       int    `json:"medium"`
	Type         int    `json:"type"`
	Value        string `json:"value"`
	responseChan chan string
	promptChan   chan string
	userChan     chan string
}

func NewGetNotifyPrefs() *CommandGetNotifyPrefs {
	return &CommandGetNotifyPrefs{}
}

func (g *CommandGetNotifyPrefs) Start(respChan chan string, promptChan chan string, userChan chan string) {
	g.responseChan = respChan
	g.promptChan = promptChan
	g.userChan = userChan
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = g.getNotifyPrefs
	cargo := CommandGetNotifyPrefs{}
	stateMachine.Start(cargo)
}

func (g *CommandGetNotifyPrefs) getNotifyPrefs(cargs interface{}) statemachiner.StateFn {
	o := cloudsigma.NewNotificationPreferences()
	args := o.NewGet()
	fmt.Println("Username:", config.Login().Username)
	args.Username = config.Login().Username
	args.Password = config.Login().Password
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
	}
	g.responseChan <- string(resp)
	return nil
}
