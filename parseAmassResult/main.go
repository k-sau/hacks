package main

import (
	"encoding/json"
	"os"
	"flag"
	"fmt"
	"bufio"
)

type Obj struct {
	Name	string	`json:"name"`
	Domain	string	`json:"domain"`
	Tag	string	`json:"tag"`
	Source	string	`json:"source"`
    Addresses []Address `json:"addresses"`
}

type Address struct {
	Ip	string `json:"ip"`
	Cidr	string	`json:"cidr"`
	Asn	int	`json:"asn"`
	Desc string	`json:"desc"`
}


func main() {
	path := flag.String("p", "", "path of the amass result")
	cidr := flag.Bool("cidr", false, "prints only cidr")
	host := flag.Bool("host", false, "prints only hosts")
	ip := flag.Bool("ip", false, "prints only ip")
	asn := flag.Bool("asn", false, "prints only asn")
	help := flag.Bool("h", false, "prints available flags")
	flag.Parse()

	if *path == "" {
		fmt.Println("argument `p` is required")
		flag.Usage()
		os.Exit(1)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	result, err := os.Open(*path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer result.Close()

	scanner := bufio.NewScanner(result)
	for scanner.Scan() {
		var obj Obj
		json.Unmarshal([]byte(scanner.Text()), &obj)
		if *cidr {
			for _, v := range obj.Addresses {
				fmt.Println(v.Cidr)
			}
		} else if *ip {
			for _, v := range obj.Addresses {
				fmt.Println(v.Ip)
			}
		} else if *asn {
			for _, v := range obj.Addresses {
				fmt.Println(v.Asn)
			}
		} else if *host {
			fmt.Println(obj.Name)
		} else {
			flag.Usage()
			os.Exit(1)
		}
    }	
}