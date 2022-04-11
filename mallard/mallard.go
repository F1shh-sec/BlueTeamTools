package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"time"
)

const (
	cReset  = "\033[0m"
	cRed    = "\033[31m"
	cGreen  = "\033[32m"
	cBlue   = "\033[34m"
	cCyan   = "\033[36m"
	cLBlue  = "\033[94m"
	cLCyan  = "\033[96m"
	cYellow = "\033[33m"
	cLGrey  = "\033[37m"
	cBrown  = "\033[38;5;94m"
)

var source = rand.NewSource(time.Now().UnixNano())
var randomSrc = rand.New(source)
var knownAccounts []string
var AGGRESSIVE bool

type connect struct {
	pid  []string
	name string
}

func init() {
	aggressive := flag.Bool("a", false, "Runs all commands on start")
	flag.Parse()
	AGGRESSIVE = *aggressive
	if AGGRESSIVE == true {
		fmt.Println(cBlue + "GETTING USERS WITH A SHELL:" + cReset)
		users()
		fmt.Println(cBlue + "CHANGING PASSWORDS FOR USERS TO: Sup3rS3cur3S3cret!" + cReset)
		change_passwd("Sup3rS3cur3S3cret!")
		fmt.Println(cBlue + "DISABLING SHELL ACCESS FOR USERS:" + cReset)
		disableAccounts()
	}
}

func main() {
	logo()
	mallard()
}

/**
Prints out the logo on startup
*/
func logo() {
	fmt.Println(cGreen + `   ▄▄▄▄███▄▄▄▄      ▄████████  ▄█        ▄█          ▄████████    ▄████████ ████████▄  
 ` + cGreen + `▄██▀▀▀███▀▀▀██▄   ███    ███ ███       ███         ███    ███   ███    ███ ███   ▀███ 
 ` + cGreen + `███   ███   ███   ███    ███ ███       ███         ███    ███   ███    ███ ███    ███ 
 ` + cGreen + `███   ███   ███   ███    ███ ███       ███         ███    ███  ▄███▄▄▄▄██▀ ███    ███ 
 ` + cLGrey + `███   ███   ███ ▀███████████ ███       ███       ▀███████████ ▀▀███▀▀▀▀▀   ███    ███ 
 ` + cBrown + `███   ███   ███   ███    ███ ███       ███         ███    ███ ▀███████████ ███    ███ 
 ` + cBrown + `███   ███   ███   ███    ███ ███▌    ▄ ███▌    ▄   ███    ███   ███    ███ ███   ▄███ 
  ` + cBrown + `▀█   ███   █▀    ███    █▀  █████▄▄██ █████▄▄██   ███    █▀    ███    ███ ████████▀  
                              ▀         ▀                        ███    ███` + cReset)

	messages := []string{
		"Any machine can become un-hackable if you turn it off!",
		"This is why you take snapshots.",
		"Quack",
		"Now in HD",
		"Try not to brick your box.",
		"Exiting Program.\nFuck, wrong message.",
		"Your friendly neighborhood watchduck.",
		"Sniff packets not glue.",
	}
	ranMessage := randomSrc.Intn(len(messages))
	fmt.Println(cRed + messages[ranMessage] + cReset)
}

/**
Main function that starts the watchers and command handler
*/
func mallard() {
	fmt.Println(cBlue + "Generating Initial Info Report..." + cReset)
	getInfo()
	reader := bufio.NewReader(os.Stdin)
	go watchAccounts()
	go watchConnections()
	for {
		printPrefix()
		input, _ := reader.ReadString('\n')
		commandHandle(input)
	}

}

/**
Prints the command prefix when called
*/
func printPrefix() {
	fmt.Print(cGreen + "# " + cReset)
}

/**
Command Handler
*/
func commandHandle(input string) {
	input_split := strings.Split(input, " ")
	input_trimmed := strings.TrimSpace(input_split[0])
	switch input_trimmed {
	case "exit":
		fmt.Println(cRed + "Exiting Program." + cReset)
		os.Exit(1)
	case "users":
		users()
	case "passwd":
		if len(input_split) > 1 {
			change_passwd(strings.TrimSpace(input_split[1]))
		} else {
			change_passwd("SuperSecurePassword")
		}
	case "disable":
		disableAccounts()
	case "info":
		getInfo()
	case "help":
		help()
	case "conn":
		getConnections()
	default:
		fmt.Println("Command Not Found...\n")
	}
}

