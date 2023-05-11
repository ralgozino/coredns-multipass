# multipass

## Name

*multipass* - dynamically adds machines created with [`multipass`](https://multipass.run/) to a zone.

## Description

The *multipass* plugin adds `A` records to a zone for all the VMs managed by multipass.

The plugin will match the first part of the request to the VM name and use the first IPv4 address of the VM as `A` record in the response (it could be possible to answer all the IPs, but I do not need it).

### But Why?

multipass uses macos' virtualization framework to create the virtual machines. The virtualization framework uses an isolated network, so you can't for example use your network's DHCP+DNS servers to access the machines. mDNS is not great for general purposes.

Using CoreDNS as resolver for your machine, allows you to have more flexibility and the multipass plugins gives you automatic DNS resolution for all the machines managed by multipass.

## Examples

Consider the following `Corefile`:

```Corefile
example.org:1053 {
    bind en0
    multipass
}
```

and that the following VM has been created with `multipass`:

```console
$ multipass list
Name                    State             IPv4             Image
primary                 Running           192.168.64.127   Ubuntu 22.04 LTS
```

querying CoreDNS for `primary.example.org` will return `192.168.64.127`.
