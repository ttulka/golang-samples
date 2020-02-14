package main

import (
  "fmt"
  "log"
  "os"
  "net"
  "github.com/google/gopacket"
  "github.com/google/gopacket/pcap"
)

var (
  snaplen = int32(1600)
  promisc = false
  timeout = pcap.BlockForever
  filter = "tcp and port 80"
  devFound = false
)

func main() {
  if len(os.Args[1:]) <= 0 {
    log.Fatalln("Usage: <IP-address-to-filter>")
  }
  ip := net.ParseIP(os.Args[1:][0])
  
  devices, err := pcap.FindAllDevs()
  if err != nil {
    log.Panicln(err)
  }
  
  var iface string
  
  for _, device := range devices {
    for _, address := range device.Addresses {
      if address.IP.Equal(ip) {
        iface = device.Name
        devFound = true
      }
    }
  }
  
  if !devFound {
    log.Panicf("Device with IP %s does not exist\n", ip)
  }
  
  handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
  if err != nil {
    log.Panicln(err)
  }
  defer handle.Close()
  
  if err := handle.SetBPFFilter(filter); err != nil {
    log.Panicln(err)
  }
  
  source := gopacket.NewPacketSource(handle, handle.LinkType())
  for packet := range source.Packets() {
    fmt.Println(packet)
  }
}