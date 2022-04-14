package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"sync"
)

const ChunkDirMapping = "/Users/terrill/chunk_dir2.dat"

//const ChunkDirMapping = "/ardir/chunk_dir.dat"

var bootMap = func() *BootMap {
	return &BootMap{
		lock:      &sync.Mutex{},
		Connected: make(map[string]*Peer),
	}
}()

type Peer struct {
	lock      *sync.Mutex
	Connected map[string]*Peer
}

type BootMap struct {
	lock      *sync.Mutex
	Connected map[string]*Peer
}

type Chunk struct {
	Name string
	Path string
}

// start copier
func main() {
	Mod := os.Args[1]
	switch Mod {
	case "bootstrap":
		fmt.Println("started as bootstrap node")
		// peer as bootstrap
		startBootstrap()
	case "node":
		Local, err := LocalIPv4s()
		if err != nil {
			panic("can't get local net ipv4 ip")
		}
		fmt.Println("started as a peer ", Local)
		// peer as node
		startServing()
	default:
		panic("start up wrong mode")
	}
}

// accept all connections up limit 1000,
// then connected nodes with alive peers
func startBootstrap() {

}

// start connection to bootstrap node asking for avail nodes
func startServing() {
	// ask bootstrap
	peers, err := askForAvailNodesIP()
	if err != nil {
		panic("can't connect boot node")
	}
	if peers == nil {
		fmt.Println("genesis peer")
		fmt.Println("full state")
	}
	// if one peer is full state and
}

// LocalIPs return all non-loopback IPv4 addresses
func LocalIPv4s() (string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic("wrong network")
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	if ips == nil {
		panic("wrong network")
	}
	return ips[0], nil
}

func askForAvailNodesIP() (peers []string, err error) {
	return nil, nil
}

func chunks() (res []Chunk) {
	err := ReadLine(ChunkDirMapping, func(s string) {
		s = strings.Trim(s, ".")
		s = strings.Trim(s, "{")
		s = strings.Trim(s, "}")
		s = strings.Trim(s, "\"")
		mid, err := GetAllFile(s)
		if err != nil {
			panic("can't load files correctly")
		}
		res = append(res, mid...)
	})
	if err != nil {
		panic("can't load files correctly")
	}
	return
}

func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			return nil
		}
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func GetAllFile(pathname string) (s []Chunk, err error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}

	for _, fi := range rd {
		c := Chunk{}
		if !fi.IsDir() {
			fullName := pathname + "/" + fi.Name()
			c.Name = fi.Name()
			c.Path = fullName
			s = append(s, c)
		}
	}
	return s, nil
}
