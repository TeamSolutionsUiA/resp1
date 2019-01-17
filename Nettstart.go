package main

import "os/exec"

func main() {
	chrome := "C:/Program Files (x86)/Google/Chrome/Application/chrome.exe"
	cmd := exec.Command(chrome, "--chrome-frame", "--kiosk", "http://www.vg.no", "http://www.e24.no","http://www.nordnet.no")
	err := cmd.Start()
	if err != nil {
		println("Failed to start chrome:", err)
	}
}
