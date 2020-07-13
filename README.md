# Network Programming using Go

This contains all the learnings and some code when study Network Programming with Go, using by that title, written by Jan Newmarch https://www.apress.com/gp/book/9781484226919.

## Contents and Topics Studied

Internet Addresses: In order to use a service, you must be able to find it, the internet uses an address scheme for devices to be located.

To find the network of a device:  bitwise AND `&` it's IP address with the network mask while the device address within the subnet is found with bitwise AND of the 1's complement of the mask with the IP address.

IP4 = 8:4 (32 bits)
IP6 = 16:8 (128 bits)
Localhost: “0:0:0:0:0:0:0:1” i.e ::1


## Services (Daemons / Agents) - Background processes running on a computer.

To get a list of background services on `Mac`, it might be different for Linux but will definitely be different for Windows.
https://apple.stackexchange.com/questions/55214/whats-the-services-running-processes-manager-in-mac-os-x

Run:

        $ sudo launchctl list | less


Services run on host machines. They are typically long lived and are designed to wait for requests and respond to them. They are based on 2 methods of communication, `TCP` and `UDP`.

There is also a new Transport Protocol `SCTP - Stream Control Transport Protocol`.
P2P, RPC, Communicating agents are all built on TCP and UDP.
They are the background processes on an host machine waiting for requests or events to occur

## Ports

- Services exists on host machines


- The IP Address is needed to locate the host machine but how will we locate the services or distinguish the services on the host machines? Ports are used to distinguish each service, i.e a service on an host machine will use only one port or bind to a port and listen for connection or request from that port.


- TCP, UDP and SCTP uses ports to send packets to another process or service. Port is an unsigned integer between 1 - 65535

- Each service will allocate itself with one or more of these ports. On Unix system, the commonly used ports are listed in the file `/etc/services`. Telnet - 23, DNS - 53, SSH - 22, HTTP 80, HTTPS - 443 etc

- If the address is an IPv6 address, which already has colons in it, then the host part must be enclosed in square brackets, such as host:port => `[::1]:23`. Another special case is often used for servers, where the host address is zero, so that the TCP address is really just the port name, as in ":80" for an HTTP server.”

- On Unix, you cannot listen on a port below 1024 except you are the root but ports below 128 are standards

## TCP Client

- Once a client has established a TCP address for a service, it "dials" the service”
If successful, the service returns a connection object.


- One of the possible messages that a client can send is the "HEAD" message.

- It is normal for failures in networking programs: the opportunities for failure are substantially greater than for standalone programs. Hardware may fail on the client, the server, or on any of the routers and switches in the middle; communication may be blocked by a firewall; timeouts may occur due to network load; the server may crash while the client is talking to it.

- Server blocks on an accept call waiting for the client to send a connection and the server returns a connection object which will form the communication path between the 2 agents

Important to Note:

    The server should run forever, so that if any error occurs with a client, the server just ignores that client and carries on.
    Server should not shut down due to any client Requests, it should only fail to start up when the resources required for it to function properly is not there

### Controlling TCP Connections

- Time Out: server or client may wish to timeout each other so that if there’s a delay in response after a certain time it should end the connection with the respective hosts.

- Keep Alive or Staying Alive is used when you want to prolong the connection between the 2 parties or hosts

## UDP Datagrams - User Datagram Protocol

Connectionless protocol.
Message contains information about origin and destination
UDP client and servers makes use of datagrams to exchange information
Messages are not guaranteed to arrive or may arrive out of order

Also, note:

    A server may decide to listen to a client on more than one port, in that case the server has to implement a polling mechanism between the ports.

## Raw Sockets and the type IPConn - Go

- Raw sockets, allow the programmer to build their own IP protocols, or use protocols other than TCP or UDP


- TCP and UDP are not the only protocols built above the IP layer. The file lists about 140 of them (this list is often available on Unix systems in the file /etc/protocols).

Run:

    $ cat /etc/protocols | less


- Go allows you to build raw sockets, to enable you to communicate using one of these other protocols besides TCP or UDP. Even build your own protocol but it gives minimal support.

- Ping uses echo command from the `ICMP protocol`, it is a byte-oriented protocol. Client sends streams of bytes or continuous data to another host and host replies.

