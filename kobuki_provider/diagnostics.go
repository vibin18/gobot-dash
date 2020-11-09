package kobuki_provider
//package main

import (
	"github.com/aler9/goroslib"
	"github.com/aler9/goroslib/msg"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

const (
	OK int8 = 0
	WARN int8 = 1
	ERROR int8 = 2
	STALE int8 = 3
	MASTER string = "192.168.178.63"
)

type KeyValue struct {
	msg.Package `ros:"main"`
	Key          string
	Value        string
}
type DiagnosticStatus struct {
	msg.Package     `ros:"main"`
	msg.Definitions `ros:"byte OK=0,byte WARN=1,byte ERROR=2,byte STALE=3"`
	Level            int8 `ros:"byte"`
	Name             string
	Message          string
	HardwareId       string
	Values           []KeyValue
}

type DiagnosticArray struct {
	msg.Package `ros:"main"`
	Header       std_msgs.Header
	Status       []DiagnosticStatus
}


var Messages = make(chan *DiagnosticArray, 1)

func OnMessage ( msg *DiagnosticArray ) {
	Messages <- msg
}


func Update () {
	// create a node with given name and linked to given master.
	// master can be reached with an ip or hostname.
	n, err := goroslib.NewNode(goroslib.NodeConf{
		Name:       "/goroslib-sub",
		MasterHost: MASTER,
	})
	if err != nil {
		panic(err)
	}
	defer n.Close()

	// create a subscriber
	sub, err := goroslib.NewSubscriber(goroslib.SubscriberConf{
		Node:     n,
		Topic:    "/diagnostics",
		Callback: OnMessage,
	})
	if err != nil {
		panic(err)
	}

	defer sub.Close()

	// freeze main loop
	select {}
}

func init (){
	go Update()
}
