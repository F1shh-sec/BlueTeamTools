package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

var colorReset, colorRed, colorGreen, colorBlue = "\033[0m", "\033[31m", "\033[32m", "\033[34m"
var source = rand.NewSource(time.Now().UnixNano())
var randomSrc = rand.New(source)

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
		"Try not to brick your box."}
	ranMessage := randomSrc.Intn(len(messages))
	fmt.Println(colorRed + messages[ranMessage] + colorReset)
}

func mallard() {
	reader := bufio.NewReader(os.Stdin)
	for {
		printPrefix()
		input, _ := reader.ReadString('\n')
		commandHandle(input)
	}

}

func printPrefix() {
	fmt.Print(colorGreen + "# " + colorReset)
}

func commandHandle(input string) {
	input_split := strings.Split(input, " ")
	input_trimmed := strings.TrimSpace(input_split[0])
	switch input_trimmed {
	case "exit":
		fmt.Println(colorRed + "Exiting Program" + colorReset)
		os.Exit(1)
	case "users":
		users()
	case "passwds":
		change_passwd()
	default:
		fmt.Println("Command Not Found...")
	}
}

func users() {
	cmd, err := exec.Command("bash", "../scripts/getUserList.sh").Output()
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
