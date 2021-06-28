# secure-channel

<p>
  <img src="./sc.png" />
</p>

Secure channel provides a secure means of communicating where eavesdropping is not possible, and it happens through secure RSA Encryption 

## Installation

```
$ git clone github.com/sap200/secure-channel

$ cd secure-channel

$ go install 
```

or 

you can directly install the binary

For linux

```
curl https://github.com/sap200/secure-channel/releases/download/v0.1-beta/secure-channel

sudo mv secure-channel /usr/local/bin
```

For Windows

```
curl https://github.com/sap200/secure-channel/releases/download/v0.1-beta/secure-channel.exe
```

## Usage

```
$ ./secure-channel -h

Usage of ./secure-channel:

  -command string

    	command: either server or client

  -ip string

    	ip: if command is server then launch ip, if command is client then server's ip  to connect

  -port uint

    	port: if command is server then launch port, if command is client then server's port to connect
```

For Server

```
./secure-channel -command server -ip <ipv4> -port <PORT>
```

For Client 

```
./secure-channel -command client -ip <ipv4> -port <PORT>
```

