package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/statemachiner"
)

type CommandGetNotifyPrefs struct {
	responseChan chan string
	promptChan   chan string
	userChan     chan string
}

type CommandSetNotifyPrefs struct {
	Contact      string `json:"contact"`
	Medium       string `json:"medium"`
	Type         string `json:"type"`
	Value        bool   `json:"value"`
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

func (m *CommandSetNotifyPrefs) Start(respChan chan string, promptChan chan string, userChan chan string) {
	m.responseChan = respChan
	m.promptChan = promptChan
	m.userChan = userChan
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.setNotifyPrefsContact
	cargo := CommandSetNotifyPrefs{}
	stateMachine.Start(cargo)
}

func (m *CommandSetNotifyPrefs) setNotifyPrefsContact(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.promptChan <- "Contact:"
	s := <-m.userChan
	c, ok := cargo.(CommandSetNotifyPrefs)
	if ok {
		c.Contact = s
	} else {
		fmt.Println("assertion not ok")
	}
	return m.setNotifyPrefsMedium(c)
}

func (m *CommandSetNotifyPrefs) setNotifyPrefsMedium(cargo interface{}) statemachiner.StateFn {
	m.promptChan <- "Medium:"
	s := <-m.userChan
	c, ok := cargo.(CommandSetNotifyPrefs)
	if ok {
		c.Medium = s
	} else {
		fmt.Println("asserton not ok")
	}
	return m.setNotifyPrefsType(c)
}

func (m *CommandSetNotifyPrefs) setNotifyPrefsType(cargo interface{}) statemachiner.StateFn {
	m.promptChan <- "Type:"
	s := <-m.userChan
	c, ok := cargo.(CommandSetNotifyPrefs)
	if ok {
		c.Type = s
	} else {
		fmt.Println("assertion not ok")
	}
	return m.setNotifyPrefsValue(c)
}

func (m *CommandSetNotifyPrefs) setNotifyPrefsValue(cargo interface{}) statemachiner.StateFn {
	//o := cloudsigma.NewNotificationPreferences()
	//args := o.NewGet()
	//client := &cloudsigma.Client{}
	//resp, err := client.Call(args)
	//if err != nil {
	//	fmt.Println("Error calling client.", err)
	//}
	m.responseChan <- "Not yet implemented"
	return nil
}
