package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/statemachiner"
)

type CommandBalance struct {
	responseChan chan string
	promptChan   chan string
	userChan     chan string
}

func NewBalance() *CommandBalance {
	return &CommandBalance{}
}

func (m *CommandBalance) Start(respChan chan string, promptChan chan string, userChan chan string) {
	m.responseChan = respChan
	m.promptChan = promptChan
	m.userChan = userChan
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getBalance
	cargo := CommandBalance{}
	stateMachine.Start(cargo)
}

func (m *CommandBalance) getBalance(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewBalance()
	args := o.NewGet()
	fmt.Println("Username:", config.Login().Username)
	args.Username = config.Login().Username
	args.Password = config.Login().Password
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
	}
	m.responseChan <- string(resp)
	return nil
}
