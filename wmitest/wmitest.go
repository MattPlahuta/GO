package main

import (
	"fmt"
	//"strings"
	"os"
	"regexp"
	"bufio"
	ps "github.com/gorillalabs/go-powershell"
	"github.com/gorillalabs/go-powershell/backend"
	"github.com/gorillalabs/go-powershell/middleware"
)

func main() {
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	var task_n string
	var task_s string
	var ar[]
	

	// enter hostname
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	// choose a backend
	back := &backend.Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}

	// prepare remote session configuration
	config := middleware.NewSessionConfig()
	config.ComputerName = "t2300wks0250"

	// create a new shell by wrapping the existing one in the session middleware
	session, err := middleware.NewSession(shell, config)
	if err != nil {
		panic(err)
	}
	defer session.Exit() // will also close the underlying ps shell!

	// everything run via the session is run on the remote machine
	usb, stderr, err := session.Execute("Get-WmiObject -Class Win32_UsbHub | Select-Object Name, Status")

	_ = stderr
	if err != nil {
		panic(err)
	}

	fmt.Println(usb)
	
	temp5 := strings.Split(usb, "\r\n")
	for _,element := range temp5 {
		if (strings.HasPrefix(element, "Name") || strings.HasPrefix(element, "Status")){
			fmt.Println(element)
		}
	}
	
	
	
}