/**
Watcher that prevents new account creation
*/
func watchAccounts() {
	GetInitialUsers, err := exec.Command("bash", "-c", "mapfile -t usersArray < <(awk -F\":\" '{print $1}' /etc/passwd);echo \"${usersArray[@]}\"\n").Output()
	if err != nil {
		fmt.Println(err)
	}
	knownAccounts = strings.Split(strings.TrimSpace(string(GetInitialUsers)), " ")
	for {
		getUserCommand, err := exec.Command("bash", "-c", "mapfile -t usersArray < <(awk -F\":\" '{print $1}' /etc/passwd);echo \"${usersArray[@]}\"\n").Output()
		if err != nil {
			fmt.Println(err)
		}
		getUserSplit := strings.Split(strings.TrimSpace(string(getUserCommand)), " ")
		if !reflect.DeepEqual(knownAccounts, getUserSplit) {
			if len(getUserSplit) > len(knownAccounts) {
				newusers := getUserSplit[len(knownAccounts):]
				for _, elm := range newusers {
					fmt.Println(cBlue + "\nA NEW USER HAS BEEN CREATED: " + cRed + strings.TrimSpace(string(elm)) + cReset)
					deletecommand := "userdel -f " + strings.TrimSpace(string(elm))
					stopServicesCommand := "killall -u " + strings.TrimSpace(string(elm))
					logoutCommand := "skill -kill -u " + strings.TrimSpace(string(elm))
					_, err := exec.Command("bash", "-c", stopServicesCommand).Output()
					_, err = exec.Command("bash", "-c", logoutCommand).Output()
					_, err = exec.Command("bash", "-c", deletecommand).Output()
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(cBlue + "DELETED USER: " + cRed + strings.TrimSpace(string(elm)) + cReset)
					printPrefix()
				}

			}
		}
		time.Sleep(time.Duration(500) * time.Millisecond)
	}

}

func parseConnections(connections []string) []connect {
	var foundConnections []connect
	for _, elm := range connections {
		connection := strings.Split(strings.TrimSpace(string(elm)), ":")
		pids := strings.Split(strings.TrimSpace(string(connection[0])), " ")
		serviceName := connection[1]
		newConnect := connect{pids, serviceName}
		foundConnections = append(foundConnections, newConnect)
	}
	return foundConnections
}

func watchConnections() {
	getInitConns, err := exec.Command("bash", "../scripts/getconn.sh").Output()
	if err != nil {
		fmt.Println(err)
	}
	initConnSplit := strings.Split(strings.TrimSpace(string(getInitConns)), "\n")
	// Parses the initial Connection list
	initParsed := parseConnections(initConnSplit)

	connectionMap := make(map[string][]string)
	for _, elm := range initParsed {
		connectionMap[elm.name] = elm.pid
	}

	for {
		getNewConns, err := exec.Command("bash", "../scripts/getconn.sh").Output()
		if err != nil {
			fmt.Println(err)
		}
		getConnsSplit := strings.Split(strings.TrimSpace(string(getNewConns)), "\n")
		// Parses the new connection into the array
		getConnParsed := parseConnections(getConnsSplit)

		// CHECKS TO SEE IF A NEW CONNECTION IS MADE
		// For each connection in the new command
		for _, elm := range getConnParsed {
			// Check and kill the process if its malicious

			/**
			TODO: check the prefixes and format the output print accordingly
			*/
			// Returns true on malicious and false if clean
			ismalicious := checkAndKill(elm.name, elm.pid)
			if ismalicious {
				fmt.Println(cBlue + "\nNew Connection Found: " + cRed + elm.name + cReset)
				fmt.Println(cBlue + "Killing Malicious Process: " + cRed + elm.name + cReset)
			}
			// Check if we have the name of the service in the list
			_, ok := connectionMap[elm.name]
			if ok {
				// If we do, Check to see if the pids are the same
				if !reflect.DeepEqual(connectionMap[elm.name], elm.pid) {
					// If they are not, we have a new process
					if len(connectionMap[elm.name]) > len(elm.pid) {
						if ismalicious {
							fmt.Println(cBlue + "Connection Removed: " + cRed + elm.name + cReset)
						} else {
							fmt.Println(cBlue + "\nConnection Removed: " + cRed + elm.name + cReset)
						}
						printPrefix()
						connectionMap[elm.name] = elm.pid
						initParsed = getConnParsed
					} else {
						if !ismalicious {
							fmt.Println(cBlue + "\nNew Connection Found: " + cRed + elm.name + cReset)
							printPrefix()
						}
						connectionMap[elm.name] = elm.pid
						initParsed = getConnParsed
					}
				}
			} else {
				// If the name is not in the list, We have a new process
				if !ismalicious {
					fmt.Println(cBlue + "\nNew Connection Found: " + cRed + elm.name + cReset)
					printPrefix()
				}
				// Add the new process to the list
				connectionMap[elm.name] = elm.pid
				initParsed = getConnParsed
			}
		}

		NewConnectionMap := make(map[string][]string)
		for _, elm := range getConnParsed {
			NewConnectionMap[elm.name] = elm.pid
		}

		// CHECKS TO SEE IF A CONNECTION IS REMOVED
		for _, elm := range initParsed {
			// If the new connection map has a elm of the old connection map
			_, ok := NewConnectionMap[elm.name]
			if ok == true {
				//fmt.Println("Matching Element Detected: " + elm.name)
				// Check to see if the PID maps are =
				if !reflect.DeepEqual(NewConnectionMap[elm.name], elm.pid) {
					// If they are not, Then a process was removed and we need to update the connection map
					fmt.Println(cBlue + "\nConnection Removed: " + cRed + elm.name + cReset)
					connectionMap[elm.name] = elm.pid
					printPrefix()
				}
			} else {
				// If the name is not in the list, Then the process was removed
				fmt.Println(cBlue + "\nConnection Removed: " + cRed + elm.name + cReset)
				printPrefix()
				// Add the new process to the list
				delete(connectionMap, elm.name)
			}
		}
		initParsed = getConnParsed
		time.Sleep(time.Duration(500) * time.Millisecond)
	}

}