The format is as follows:

    First byte is 8: the echo message
    Second bytes is 0
    3rd and 4th bytes are a checksum on the entire message
    5th and 6th bytes are arbitrary identifier
    7th and 8th are arbitrary sequence number
    The rest of the packet is user data

## Communication

- Communication between a client and a service requires the exchange of data.

- The data has to be serialized or marshalled for transport over the wire. Messages are sent across the network as a sequence of bytes, which has no structure except for a linear stream of bytes.

- IP, TCP, UDP don’t know anything about whatever data structure is in our project which holds the state of the application. All they want is a sequence bytes. This means an application has to serialize any data into a stream of bytes in order to write it or send across the  network and deserialize it back into the data structure it was serialized from.

The unmarshalling side has to know exactly how the data is serialised in order to unmarshal it correctly.

### Encoding and Decoding (Serialization and Deserialization)

`Encoding` is the representation or the structure of data to bytes and then `Decoding` is how to use that pattern to match it to the direct representation.

Some encoding formats:

#### ASN.1

Abstract Syntax Notation One (ASN.1) was originally designed in 1984 for the telecommunications industry. ASN.1 is a complex standard, and a subset of it is supported by Go in the package `asn1`.
Its primary use in current networking systems is as the encoding for `X.509 certificates` which are heavily used in authentication systems.

Unmarshalling requires allocated memory to unmarshall data and the type of  storage must be enough to contain or represent data, if not it’ll overflow because it can not occupy the allocated memory space.

Go uses the `reflect package` to marshall/unmarshall structures, so it must be able to examine all fields - why you had to export a field from the type if you want it to be encoded / decoded.

#### JSON

Javascript Object Notation

#### Gob

Gob is a serialisation technique specific to Go. 
It is designed to encode Go data types specifically.
To use Gob to marshall a data value, you first need to create an Encoder.
The encoder has a method Encode which marshalls the value to the stream i.e bytes


#### Base64 String

A way to transmit string over the network as stream of bytes

## Application-Level Protocols

A client and a server exchange messages consisting of message types and message data. 
This requires design of a suitable message exchange protocol.
A client and server needs to exchange information through messages.
TCP, UDP, SCTP are transport mechanisms to get the message sent.

A protocol defines what type of conversation can take place between two components of a distributed application, by specifying messages, data types, encoding formats and so on.

Protocols have versions as it changes, Factors to be considered are:

- Communication: broadcast (UDP) or point-to-point (TCP, UDP)
- State: Knowing or having information about the client or the server (Stateless or Stateful)
- Reliable transport protocol or unreliable protocol
- Are response needed
- Data format
- Bursty communication stream (Ethernet or Internet) or steady stream (video and voice)

Data Format
Byte encoded
Character encoded
You can use tools like tcpdump to snoop on TCP traffic and see immediately what clients are sending to servers.

State
Applications often make use of state information to simplify what is going on
In a distributed system, such state information may be kept in the client, in the server, or in both.

## Managing Character Sets and Encodings

There are many languages in use throughout the world and they use many different character set.

Language characters get encoded into bytes using different encoding formats specific for it.

Text files and I/O consists of stream of bytes with each byte representing a single character

Internationalisation (i18n): is how you write your applications so that they can handle the variety of languages and cultures of people.

Localisation (l10n): is the process of customising your internationalised application to a particular cultural group.

### Character

A character is a unit of information that roughly corresponds to a grapheme (written symbol) of a natural language, such as a letter, numeral, or punctuation mark.

The concept of character also includes control characters, which do not correspond to natural language symbols but to other bits of information used to process texts of the language.

A character does not have any particular appearance, although we use the appearance to help recognise the character. However, even the appearance may have to be understood in a context.

Character set: A character repertoire is a set of distinct characters, such as the Latin alphabet.

### Character Code

A character code is a mapping from characters to integers. 
The mapping for a character set is also called a coded character set or code set

Code point: Is the value of each character in this mapping i.e character code
The codepoint for 'a' is 97 and for 'A' is 65 (decimal).

### Character Encoding

To communicate or store a character you need to encode it in some way.
To transmit a string, you need to encode all characters in the string.

There are many possible encodings for any code set.
`Character (A) -> Character Code (65) -> Character Encoding ( 7-bit ASCII code) -> Machine representation (8-bit bytes )  01000001.`

