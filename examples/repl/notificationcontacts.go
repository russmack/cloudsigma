package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/statemachiner"
)

type CommandGetNotifyContacts struct {
	responseChan chan string
	promptChan   chan string
	userChan     chan string
}

type CommandSetNotifyContacts struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
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
	stateMachine.StartState = m.setNotifyContactsEmail
	cargo := CommandSetNotifyContacts{}
	stateMachine.Start(cargo)
}

func (m *CommandSetNotifyContacts) setNotifyContactsEmail(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.promptChan <- "Email:"
	s := <-m.userChan
	c, ok := cargo.(CommandSetNotifyContacts)
	if ok {
		c.Email = s
	} else {
		fmt.Println("assertion not ok")
	}
	return m.setNotifyContactsName(c)
}

func (m *CommandSetNotifyContacts) setNotifyContactsName(cargo interface{}) statemachiner.StateFn {
	m.promptChan <- "Name:"
	s := <-m.userChan
	c, ok := cargo.(CommandSetNotifyContacts)
	if ok {
		c.Name = s
	} else {
		fmt.Println("asserton not ok")
	}
	return m.setNotifyContactsPhone(c)
}

func (m *CommandSetNotifyContacts) setNotifyContactsPhone(cargo interface{}) statemachiner.StateFn {
	m.promptChan <- "Phone:"
	s := <-m.userChan
	c, ok := cargo.(CommandSetNotifyContacts)
	if ok {
		c.Phone = s
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
