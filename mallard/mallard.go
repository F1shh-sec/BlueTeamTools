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
		users()
		change_passwd()
		disableAccounts()
	}
}

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
	getInfo()
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
		change_passwd()
	case "disable":
		disableAccounts()
	default:
		fmt.Println("Command Not Found...\n")
	}
}

func users() {
	fmt.Print("Users with shell access: ")
	cmd, err := exec.Command("bash", "../scripts/getusers.sh").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(cmd))
}

func change_passwd() {
	cmd, err := exec.Command("bash", "./changepasswords.sh").Output()
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
func help() {
	fmt.Println("exit: Exits the program")
	fmt.Println("users: Gets a list of all users with a shell")
	fmt.Println("passwd: Changes the password of all users to a set string")
	fmt.Println("disable: Disables shell access for all users with a shell")
}