The character encoding is where we function at the programming level.
32-bit word-length encoding. `ASCII 'A' would be encoded as 00000000 00000000 0000000 01000001`.

### Transport Encoding

A character encoding will suffice for handling characters within a single application. 
However, once you start sending text between applications, then there is the further issue of how the bytes, shorts or words are put on the wire.

Encoding could be based on space and bandwidth-saving techniques like zipping the text.

It could be reduced to 7-bit format to allow parity checking base64.

However the convention of how data sent over is encoded is checked in the header of the request

    Content-Type: text/html; charset=ISO-8859-4: character encoding
    Content-Encoding: gzip : transfer encoding

#### Unicode

Unicode is an embracing standard character set intended to cover all major character sets in use.
The first 256 code points correspond to ISO 8859-1, with US ASCII as the first 128. There is thus a backward compatibility with these major character sets, as the code points for ISO 8859-1 and ASCII are exactly the same in Unicode.

To represent Unicode characters in a computer system, an encoding must be used.

    UTF-32 is a 4-byte encoding in machine code
    UTF-16 encodes the most common characters into 2 bytes with a further 2 bytes for the "overflow", with ASCII and ISO 8859-1 having the usual values
    UTF-8 uses between 1 and 4 bytes per character, with ASCII having the usual values (but not ISO 8859-1)
    UTF-7 is used sometimes, but is not common


Go uses UTF-8 encoded characters in its strings. Each character is of type rune. This is an alias for int32 as a Unicode character can be 1, 2 or 4 bytes in UTF-8 encoding. In terms of characters, a string is an array of runes.

Please, note:

    A string is also an array of bytes, but you have to be careful: only for the ASCII subset is a byte equal to a character. All other characters occupy two, three or four bytes.

    A string of bytes is only equal to a string of runes when they both contain ASCII characters.

    str := "百度一下，你就知道"

    println("String length", len([]rune(str))) // represented as int32 and the byte they occupy, an array of the Unicode code points which is generally the number of characters

    println("Byte length", len(str)) // represented as the number of bytes

If Go cannot properly decode bytes into Unicode characters, then it gives the Unicode Replacement Character `\uFFFD`.

Converting strings to slice or runes brings the code point representation.


    UTF-16
    str := "百度一下，你就知道"

    runes := utf16.Encode([]rune(str))
    ints := utf16.Decode(runes)


## Security

Though the internet was created with the mindset that it would be safe, the reality is that it is not - so many security issues, Spam mails, denial of service attacks, phishing attempts.

Now applications have to be built to survive hostile environment.
Ensuring privacy and integrity of data transferred, access only to legitimate users and other issues.

### Data Integrity

Means supplying a means of testing to ensure that data has not been tampered with.
Easily or usually done by forming a simple number out of the bytes in the data.
This above process is called `hashing` which produces a hash value or hash.
Ensure that the hashing algorithm chances of collision with another sequence of bytes is 0 or almost impossible

### Symmetric Key Encryption

There are two major mechanisms used for encrypting data. 
The first uses a single key that is the same for both encryption and decryption. 
This key needs to be known to both the encrypting and the decrypting agents
In general encryption algorithms become weaker over time as computers get faster.
Once you have a cipher, you can use it to encrypt and decrypt blocks of data.

### Public Key Encryption

This type of encryption requires 2 keys, one key to encrypt and another to decrypt.
The encryption key is usually made public in some way so that anyone can encrypt messages to you. 
The decryption key must stay private, otherwise everyone would be able to decrypt those messages.

Public key systems are asymmetric, with different keys for different uses
E.g RSA Scheme

### X.509 Certificates

A `Public Key Infrastructure (PKI)` is a framework for a collection of public keys, along with additional information such as owner name and location, and links between them giving some sort of approval mechanism.
Principal PKI today is `X.509 certificates`, 
Web browsers use them to verify the identity of websites.

### TLS (Transport Layer Security formerly Secure Sockets Layer - SSL)

It handles the encryption and decryption of messages automatically, avoiding the manual toil of doing it ourselves.

