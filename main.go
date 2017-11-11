package main
import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"github.com/olekukonko/tablewriter"
	"strings"
	"strconv"
	"os/exec"
	"os/user"
)

type ShuttleConfig struct {
	Terminal string `json:"terminal"`
	ITermVersion string `json:"iTerm_version"`
	DefaultTheme string `json:"default_theme"`
	LaunchAtLogin bool `json:"launch_at_login"`
	ShowSSHConfigHosts bool `json:"show_ssh_config_hosts"`
	SSHConfigIgnoreHosts []interface{} `json:"ssh_config_ignore_hosts"`
	SSHConfigIgnoreKeywords []interface{} `json:"ssh_config_ignore_keywords"`
	Hosts []interface{} `json:"hosts"`
}

type Host struct {
	Name string
	Cmd string
}

var Hosts []Host

func main() {
	if (len(os.Args) <= 1) {
		printUsage()
		os.Exit(1)
	}

	usr, _ := user.Current()
	raw, err := ioutil.ReadFile(usr.HomeDir + "/.shuttle.json")
	if err != nil {
		fmt.Println("Dude there is something wrong with ~/.shuttle.json")
		os.Exit(1)
	}

	var c ShuttleConfig
	json.Unmarshal(raw, &c)

	ParseHosts(c.Hosts)

	switch os.Args[1] {
		// Print all shuttle hosts config
		case "ls": {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"#", "name", "command"})

			for index, v := range Hosts {
				table.Append([]string{fmt.Sprint(index), v.Name, v.Cmd})
			}
			table.Render() // Send output
		}
		// Edit shuttle config
		case "e": {
			execCommand("vi", usr.HomeDir + "/.shuttle.json")
		}
		default: {
			// Connect to the selected host
			if len(os.Args) == 2 {
				connect()
			} else {
				printUsage()
				os.Exit(1)
			}
		}
	}
}

func connect() {
	var i, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("The first argument should be a number but found ", os.Args[1])
		os.Exit(1)
	}
	if len(Hosts) == 0 {
		fmt.Println("No hosts to connect")
		os.Exit(1)
	}
	if i < 0 || i >= len(Hosts) {
		fmt.Println(os.Args[1] + " is invalid or not in the list")
		os.Exit(1)
	}
	Cmd := Hosts[i].Cmd
	CmdFields := strings.Fields(Cmd)
	ProgramName := CmdFields[0]
	ProgramArgs := CmdFields[1:]
	execCommand(ProgramName, ProgramArgs...)
}

func execCommand(program string, args ...string) {
	cmd := exec.Command(program, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func printUsage() {
	fmt.Println("Usage: ")
	fmt.Println("	shuttle <number> (Connect to host)")
	fmt.Println("	shuttle ls (Show hosts)")
	fmt.Println("	shuttle e (Edit shuttle configuration)")
}

func AppendHost(HostMap map[string]interface{}) {
	Name, _ := HostMap["name"].(string)
	Cmd, _ := HostMap["cmd"].(string)
	Hosts = append(Hosts, Host{Name, Cmd})
}

func ParseHost(Host map[string]interface{}) {
	L:
	for _, groupVal := range Host {
		switch groupValConcrete := groupVal.(type) {
		case []interface{}:
			ParseHosts(groupValConcrete)
		case map[string]interface{}:
			AppendHost(groupValConcrete)
		case string:
			AppendHost(Host)
			break L
		}
	}
}

func ParseHosts(Hosts []interface{}) {
	for _, val := range Hosts {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			ParseHost(concreteVal)
		}
	}
}