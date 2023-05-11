# multipass

## Name

*multipass* - dynamically adds machines created with [`multipass`](https://multipass.run/) to a zone.

## Description

The *multipass* plugin adds `A` records to a zone for all the VMs managed by multipass.

The plugin will match the first part of the request to the VM name and use the first IPv4 address of the VM as `A` record in the response (it could be possible to answer all the IPs, but I do not need it).

The plugin calls the `multipass` binary with the `list` command and the `--format=json` flags. So you need a working `multipass` client in the same machine that CoreDNS is running.

### But Why?

multipass uses macos' virtualization framework to create the virtual machines. The virtualization framework uses an isolated network, so you can't for example use your network's DHCP+DNS servers to access the machines. mDNS is not great for general purposes.

Using CoreDNS as resolver for your machine, allows you to have more flexibility and the multipass plugins gives you automatic DNS resolution for all the machines managed by multipass.

## Examples

Consider the following `Corefile`:

```Corefile
example.org:1053 {
    multipass
}
```

and that the following VM has been created with `multipass`:

```console
$ multipass list
Name                    State             IPv4             Image
primary                 Running           192.168.64.127   Ubuntu 22.04 LTS
```

querying CoreDNS for `primary.example.org` will return `192.168.64.127`:

```console
$ dig @127.0.0.1 -p 1053 primary.example.org

; <<>> DiG 9.10.6 <<>> @127.0.0.1 -p 1053 primary.example.org
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 35613
;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1
;; WARNING: recursion requested but not available

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
;; QUESTION SECTION:
;primary.example.org.		IN	A

;; ANSWER SECTION:
primary.example.org.	3600	IN	A	192.168.64.127

;; Query time: 560 msec
;; SERVER: 127.0.0.1#1053(127.0.0.1)
;; WHEN: Thu May 11 16:27:51 CEST 2023
;; MSG SIZE  rcvd: 83
```
