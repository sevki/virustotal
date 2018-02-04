package main

import "me/vighnesh/api/virustotal"
import "fmt"

func main() {
	virustotalapi, _ := virustotal.Configure("APIKEY")
	response, e := virustotalapi.DomainReport("vighnesh.me")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("Resolutions", response.Resolutions)
	fmt.Println("Detected", response.DetectedUrls)
}
