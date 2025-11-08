# 0xDNS (HexDNS) — Simple DNS-Level Adblocker

A lightweight DNS server in Go that blocks ads at the DNS level. Works with **IPv4 (A) records**, supports custom blocklists, and forwards non-blocked queries to an upstream DNS server.

> Made using [miekg/dns](https://codeberg.org/miekg/dns)


---

## Features

- DNS-level adblocking for faster browsing and ad-free apps  
- Customizable blocklist from a `hosts.txt`-style file  
- Supports subdomain matching for ad hosts  
- Forwards non-blocked queries to a public DNS (default: Google 8.8.8.8)  
- Lightweight, written in pure Go  

---

## Installation

```bash
git clone https://github.com/0xMatheesha/0xdns.git
cd 0xdns
```
## Building Project

```bash
go build -o 0xdns server.go
./0xdns
```
Example output:
```bash
Getting ready 📝 \
Starting the server 🛫
``` 
## Testing

```bash
dig @127.0.0.1 -p 8053 doubleclick.net
dig @127.0.0.1 -p 8083 ads.google.com
```
Example output:
```bash
Blocked:  doubleclick.net
Blocked:  ads.google.com
```
## Blacklist Files
This project uses blocklists from the following sources:
- [Steven Black's Hosts](https://github.com/StevenBlack/hosts)
- [GoodbyeAds](https://github.com/jerryn70/GoodbyeAds)
## License 🪪
MIT - © 2025 [0xMatheesha](https://github.com/0xMatheesha)