/**
Returns true if the process has a malicious name, false if it does not.
I plan to use hashing for this at some point.
*/
func checkAndKill(name string, pids []string) bool {

	maliciousProcessNames := []string{"nc", "mimikatz", "meterpreter"}
	for _, malName := range maliciousProcessNames {
		if name == malName {
			for _, elm := range pids {
				filePathString := "lsof -p " + elm + " grep -m 1 txt | awk '{print $9}'"
				filepath, err := exec.Command("bash", "-c", filePathString).Output()

				hashString := "md5sum " + string(filepath) + " | awk '{print $1}'"
				md5hash, err := exec.Command("bash", "-c", hashString).Output()
				killProcess := "kill -9 " + strings.TrimSpace(string(elm))
				processKilled, err := exec.Command("bash", "-c", killProcess).Output()
				if err != nil {
					fmt.Println(md5hash)
					fmt.Println(processKilled)
					fmt.Println(err)
				}
				return true
			}
		}
	}
	return false
}
func getConnections() {
	fmt.Println(cBlue + "Active Connections and associated PIDs: " + cReset)
	cmd, err := exec.Command("bash", "../scripts/getconn.sh").Output()
	if err != nil {
		fmt.Println(err)
	}
	cmdSplit := strings.Split(strings.TrimSpace(string(cmd)), "\n")
	for _, elm := range cmdSplit {
		commandSplit := strings.Split(elm, ":")
		fmt.Println(cRed + commandSplit[0] + cReset + ":" + cGreen + commandSplit[1] + cReset)
	}
}

/**
Calls the get users script
*/
func users() {
	fmt.Print("Users with shell access: ")
	cmd, err := exec.Command("bash", "../scripts/getusers.sh").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(cmd))
}

/**
Calls the change password script
*/
func change_passwd(newPassword string) {
	// Makes sure the script has run permissions
	_, err := exec.Command("bash", "-c", "chmod +x ../scripts/changepasswords.sh").Output()
	if err != nil {
		fmt.Println(err)
	}
	commandString := "../scripts/changepasswords.sh " + newPassword
	cmd, err := exec.Command("bash", "-c", commandString).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(cmd))
}

/**
Disables accounts
*/
func disableAccounts() {
	cmd, err := exec.Command("bash", "../scripts/disableusers.sh").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(cmd))
}

/**
Runs a script to get a lot of useful information
*/
func getInfo() {
	cmd, err := exec.Command("bash", "../scripts/infocollect.sh").Output()
	if err != nil {
		fmt.Println(err)
	}
	curtime := time.Now()
	filename := "Info_" + string(curtime.Format("Jan _2 15:04:05.000"))
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	_, err = io.WriteString(file, string(cmd))
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
	fmt.Println(cBlue + "Created File: " + cGreen + filename + cReset + "\n")
}

/**
Prints out help
*/
func help() {
	fmt.Println(cBlue + "exit:" + cGreen + "Exits the program" + cReset)
	fmt.Println(cBlue + "users:" + cGreen + " Gets a list of all users with a shell" + cReset)
	fmt.Println(cBlue + "passwd <new password>:" + cGreen + " Changes the password of all users to a set string" + cReset)
	fmt.Println(cBlue + "disable:" + cGreen + " Disables shell access for all users with a shell" + cReset)
	fmt.Println(cBlue + "info:" + cGreen + " Gets the initial state of the machine" + cReset)
	fmt.Println(cBlue + "conn:" + cGreen + " Prints out active connections and the associated PIDs" + cReset)
	fmt.Println()
}
