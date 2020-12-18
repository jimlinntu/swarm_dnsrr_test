package main

import (
    "fmt"
    "net"
    "net/http"
    "regexp"
)

func interfaces(w http.ResponseWriter, req *http.Request){
    // https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
    ifaces, err := net.Interfaces()
    if err != nil{
        fmt.Println("panic")
        panic(err)
    }
    var ips []net.IP
    for _, i := range ifaces {
        addrs, err := i.Addrs()
        if err != nil {
            fmt.Println("panic")
            panic(err)
        }
        for _, addr := range addrs {
            var ip net.IP
            switch v := addr.(type) {
            case *net.IPNet:
                ip = v.IP
            case *net.IPAddr:
                ip = v.IP
            }
            ips = append(ips, ip)
        }
    }
    re := regexp.MustCompile(`[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`)
    for _, ip := range ips{
        str_ip:= ip.String()
        if re.Match([]byte(str_ip)) {
            _, err = fmt.Fprintf(w, "IP: %s\n", str_ip)
            if err != nil {
                panic(err)
            }
        }
    }
}

func main(){
    http.HandleFunc("/interfaces", interfaces)
    err := http.ListenAndServe("0.0.0.0:8090", nil)
    if err != nil{
        panic(err)
    }
}
