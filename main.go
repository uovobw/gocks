package main

import (
	"fmt"
	"log"
	"net"

	"github.com/alecthomas/kingpin"
	"github.com/armon/go-socks5"
)

type StaticAuth struct {
	user string
	pass string
}

func (s StaticAuth) Valid(u, p string) bool {
	return (u == s.user && p == s.pass)
}

func newStaticAuth(u, p string) (s StaticAuth) {
	s = StaticAuth{
		user: u,
		pass: p,
	}
	return s
}

type StaticRuleset struct {
	destination net.IP
}

func (s StaticRuleset) AllowBind(dstIP net.IP, dstPort int, srcIP net.IP, srcPort int) bool {
	return false
}

func (s StaticRuleset) AllowAssociate(dstIP net.IP, dstPort int, srcIP net.IP, srcPort int) bool {
	return false
}

func (s StaticRuleset) AllowConnect(dstIP net.IP, dstPort int, srcIP net.IP, srcPort int) bool {
	return dstIP.Equal(s.destination)
}

func main() {

	username := kingpin.Flag("username", "SOCKS username").Short('u').Default("admin").String()
	password := kingpin.Flag("password", "SOCKS password").Short('p').Default("password").String()
	bind := kingpin.Flag("bind", "Address to bind to").Short('b').Default("127.0.0.1").IP()
	destination := kingpin.Flag("destination", "Remote address that can be reached").Short('d').Default("127.0.0.1").IP()
	port := kingpin.Flag("port", "Port to bind to").Short('P').Default("9186").String()
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	auth := newStaticAuth(*username, *password)
	ruleset := StaticRuleset{
		destination: *destination,
	}

	conf := &socks5.Config{
		Credentials: auth,
		BindIP:      *bind,
		Rules:       ruleset,
	}
	server, err := socks5.New(conf)
	if err != nil {
		log.Fatal(err)

	}

	// Create SOCKS5 proxy on localhost port 8000
	log.Printf("Running on %s:%s\n", *bind, *port)
	log.Printf("Allowed destination %s\n", (*destination).String())
	if err := server.ListenAndServe("tcp", fmt.Sprintf("%s:%s", *bind, *port)); err != nil {
		log.Fatal(err)
	}
}
