package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	ps "github.com/gorillalabs/go-powershell"
	"github.com/gorillalabs/go-powershell/backend"
	"github.com/gorillalabs/go-powershell/middleware"
)

func main() {
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
	config.ComputerName = "t1300wks0320"

	// create a new shell by wrapping the existing one in the session middleware
	session, err := middleware.NewSession(shell, config)
	if err != nil {
		panic(err)
	}
	defer session.Exit() // will also close the underlying ps shell!

	// everything run via the session is run on the remote machine
	comsys, stderr,err := session.Execute("Get-WmiObject -Class win32_computersystem | Select-Object username, PowerSupplyState, Model")
	sn, stderr, err := session.Execute("Get-WmiObject -Class Win32_BIOS | Select-Object SerialNumber, Manufacturer")
	opsys, stderr, err := session.Execute("Get-WMIObject -Class Win32_OperatingSystem | Select-Object caption, LastBootUpTime, Status, FreePhysicalMemory")
	usb, stderr, err := session.Execute("Get-WmiObject -Class Win32_UsbHub | Select-Object Name, Status")
	fan, stderr, err := session.Execute("Get-WmiObject -class Win32_Fan | Select-Object Name, Status")
	prt, stderr, err := session.Execute("Get-WmiObject -Class Win32_Printer | Select-Object Name, DriverName, SpoolEnabled, PortName")
	drvstat, stderr, err := session.Execute("Get-WmiObject -Class Win32_DiskDrive | Select-Object Caption, Name, InterfaceType, Status")
	

	}

	fmt.Println("")
	fmt.Println ("Computer Info: ")
	temp := strings.Split(opsys, "\r\n")
	for _,element := range temp {
		if (strings.HasPrefix(element, "caption") || strings.HasPrefix(element, "Status")){
			fmt.Println(element)
		}
	}
	
	temp1 := strings.Split(comsys, "\r\n")
	for _,element := range temp1 {
		if (strings.HasPrefix(element, "username") || strings.HasPrefix(element, "Model")){
			fmt.Println(element)
		}
	}
	
	
	temp2 := strings.Split(sn, "\r\n")
	for _,element := range temp2 {
		if (strings.HasPrefix(element, "SerialNumber") || strings.HasPrefix(element, "Manufacturer")){
			fmt.Println(element)
		}
	}
	
	fmt.Println("")
	fmt.Println ("Drive Info: ")
	temp6 := strings.Split(drvstat, "\r\n")
	for _,element := range temp6 {
		if (strings.HasPrefix(element, "Caption") || strings.HasPrefix(element, "InterfaceType") || strings.HasPrefix(element, "Status")){
			fmt.Println(element)
		}
	}
	
	fmt.Println("")
	fmt.Println ("Fan Info: ")
	temp3 := strings.Split(fan, "\r\n")
	for _,element := range temp3 {
		if (strings.HasPrefix(element, "Name") || strings.HasPrefix(element, "Status")){
			fmt.Println(element)
		}
	}
	
	fmt.Println("")
	fmt.Println ("Printer Info: ")
	temp4 := strings.Split(prt, "\r\n")
	for _,element := range temp4 {
		if (strings.HasPrefix(element, "Name") || strings.HasPrefix(element, "DriverName") || strings.HasPrefix(element, "SpoolEnabled") || strings.HasPrefix(element, "PortName")){
			fmt.Println(element)
		}
	}
	fmt.Println("")
	fmt.Println ("USB Info: ")
	temp5 := strings.Split(usb, "\r\n")
	for _,element := range temp5 {
		if (strings.HasPrefix(element, "Name") || strings.HasPrefix(element, "Status")){
			fmt.Println(element)
		}
	}


	
}