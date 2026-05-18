package core

import (
	"fmt"
	"net"
	"path/filepath"

	"github.com/Turtle-In-Space/theia/helpers"
	"github.com/Turtle-In-Space/theia/models"
	"github.com/Turtle-In-Space/theia/output"
)

func CreateMsfFile(host models.Host, nmap_data string) {
	fileName := filepath.Join(host.Dir, fmt.Sprint("../start.rc")) // bad fix?
	file := helpers.CreateFile(fileName)
	defer file.Close()

	helpers.WriteLine(file, "# Created by theia")
	helpers.WriteLine(file, fmt.Sprintf("workspace -a %s", host.Name))
	helpers.WriteLine(file, fmt.Sprintf("db_import %s", nmap_data))
	helpers.WriteLine(file, fmt.Sprintf("setg lhost %s", getTun0IP())) // hardcode tun0?

	output.Debug("created msf file for: %s", host.Name)
}

func getTun0IP() string {
	iface, err := net.InterfaceByName("tun0")
	if err != nil {
		return "" // or a fallback IP
	}
	addrs, err := iface.Addrs()
	if err != nil || len(addrs) == 0 {
		return ""
	}
	// Assume the first address is IPv4
	ip := addrs[0].(*net.IPNet).IP.String()
	return ip
}
