package main

import (
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/pamelasanchezvi/protobuf-message-examples/messages"
	"net"
	"os"
	"strconv"
)

func createResourceOfferMessage() ([]byte, error) {

	f := float64(2016)

	scalar := &messages.Value_Scalar{
		Value: &f,
	}
	/*
		t := "testtext"

			text := &messages.Value_Text{
				Value: &t,
			}

			           val := &Value{
			   		Type: int32(2),
			   		Scalar: scalar,
			   		Text: text,
			   	}
	*/
	offerId := new(messages.OfferID)
	testoffer := "pamtestoffer"
	offerId.Value = &testoffer

	fwId := new(messages.FrameworkID)
	testfw := "pamframeworkid"
	fwId.Value = &testfw

	testslave := "pamslave"
	slave := new(messages.SlaveID)
	slave.Value = &testslave

	testhost := "pamubuntuhost"
	testrole := "testpamrole"
	memres := "mem"

	testres := new(messages.Resource)
	testres.Name = &memres
	testType := messages.Value_Type(3)
	testres.Type = &testType
	testres.Scalar = scalar
	testres.Role = &testrole
	resourceArray := []*messages.Resource{testres}

	OfferMsg := new(messages.Offer)
	OfferMsg.Id = offerId
	OfferMsg.FrameworkId = fwId
	OfferMsg.SlaveId = slave
	OfferMsg.Hostname = &testhost
	OfferMsg.Resources = resourceArray

	OfferArray := []*messages.Offer{OfferMsg}

	ProtoMessage := new(messages.ResourceOffersMessage)
	ProtoMessage.Offers = OfferArray
	pid := proto.String("5.5.5.5")
	ProtoMessage.Pids = []string{*pid}

	//fmt.Println(ProtoMessage.Messageitems)
	return proto.Marshal(ProtoMessage)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func sendDataToDest(data []byte, dst *string) {
	conn, err := net.Dial("tcp", *dst)
	checkError(err)
	n, err := conn.Write(data)
	checkError(err)
	fmt.Println("Sent " + strconv.Itoa(n) + " bytes")
}

func main() {
	dest := flag.String("d", "127.0.0.1:8080", "Enter the destnation socket address")
	flag.Parse()
	data, err := createResourceOfferMessage()
	checkError(err)
	sendDataToDest(data, dest)
}
