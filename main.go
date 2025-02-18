package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"

	"log"
	"strconv"
	"strings"
	"syscall"
)

func protocolToString(connType uint32, ip string) string {
	isIPv6 := strings.Contains(ip, ":")
	switch connType {
	case syscall.SOCK_STREAM: // TCP
		if isIPv6 {
			return "TCP6"
		}
		return "TCP"
	case syscall.SOCK_DGRAM: // UDP
		if isIPv6 {
			return "UDP6"
		}
		return "UDP"
	default:
		return "Unknown"
	}
}

func killProcessByName(name string) error {
	processes, err := process.Processes()
	if err != nil {
		return err
	}
	for _, p := range processes {
		n, err := p.Name()
		if err != nil {
			return err
		}
		if n == name {
			return p.Kill()
		}
	}
	return fmt.Errorf("process not found")
}

func listConnections(portFilter int, processFilter string) {
	// Get network connections
	conns, err := net.Connections("all")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%-7s %-25s %-25s %-10s %-10s\n", "Proto", "Local Address", "Foreign Address", "State", "PID/Name")

	for _, conn := range conns {
		// Filter out connections with no associated process
		if conn.Pid == 0 {
			continue
		}

		// Check if connection matches port filter
		if portFilter > 0 && int(conn.Laddr.Port) != portFilter {
			continue
		}

		// Get process name
		proc, err := process.NewProcess(conn.Pid)
		if err != nil {
			log.Println(err)
			continue
		}
		procName, err := proc.Name()
		if err != nil {
			log.Println(err)
			continue
		}

		// Check if connection matches process filter
		if processFilter != "" && !strings.Contains(strings.ToLower(procName), strings.ToLower(processFilter)) {
			continue
		}

		// Convert the protocol to a string
		protocol := protocolToString(conn.Type, conn.Laddr.IP)

		// Print connection details
		localAddr := conn.Laddr.IP + ":" + strconv.Itoa(int(conn.Laddr.Port))
		remoteAddr := conn.Raddr.IP + ":" + strconv.Itoa(int(conn.Raddr.Port))
		fmt.Printf("%-7s %-25s %-25s %-10s %-10d/%s\n", protocol, localAddr, remoteAddr, conn.Status, conn.Pid, procName)
	}
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected a subcommand (e.g: 'list')")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		portFlag := listCmd.Int("port", 0, "filter by port number")
		processFlag := listCmd.String("process", "", "filter by process name")

		listCmd.Parse(os.Args[2:])

		listConnections(*portFlag, *processFlag)

	case "kill":
		killCmd := flag.NewFlagSet("kill", flag.ExitOnError)
		processName := killCmd.String("process", "", "process to kill")

		killCmd.Parse(os.Args[2:])

		if *processName == "" {
			fmt.Println("You must provide the --process flag")
			os.Exit(1)
		}

		if err := killProcessByName(*processName); err != nil {
			fmt.Printf("Error killing process: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Successfully killed process: %s\n", *processName)

	default:
		fmt.Println("Unknown subcommand")
		os.Exit(1)
	}
}
