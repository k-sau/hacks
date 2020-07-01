### Description

Parses the amass results stored in json file.

It have following filter:
+ `cidr`: Prints only cidr
+ `ip`: Prints only ip
+ `host`: Prints only host
+ `asn`: Prints only asn

### Usage

``go run main.go -p path/to/amass/result -ip``

``go run main.go -host -p verizon/mobile.example.com.json | httprobe -p xlarge``