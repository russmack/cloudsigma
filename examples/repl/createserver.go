package main

import (
	"fmt"
	"github.com/russmack/statemachiner"
	"strconv"
)

type CommandCreateServer struct {
	Name         string `json:"name"`
	Cpu          int    `json:"cpu"`
	Memory       int    `json:"mem"`
	VncPassword  string `json:"vnc_password"`
	responseChan chan string
	promptChan   chan string
	userChan     chan string
}

func NewCreateServer() *CommandCreateServer {
	return &CommandCreateServer{}
}

func (m *CommandCreateServer) Start(respChan chan string, promptChan chan string, userChan chan string) {
	m.responseChan = respChan
	m.promptChan = promptChan
	m.userChan = userChan
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.createServerName
	cargo := CommandCreateServer{}
	stateMachine.Start(cargo)
}

func (m *CommandCreateServer) createServerName(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.promptChan <- "Name:"
	n := <-m.userChan
	c, ok := cargo.(CommandCreateServer)
	if ok {
		c.Name = n
	} else {
		fmt.Println("assertion not ok")
	}
	return m.createServerCpu(c)
}

func (m *CommandCreateServer) createServerCpu(cargo interface{}) statemachiner.StateFn {
	m.promptChan <- "CPU:"
	s := <-m.userChan
	c, ok := cargo.(CommandCreateServer)
	if ok {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("this should be a request to re-enter info")
		} else {
			c.Cpu = n
		}
	} else {
		fmt.Println("asserton not ok")
	}
	return m.createServerMemory(c)
}

func (m *CommandCreateServer) createServerMemory(cargo interface{}) statemachiner.StateFn {
	m.promptChan <- "Memory:"
	s := <-m.userChan
	c, ok := cargo.(CommandCreateServer)
	if ok {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("this should be a request to re-enter info")
		} else {
			c.Memory = n
		}
	} else {
		fmt.Println("assertion not ok")
	}
	return m.createServerVncPassword(c)
}

func (m *CommandCreateServer) createServerVncPassword(cargo interface{}) statemachiner.StateFn {
	//o := cloudsigma.NewCreateServer()
	//args := o.NewGet()
	//client := &cloudsigma.Client{}
	//resp, err := client.Call(args)
	//if err != nil {
	//	fmt.Println("Error calling client.", err)
	//}
	m.responseChan <- "Not yet implemented"
	return nil
}
