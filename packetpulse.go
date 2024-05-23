package main

import (
	"fmt"
	"log"

	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	// red := "\033[31m"
	green := "\033[32m"
	reset := "\033[0m"
	// Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	// Choose a device to listen on
	// For example, choose the first device
	device := devices[0].Name

	// Open the device for packet capture
	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set filter for capturing only TCP packets
	err = handle.SetBPFFilter("tcp")
	if err != nil {
		log.Fatal(err)
	}

	// Start processing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Printf("%sPacket Raw data:%s\n", green, reset)
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Println(packet.Data())
		time.Sleep(3 * time.Second)
		fmt.Printf("%sNetwork Layer data extracted:%s\n", green, reset)
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Println(packet.NetworkLayer())
		fmt.Printf("%sApplication Layer data extracted:%s\n", green, reset)
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Println(packet.ApplicationLayer())
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Printf("%sDumped data extracted:%s\n", green, reset)
		fmt.Println(packet.Dump())
	}
}
