package main

import (
	"bufio"
	"fmt"
	"github.com/russmack/cloudsigma"
	"os"
	"strings"
)

func main() {
	fmt.Println("")
	fmt.Println("Welcome to the CloudSigma IaaS REPL.")
	fmt.Println("------------------------------------")
	fmt.Println("\n\neg: cloud status")
	reploop()
}

func reploop() {
	for {
		fmt.Print("\n> ")
		in := bufio.NewReader(os.Stdin)
		userText, err := in.ReadString('\n')
		if err != nil {
			fmt.Println("Instruction not understood.")
			fmt.Println(err)
			break
		}
		instr := strings.Trim(userText, "\t \r\n")
		// TODO: cmd, err := commandParser.Parse(instr)
		/* if err != nil {
			fmt.Println("Unable to parse instruction.")
			fmt.Println(err)
			break
		} */
		args := cloudsigma.Args{}
		switch {
		case instr == "quit" || instr == "exit":
			fmt.Println("Bye!")
			os.Exit(0)
		case instr == "cloud status":
			o := cloudsigma.NewCloudStatus()
			args = o.NewGet()
		case strings.HasPrefix(instr, "current usage"):
			tokens := strings.Split(instr, " ")
			if len(tokens) < 4 {
				fmt.Println("Usage: current usage username password")
				continue
			}
			o := cloudsigma.NewCurrentUsage()
			args = o.NewGet()
			args.Username = tokens[2]
			args.Password = tokens[3]
		default:
			fmt.Println("Command not recognised.")
			continue
		}
		// TODO: resp, err := dispatcher.Dispatch(cmd)
		/* if err != nil {
			fmt.Println("Unable to execute command.")
			fmt.Println(err)
			break
		} */
		// Create a client.
		client := &cloudsigma.Client{}
		resp, err := client.Call(args)
		if err != nil {
			fmt.Println("Error calling client.", err)
			//return []byte{}
			continue
		}
		fmt.Println(string(resp))
	}
}
