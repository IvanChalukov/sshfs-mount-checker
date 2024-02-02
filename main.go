package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func isMounted(mountPoint string) bool {
	out, err := exec.Command("mount").Output()
	if err != nil {
		log.Fatal("Error checking mount:", err)
	}
	return strings.Contains(string(out), mountPoint)
}

func mountSSHFS(username string, password string, url string, localMountPoint string, remoteMountPoint string) bool {
	var cmd *exec.Cmd
	if password != "" {
		command := fmt.Sprintf("echo %s | sshfs -p23 -o reconnect,umask=000,password_stdin,allow_other %s@%s:%s %s", password, username, url, remoteMountPoint, localMountPoint)
		cmd = exec.Command("sh", "-c", command)
	} else {
		cmd = exec.Command("sshfs -p23 -o reconnect,umask=000,allow_other ", fmt.Sprintf("%s@%s:%s %s", username, url, remoteMountPoint, localMountPoint))
	}

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error mounting SSHFS:", err)
		return false
	} else {
		fmt.Println("SSHFS mounted successfully.")
		return true
	}
}

func sendUptimeKumaUpdate(uptimeKumaURL string) {
	resp, err := http.Get(uptimeKumaURL)
	if err != nil {
		fmt.Println("Failed to send update to Uptime Kuma:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Update sent to Uptime Kuma successfully.")
	} else {
		fmt.Println("Uptime Kuma responded with status code:", resp.StatusCode)
	}
}

func main() {
	username := flag.String("username", "", "Username for SSHFS")
	url := flag.String("url", "", "URL for SSHFS")
	remoteMountPoint := flag.String("remotemountpoint", "", "Mount point on Storage machine for SSHFS")
	localMountPoint := flag.String("localmountpoint", "", "Mount point on Current machine for SSHFS")
	healthUrl := flag.String("healthurl", "", "Url to ping after successful operation")
	flag.Parse()

	password := os.Getenv("SSHFS_PASSWORD")
	if *username == "" || *url == "" || *remoteMountPoint == "" || *localMountPoint == "" || *healthUrl == "" {
		log.Fatal("username, URL, remoteMountPoint, localMountPoint and healthUrl  must all be provided.")
	}

	if !isMounted(*localMountPoint) {
		fmt.Println("SSHFS is not mounted. Attempting to mount...")
		if mountSSHFS(*username, password, *url, *localMountPoint, *remoteMountPoint) {
			sendUptimeKumaUpdate(*healthUrl)
		}
	} else {
		fmt.Println("SSHFS is already mounted.")
		sendUptimeKumaUpdate(*healthUrl)
	}
}
