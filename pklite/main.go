package main

import (
	"context"
	"flag"
	"fmt"
	"pkoderlite"
)

var (
	buildTime    string
	compiler     string
	majorVersion string
	minorVersion string
	patchVersion string
)

func main() {
	var (
		protocol string
		rtmp     string
		mpegts   string
		version  bool
	)
	flag.StringVar(&rtmp, "rtmp", "", "rtmp://<ip>:<port>/appid")
	flag.StringVar(&mpegts, "mpegts", "", "udp://<ip>:<port>")
	flag.StringVar(&protocol, "protocol", "rtmp", "standalone protocol to use: rtp | tcp | rtmp (default)")
	flag.BoolVar(&version, "version", false, "print version")
	flag.Parse()

	smallVersion := fmt.Sprintf("%s.%s.%s", majorVersion, minorVersion, patchVersion)
	largeVersion := fmt.Sprintf("PKST Encoder (lite) version %s.%s.%s - Copyright @ Studio PKST LLC 2023\nCompiler: %s - Build Time: %s",
		majorVersion, minorVersion, patchVersion, compiler, buildTime)

	if version {
		fmt.Printf("%s\n", smallVersion)
		return
	}

	fmt.Printf("%s\n-\n", largeVersion)

	port := 3000

	ip, err := pkoderlite.PrivateIPv4()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	listen := pkoderlite.ListenResource(protocol, ip.String(), port)

	fmt.Println("Listen on: ", listen)
	fmt.Println("Destination: ", rtmp, mpegts)
	worker := pkoderlite.BaseWorker{
		StreamID:  "std",
		ListenURI: listen,
		Protocol:  protocol,
		Timeout:   0,
		RtmpDst:   []string{rtmp},
		MpegTs:    []string{mpegts},
		Ctx:       context.TODO(),
	}
	fmt.Println(worker.Do())
}
