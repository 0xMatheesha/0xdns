// SPDX-License-Identifier: MIT
// Copyright (c) 2025 0xMatheesha

package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"codeberg.org/miekg/dns"
)

var dnsPort string = ":8053"
var protocol string = "udp"
var upstream string = "8.8.8.8:53"

var blacklist map[string]struct{}

const (
	Green  = "\033[32m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

func main() {
	fmt.Println(Yellow, "Getting ready 📝", Reset)
	getBlockedHosts()
	//initialize the server 📃
	dns.HandleFunc(".", handleDnsReq)
	server := &dns.Server{
		Addr: dnsPort,
		Net:  protocol,
	}
	fmt.Println(Yellow, "Starting the server 🛫", Reset)
	//start the dns server
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Failed to connect! 🔌", err)
	}

	// fmt.Println(Green, "Server started ✈️", Reset)
	// select {}
}

func handleDnsReq(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) {
	if len(r.Question) == 0 {
		m := new(dns.Msg)
		m.Rcode = dns.RcodeFormatError
		m.WriteTo(w)
		return
	}

	client := new(dns.Client)
	q := r.Question[0]
	domain := strings.TrimSuffix(q.Header().Name, ".")

	msg := new(dns.Msg)
	msg.ID = r.ID
	msg.Response = true
	msg.Question = r.Question

	r.RecursionDesired = true

	if isBlocked(domain) {
		msg.Rcode = dns.RcodeNameError
	} else {
		res, _, err := client.Exchange(ctx, r, "udp", upstream)
		if res == nil || err != nil || len(res.Question) == 0 || len(res.Answer) == 0 {
			m := new(dns.Msg)
			m.ID = r.ID
			m.Rcode = dns.RcodeServerFailure
			m.WriteTo(w)
			fmt.Println(Red, err, Reset)
			return
		}

		for _, ans := range res.Answer {
			if a, ok := ans.(*dns.A); ok {
				msg.Answer = append(msg.Answer, a)
			}
		}

	}
	msg.WriteTo(w)
}

func getBlockedHosts() {
	blacklist = make(map[string]struct{})

	f, err := os.Open("blacklist.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		//check if the line is a comment or is empty
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		for _, host := range fields[1:] {
			h := strings.TrimSuffix(strings.ToLower(strings.TrimSpace(host)), ".")
			blacklist[h] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func isBlocked(d string) bool {
	if _, blocked := blacklist[d]; blocked {
		fmt.Println(Red, "Blocked:", Reset, d)
		return true
	}
	return false
}
