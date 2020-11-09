package test

//package main

import (
	"github.com/aler9/goroslib"
	"github.com/aler9/goroslib/msg"
	"github.com/aler9/goroslib/msgs/geometry_msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

const (
	OK     int8   = 0
	WARN   int8   = 1
	ERROR  int8   = 2
	STALE  int8   = 3
	MASTER string = "192.168.178.63"
)

type Quaternion struct {
	msg.Package `ros:"main"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	Z           float64 `json:"z"`
	W           float64 `json:"w"`
}

type Point struct {
	msg.Package `ros:"main"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	Z           float64 `json:"z"`
}

type Vector3 struct {
	msg.Package `ros:"main"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	Z           float64 `json:"z"`
}

type Pose struct {
	msg.Package `ros:"main"`
	Position    Point
	Orientation Quaternion
}

type Twist struct {
	msg.Package `ros:"main"`
	Linear      Vector3
	Angular     Vector3
}

type TwistWithCovariance struct {
	msg.Package `ros:"main"`
	Twist       Twist
	Covariance  [36]float64
}

type PoseWithCovariance struct {
	msg.Package `ros:"main"`
	Pose        Pose
	Covariance  [36]float64
}

type Odometry struct {
	msg.Package  `ros:"main"`
	Header       std_msgs.Header
	ChildFrameId string `json:"ChildFrameId"`
	Pose         geometry_msgs.PoseWithCovariance
	Twist        geometry_msgs.TwistWithCovariance
}

var OdomMessages *Odometry

//var msgOdom *Odometry
//
//var OdomMessages *Odometry

//func OnMessageOdom(msgOdom *Odometry) {
//	e, err := json.Marshal(msgOdom)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	//fmt.Println(string(e))
//	OdomMessages = fmt.Sprintf(string(e))
//	fmt.Println("Message received")
//	fmt.Println(OdomMessages)
//}

func OdomUpdate() {
	//var msgOdom *Odometry

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

	Getty := func(msgOdom *Odometry) {
		//OdomMessages <- msgOdom
		OdomMessages = msgOdom

	}
	// create a subscriber
	sub, err := goroslib.NewSubscriber(goroslib.SubscriberConf{
		Node:     n,
		Topic:    "/odom",
		Callback: Getty,
	})
	if err != nil {
		panic(err)
	}

	defer sub.Close()

	// freeze main loop
	select {}
}
