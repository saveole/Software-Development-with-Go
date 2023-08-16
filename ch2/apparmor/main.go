package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	appArmorEnabledPath = "/sys/module/apparmor/parameters/enabled"
	appArmorModePath = "/sys/module/apparmor/parameters/mode"
)

func main() {
	fmt.Println("AppArmor mode : ", appArmorMode())
	fmt.Println("AppArmor is enabled : ", appArmorEnbaled())
}

func appArmorMode() (mode string) {
	content, err := ioutil.ReadFile(appArmorModePath)
	if err != nil {
		log.Fatal("error (read mode) : ", err)
	}
	return strings.TrimSpace(string(content))
}

func appArmorEnbaled() (support bool) {
	content, err := ioutil.ReadFile(appArmorEnabledPath)
	if err != nil {
		log.Fatal("error (read mode) : ", err)
	}
	return strings.TrimSpace(string(content)) == "Y"
}