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

var colorReset, colorRed, colorGreen, colorBlue = "\033[0m", "\033[31m", "\033[32m", "\033[34m"
var source = rand.NewSource(time.Now().UnixNano())
var randomSrc = rand.New(source)
var knownAccounts []string
var AGGRESSIVE bool

func init() {
	aggressive := flag.Bool("a", false, "Runs all commands on start")
	flag.Parse()
	AGGRESSIVE = *aggressive
	if AGGRESSIVE == true {
		fmt.Println(colorBlue + "GETTING USERS WITH A SHELL:" + colorReset)
		users()
		fmt.Println(colorBlue + "CHANGING PASSWORDS FOR USERS TO: Sup3rS3cur3S3cret!" + colorReset)
		change_passwd("Sup3rS3cur3S3cret!")
		fmt.Println(colorBlue + "DISABLING SHELL ACCESS FOR USERS:" + colorReset)
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
	fmt.Println(colorBlue + `   ▄▄▄▄███▄▄▄▄      ▄████████  ▄█        ▄█          ▄████████    ▄████████ ████████▄  
 ▄██▀▀▀███▀▀▀██▄   ███    ███ ███       ███         ███    ███   ███    ███ ███   ▀███ 
 ███   ███   ███   ███    ███ ███       ███         ███    ███   ███    ███ ███    ███ 
 ███   ███   ███   ███    ███ ███       ███         ███    ███  ▄███▄▄▄▄██▀ ███    ███ 
 ███   ███   ███ ▀███████████ ███       ███       ▀███████████ ▀▀███▀▀▀▀▀   ███    ███ 
 ███   ███   ███   ███    ███ ███       ███         ███    ███ ▀███████████ ███    ███ 
 ███   ███   ███   ███    ███ ███▌    ▄ ███▌    ▄   ███    ███   ███    ███ ███   ▄███ 
  ▀█   ███   █▀    ███    █▀  █████▄▄██ █████▄▄██   ███    █▀    ███    ███ ████████▀  
                              ▀         ▀                        ███    ███` + colorReset)

	messages := []string{
		"Any machine can become unhackable if you turn it off!",
		"This is why you take snapshots.",
		"Quack",
		"Now in HD",
		"Try not to brick your box.",
		"Exiting Program.\nFuck, Wrong message."}
	ranMessage := randomSrc.Intn(len(messages))
	fmt.Println(colorRed + messages[ranMessage] + colorReset)
}

/**
Main function that starts the watchers and command handler
*/
func mallard() {
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
	fmt.Print(colorGreen + "# " + colorReset)
}

/**
Command Handler
*/
func commandHandle(input string) {
	input_split := strings.Split(input, " ")
	input_trimmed := strings.TrimSpace(input_split[0])
	switch input_trimmed {
	case "exit":
		fmt.Println(colorRed + "Exiting Program." + colorReset)
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
					fmt.Println(colorBlue + "\nA NEW USER HAS BEEN CREATED: " + colorRed + strings.TrimSpace(string(elm)) + colorReset)
					deletecommand := "userdel -f " + strings.TrimSpace(string(elm))
					stopServicesCommand := "killall -u " + strings.TrimSpace(string(elm))
					logoutCommand := "skill -kill -u " + strings.TrimSpace(string(elm))
					_, err := exec.Command("bash", "-c", stopServicesCommand).Output()
					_, err = exec.Command("bash", "-c", logoutCommand).Output()
					_, err = exec.Command("bash", "-c", deletecommand).Output()
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(colorBlue + "DELETED USER: " + colorRed + strings.TrimSpace(string(elm)) + colorReset)
					printPrefix()
				}

			}
		}
		time.Sleep(time.Duration(500) * time.Millisecond)
	}

}

type connect struct {
	pid  []string
	name string
}

func parseConnections(connections []string) []connect {
	var foundConnections []connect
	for _, elm := range connections {
		connection := strings.Split(string(elm), ":")
		pids := strings.Split(string(connection[0]), " ")
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
		for _, elm := range getConnParsed {
			if _, ok := connectionMap[elm.name]; ok {
				if !reflect.DeepEqual(connectionMap[elm.name], elm.pid) {
					fmt.Println("New Connection Found!")
					connectionMap[elm.name] = elm.pid
				}
			} else {
				fmt.Println("New Connection Found!")
				connectionMap[elm.name] = elm.pid
			}
		}
		/**
		TODO: Write a function to compare getConn Parsed and InitParsed. Potentially rewrite the structure to map
		TODO: the service to the pids. Then check to see if a value is in the map.
		*/

		time.Sleep(time.Duration(500) * time.Millisecond)
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
	file, err := os.Create("InitialInfo.txt")
	if err != nil {
		fmt.Println(err)
	}
	_, err = io.WriteString(file, string(cmd))
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
}

/**
Prints out help
*/
func help() {
	fmt.Println(colorBlue + "exit:" + colorGreen + "Exits the program" + colorReset)
	fmt.Println(colorBlue + "users:" + colorGreen + " Gets a list of all users with a shell" + colorReset)
	fmt.Println(colorBlue + "passwd <new password>:" + colorGreen + " Changes the password of all users to a set string" + colorReset)
	fmt.Println(colorBlue + "disable:" + colorGreen + " Disables shell access for all users with a shell" + colorReset)
	fmt.Println(colorBlue + "info:" + colorGreen + " Gets the initial state of the machine" + colorReset)
	fmt.Println()
}
