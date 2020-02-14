package main

import (
  "fmt"
  "log"
  "github.com/google/gopacket"
  "github.com/google/gopacket/pcap"
)

var (
  snaplen = int32(1600)
  promisc = false
  timeout = pcap.BlockForever
  filter = "tcp and port 80"
)

func main() {
  devices, err := pcap.FindAllDevs()
  if err != nil {
    log.Panicln(err)
  }  
  for _, device := range devices {
    go handle(device.Name)
  }
  for {}
}

func handle(iface string) {
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
    fmt.Println("DEVICE", iface, ":::", packet)
  }
}