A client and server negotiate their identity using a PKI - X.509 certificate, when the negotiation completes, a secret key is created between them.
The secret key is then used to handle all encryption and decryption between the client and server, though the initial stage of negotiation, when it completes, a private key is generated which makes it faster.


## HTTP

World Wide Web is distributed system applications with millions if not billions of users.
A web site may become a WWW system host by running HTTP server.
Web Clients are typically users with a Web Browser, also known as `User Agents`.

There are more User Agents like web spiders, web application clients.
Web is built on top of `HyperText Transport Protocol which is layered on TCP`.

URLs specify the location of a resource.
A resource is sometimes a static file (html document, image, video file, sound file etc), it may be a dynamically generated content like some data on a database.

A user-agent is like a client or the browser, whatever makes the request to the server.
When a user makes a request for a resource, what is returned is a representation of the resource e.g if the resource is a static file, what is returned is a copy of the file and not the actual file.

Multiple URLs may point to the same resource and the HTTP server will send the appropriate representation to the user agents.
A company might make product information available both in production(external use) and internally using separate URLs pointing to same products but their response will be different.

HTTP as a protocol has to deliver requests from user-agents to servers and return a byte of stream from the server after processing the request.

### HTTP Characteristics

- HTTP does not keep state between client and server - Stateless
- It is connectionless - once a request is done, the connection is terminated
- Each requests require a separate TCP connection, so if many resources are requested, many TCP connections have to be created and teardown for the purpose of processing those requests

HTTP Versions:

    1.1 - current
    2.0
    3.0

`HEAD` is a request from user-agents which asks for information about a resource and its HTTP server.
Using the client object in Go makes it easy to make a request


Proxy Handling:

The HTTP header of the request should contain the full URL to the destination
The Host field of the header should be set to the proxy as long as the proxy is configured to pass such requests through, then that is all that needs to be done.
Go considers Proxy to be part of the Transport layer and it exists on the Transport field of the Client object or struct.

HTTPS connections by clients:

Since HTTPS is for secure and encrypted connections, it uses TLS to achieve this i.e `HTTP + TLS = HTTPS`

Servers are required to return valid X.509 certificates before a client will accept data from them.

Many sites have invalid certificates. They may have expired, they may be self-signed instead of by a recognised Certificate Authority or they may just have errors (such as having an incorrect server name). Browsers such as Firefox put a big warning notice with a "Get me out of here!" button, but you can carry on at your risk - which many people do.

### Remote Procedure Call

Socket and HTTP Programming use a message passing paradigm.
Both client and server sends messages, they are responsible for creating messages in format that is understood by both sides and also responsible for reading it.

An RPC client will package a procedure or function up in a network message, send it over the wire and the server will unpack it and call the appropriate function or procedure.

Any RPC system needs a Transport mechanism to get the data across the network, you can use HTTP or TCP.

### Web Sockets

It is meant to solve the problem of when a server wants to push or initiate a connection or send a message to the browser or user-agent without the client or user-agent(browser) initiating the connection.

Web Sockets allow a full duplex connection for this to happen.
The browser or any user agent sends a long-lived TCP connection to a web socket server.
The connection allows either side to send packets or data.

So, any application protocol can be used with web sockets.
Client or user-agent initiates the request to the server and the TCP connection is kept open.
Client makes an HTTP connection, the server replaces the HTTP protocol with `WS` protocol using the same TCP connection.

#### Web Socket Server

A web socket server starts off by being an HTTP server, 
Accepts TCP connection and handles HTTP requests on the connection.

When a request comes in that switches that connection to being a web socket connection
The protocol handler is changed from being an HTTP handler to a WebSocket handler.

So it is only that TCP connection that came with a web socket that get its role changed.
The TCP socket underlying that web socket request is used as a web socket.

The server continues to be an HTTP server for other requests.
To handle websocket, we simply register a different kind of handler (web socket handler) based on a URL pattern, common or conventional one is `/ws`.

A complex web server can handle HTTP and Web Socket requests by having more handlers.

### The Message object

HTTP is a stream protocol.

WebSockets is a frame-based protocol, you prepare a block of data of any size and send it as a set of frames.

Frames can be UTF-8 strings or a sequence of bytes.

If you are reading data from a stream - file, connection, from wire or encoding, there should an allocated memory space or buffer to receive the data or temporarily store it.
