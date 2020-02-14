package main

import (
  "fmt"
  "log"
  "bytes"
  "github.com/google/gopacket"
  "github.com/google/gopacket/pcap"
)

var (
  snaplen = int32(1600)
  promisc = false
  timeout = pcap.BlockForever
  filter = "tcp and dst port 21"
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
    appLayer := packet.ApplicationLayer()
    if appLayer == nil {
      continue
    }
    payload := appLayer.Payload()
    if bytes.Contains(payload, []byte("USER")) {
      fmt.Print(string(payload))
    } else if bytes.Contains(payload, []byte("PASS")) {
      fmt.Print(string(payload))
    }
  }
}