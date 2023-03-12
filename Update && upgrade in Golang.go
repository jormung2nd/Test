package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Updating Kali Linux...")
	cmd := exec.Command("sudo", "apt", "update")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error updating package lists: %v", err)
	}

	cmd = exec.Command("sudo", "apt", "full-upgrade", "-y")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error upgrading packages: %v", err)
	}

	cmd = exec.Command("sudo", "apt", "autoremove", "-y")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error removing unused packages: %v", err)
	}

	fmt.Println("Kali Linux update completed.")
}
