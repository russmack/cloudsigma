package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/statemachiner"
	"strconv"
)

type CommandGetNotifyContacts struct {
	responseChan chan string
	promptChan   chan string
	userChan     chan string
}

type CommandSetNotifyContacts struct {
	Contact      string `json:"contact"`
	Medium       int    `json:"medium"`
	Type         int    `json:"type"`
	Value        string `json:"value"`
	responseChan chan string
	promptChan   chan string
	userChan     chan string
}

func NewGetNotifyContacts() *CommandGetNotifyContacts {
	return &CommandGetNotifyContacts{}
}

func (g *CommandGetNotifyContacts) Start(respChan chan string, promptChan chan string, userChan chan string) {
	g.responseChan = respChan
	g.promptChan = promptChan
	g.userChan = userChan
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = g.getNotifyContacts
	cargo := CommandGetNotifyContacts{}
	stateMachine.Start(cargo)
}

func (g *CommandGetNotifyContacts) getNotifyContacts(cargs interface{}) statemachiner.StateFn {
	o := cloudsigma.NewNotificationContacts()
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

func (m *CommandSetNotifyContacts) Start(respChan chan string, promptChan chan string, userChan chan string) {
	m.responseChan = respChan
	m.promptChan = promptChan
	m.userChan = userChan
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.setNotifyContactsContact
	cargo := CommandSetNotifyContacts{}
	stateMachine.Start(cargo)
}

func (m *CommandSetNotifyContacts) setNotifyContactsContact(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.promptChan <- "Contact:"
	n := <-m.userChan
	c, ok := cargo.(CommandSetNotifyContacts)
	if ok {
		c.Contact = n
	} else {
		fmt.Println("assertion not ok")
	}
	return m.setNotifyContactsMedium(c)
}

func (m *CommandSetNotifyContacts) setNotifyContactsMedium(cargo interface{}) statemachiner.StateFn {
	m.promptChan <- "Medium:"
	s := <-m.userChan
	c, ok := cargo.(CommandSetNotifyContacts)
	if ok {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("this should be a request to re-enter info")
		} else {
			c.Medium = n
		}
	} else {
		fmt.Println("asserton not ok")
	}
	return m.setNotifyContactsType(c)
}

func (m *CommandSetNotifyContacts) setNotifyContactsType(cargo interface{}) statemachiner.StateFn {
	m.promptChan <- "Type:"
	s := <-m.userChan
	c, ok := cargo.(CommandSetNotifyContacts)
	if ok {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("this should be a request to re-enter info")
		} else {
			c.Type = n
		}
	} else {
		fmt.Println("assertion not ok")
	}
	return m.setNotifyContactsValue(c)
}

func (m *CommandSetNotifyContacts) setNotifyContactsValue(cargo interface{}) statemachiner.StateFn {
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
