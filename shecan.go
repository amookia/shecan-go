package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"shecan-cli/dns"
	"time"
)

func main(){
	var snumber int
	interfaces,err := net.Interfaces()
	interfaces_slice := []string{}
	if err != nil {
		fmt.Println(err)
	}
	for i,o := range interfaces {
		fmt.Printf("%d -> %s\n",i+1,o.Name)
		interfaces_slice = append(interfaces_slice, o.Name)
	}
	fmt.Print("Select interface : ");fmt.Scanln(&snumber)
	if snumber > len(interfaces_slice) {
		log.Fatal("Not in range!")
	}
	//
	for _,i := range dns.GetDns() {
		network := `"` + interfaces_slice[snumber-1] +  `"`
		cmd := exec.Command("netsh","interface","ip", "add","dns",network,i,"INDEX=2")
		out,err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))
	}
	fmt.Println("Done!")
	time.Sleep(60)
}