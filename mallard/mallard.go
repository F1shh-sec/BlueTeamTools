package main

import (
	"bufio"
	"fmt"
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

func main() {
	logo()
	mallard()
}

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

func mallard() {
	reader := bufio.NewReader(os.Stdin)
	go watchAccounts()
	for {
		printPrefix()
		input, _ := reader.ReadString('\n')
		commandHandle(input)
	}

}

func printPrefix() {
	fmt.Print(colorGreen + "# " + colorReset)
}

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
				fmt.Println(newusers)
				for _, elm := range newusers {
					fmt.Println("Deleting User: " + string(elm))
					deletecommand := "userdel -z -r -f " + string(elm)
					fmt.Println(deletecommand)
					getUserCommand, err := exec.Command("bash", "-c", deletecommand).Output()
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(string(getUserCommand))
				}

			}
		}
		time.Sleep(time.Duration(5000) * time.Millisecond)
	}

}
func commandHandle(input string) {
	input_split := strings.Split(input, " ")
	input_trimmed := strings.TrimSpace(input_split[0])
	switch input_trimmed {
	case "exit":
		fmt.Println(colorRed + "Exiting Program." + colorReset)
		os.Exit(1)
	case "users":
		users()
	case "passwds":
		change_passwd()
	case "disableusers":
		disableAccounts()
	default:
		fmt.Println("Command Not Found...")
	}
}

func users() {
	cmd, err := exec.Command("bash", "../scripts/getusers.sh").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(cmd))
}

func change_passwd() {
	cmd, err := exec.Command("bash", "./changepasswords").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(cmd))
}

func disableAccounts() {
	cmd, err := exec.Command("bash", "../scripts/disableusers.sh").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(cmd))
}
