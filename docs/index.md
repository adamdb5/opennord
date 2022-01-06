
# Contents
- [Introduction](#introduction)
- [Getting started](#getting-started)
- [Functions](#functions)
- [Types](#structs)
- [Interfaces](#interfaces)
- [Example](#example)

# Introduction
OpenNord was originally created because no Linux GUI existed for NordVPN. Instead, the only way of interacting with the 
NordVPN daemon was via the official closed-source commandline utility. 

There are a few projects in the AUR, however these executed the commandline program meaning the programmer 
has little control over error handling, and is restricted in functionality to what the NordVPN client exposes
to the user.

As a result, I decided to reverse engineer the NordVPN client, uncovering the underlying gRPC API to create an
open-source library for users to build custom applications. If you are interested in this process, you can read my
blog post (coming soon).

The reason this library is written in Go is for three reasons:
- The official NordVPN client is written in Go
- There is no gRPC implementation for C
- I don't really like any other languages

If you wish to implement OpenNord in another language, then please do so. To do this, you will need to use the 
`daemon.proto` file along with the `protoc` command to generate a stub for your chosen language. More information is 
available [here](https://grpc.io/docs/protoc-installation/) 

# Getting Started
To start using OpenNord, simply add the following to your `go.mod` file:
```
require (
    github.com/adamdb5/opennord v0.0.1
)
```

Now you can start using OpenNord by importing it in your `.go` files:
```go
import (
	"github.com/adamdb5/opennord"
    "github.com/adamdb5/opennord/pb"
)
```

# Functions
To start communicating with the daemon, you will first need to create a `Client` like so:
```go
client, err := opennord.NewOpenNordClient()
```

Once you have created your client, you will be able to invoke the following functions:
- [AccountInfo](#accountinfo)
- [Cities](#cities)
- [Connect](#connect)
- [Countries](#countries)
- [Disconnect](#disconnect)
- [FrontendCountries](#frontendcountries)
- [Groups](#groups)
- [IsLoggedIn](#isloggedin)
- [Login](#login)
- [LoginOAuth2](#loginoauth2)
- [Logout](#logout)
- [Plans](#plans)
- [Ping](#ping)
- [RateConnection](#rateconnection)
- [SetAutoConnect](#setautoconnect)
- [SetCyberSec](#setcybersec)
- [SetDefaults](#setdefaults)
- [SetDns](#setdns)
- [SetFirewall](#setfirewall)
- [SetIpv6](#setipv6)
- [SetKillSwitch](#setkillswitch)
- [SetNotify](#setnotify)
- [SetObfuscate](#setobfuscate)
- [SetProtocol](#setprotocol)
- [SetTechnology](#settechnology)
- [SetWhitelist](#setwhitelist)
- [Settings](#settings)
- [SettingsProtocols](#settingsprotocols)
- [SettingsTechnologies](#settingstechnologies)
- [Status](#status)

## AccountInfo
```go
func (c Client) AccountInfo() (*pb.AccountResponse, error)
```
### Description
Returns the current user's account information including email address and subscription expiry.

### Parameters
`AccountInfo` does not take any parameters.

### Returns
- [`AccountResponse`](#accountresponse) - Information about the current user's account
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## Cities
```go
func (c Client) Cities(country string) (*pb.CitiesResponse, error)
```
### Description
Returns a list of cities (servers) located in the given country.

### Parameters
- country: string - The name of the country. If the country has a space in its name, this should be replaced with an 
underscore. e.g. United Kingdom becomes "united_kingdom".

### Returns
- [`CitiesResponse`](#citiesresponse) - Cities matching the given criteria.
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## Connect
```go
func (c Client) Connect(req *pb.ConnectRequest) (pb.Daemon_ConnectClient, error)
```
### Description
Connects to a NordVPN server.

### Parameters
- req: [`ConnectRequest`](#connectrequest) - Contains the parameters for the connection.

### Returns
- [`Daemon_ClientConnect`](#daemon_connectclient) - A stream for monitoring the connection progress.
After invoking `Connect`, the NordVPN daemon will send two messages. These can be retrieved by calling `Recv`, 
which will return a [`ConnectResponse`](#connectresponse). The NordVPN daemon will send the following two messages, it 
is up to the programmer to choose whether to react to each message, or to wait for both messages to be received.
  - Connecting
  - Connected / Failed
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## Countries
```go
func (c Client) Countries() (*pb.CountriesResponse, error)
```
### Description
Returns a list of countries containing NordVPN servers.

### Parameters
`Countries` does not take any parameters.

### Returns
- [`CountriesResponse`](#countriesresponse) - Countries containing NordVPN servers.
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## Disconnect
```go
func (c Client) Disconnect() error
```
### Description
Disconnects the current VPN session.

### Parameters
`Disconnect` does not take any parameters.

### Returns
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## FrontendCountries
```go
func (c Client) FrontendCountries() (*pb.FrontendCountriesResponse, error)
```
### Description
Returns a list of countries containing NordVPN servers. This function differs from `Countries` as `FrontendCountries` 
also returns country codes, e.g. ("united_kingdom", "uk").

### Parameters
`FrontendCountries` does not take any parameters.

### Returns
- [`FrontendCountriesResponse`](#frontendcountriesresponse) - Countries containing NordVPN servers.
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## Groups
```go
func (c Client) Groups(req *pb.GroupsRequest) (*pb.GroupsResponse, error)
```
### Description
Returns a list of NordVPN server groups.

### Parameters
- req: [`GroupsRequest`](#groupsrequest) - Contains the parameters for the request.

### Returns
- [`GroupsResponse`](#groupsresponse) - NordVPN server groups.
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## IsLoggedIn
```go
func (c Client) IsLoggedIn() (*pb.IsLoggedInResponse, error)
```
### Description
Returns whether the user is currently logged in.

### Parameters
`IsLoggedIn` does not take any parameters.

### Returns
- [`IsLoggedInResponse`](#isloggedinresponse) - The user's log in status.
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## Login
```go
func (c Client) Login(req *pb.LoginRequest) error
```
### Description
Logs the user in using username and password.

### Parameters
- req: [`LoginRequest`](#loginrequest) - Contains the user's credentials.

### Returns
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## LoginOAuth2
```go
func (c Client) LoginOAuth2() (*pb.LoginOAuth2Response, error)
```
### Description
Returns a link to allow the user in login using OAuth2.

### Parameters
`LoginOAuth2` does not take any parameters.

### Returns
- [`LoginOAuth2Response`](#loginoauth2response) - Contains a OAuth2 link to allow the user to authenticate.
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## Logout
```go
func (c Client) Logout() error
```
### Description
Logs the user out.

### Parameters
`Logout` does not take any parameters.

### Returns
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## Plans
```go
func (c Client) Plans() (*pb.PlansResponse, error)
```
### Description
Returns the available NordVPN subscription plans.

### Parameters
`Plans` does not take any parameters.

### Returns
- [`PlansResponse`](#plansresponse) - NordVPN subscription plans.
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## Ping
```go
func (c Client) Ping() error
```
### Description
Checks if the NordVPN daemon is alive.

### Parameters
`Ping` does not take any parameters.

### Returns
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## RateConnection
```go
func (c Client) RateConnection(req *pb.RateConnectionRequest) error
```
### Description
Rates the current connection.

### Parameters
- req: [`RateConnectionRequest`](#rateconnectionrequest) - The connection's rating.

### Returns
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## SetAutoConnect
```go
func (c Client) SetAutoConnect(req *pb.SetAutoConnectRequest) (*pb.Payload, error)
```
### Description
Configures the NordVPN daemon's auto-connect configuration.

### Parameters
- req: [`SetAutoConnectRequest`](#setautoconnectrequest) - The auto-connect configuration.

### Returns
- [`Payload`](#payload) - Generic NordVPN daemon response.
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## SetCyberSec
```go
func (c Client) SetCyberSec(enabled bool) error
```
### Description
Sets whether CyberSec is enabled or disabled.

### Parameters
- enabled: `bool` - Whether CyberSec should be enabled or disabled.

### Returns
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## SetDefaults
```go
func (c Client) SetDefaults() error
```
### Description
Resets all NordVPN configurations to default. This will not disconnect or log the user out.

### Parameters
`SetDefaults` does not take any parameters.

### Returns
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## SetDns
```go
func (c Client) SetDns(req *pb.SetDNSRequest) error
```
### Description
Configures the NordVPN daemon's DNS settings.

### Parameters
- req: [`SetDNSRequest`](#setdnsrequest) - The DNS configuration.

### Returns
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## SetFirewall
```go
func (c Client) SetFirewall(enabled bool) error
```
### Description
Sets whether the NordVPN firewall is enabled or disabled.

### Parameters
- enabled: `bool` - Whether the firewall should be enabled or disabled.

### Returns
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## SetIpv6
```go
func (c Client) SetIpv6(enabled bool) error
```
### Description
Sets whether IPv6 is enabled or disabled.

### Parameters
- enabled: `bool` - Whether IPv6 should be enabled or disabled.

### Returns
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## SetKillSwitch
```go
func (c Client) SetKillSwitch(enabled bool) error
```
### Description
Sets whether the killswitch is enabled or disabled.

### Parameters
- enabled: `bool` - Whether the killswitch should be enabled or disabled.

### Returns
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## SetNotify
```go
func (c Client) SetNotify(enabled bool) error
```
### Description
Sets whether desktop notifications are enabled or disabled.

### Parameters
- enabled: `bool` - Whether desktop notifications should be enabled or disabled.

### Returns
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## SetObfuscate
```go
func (c Client) SetObfuscate(enabled bool) error
```
### Description
Sets whether obfuscation is enabled or disabled.

### Parameters
- enabled: `bool` - Whether obfuscation should be enabled or disabled.

### Returns
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## SetProtocol
```go
func (c Client) SetProtocol(protocol pb.ProtocolEnum) error
```
### Description
Sets which protocol should be used.

### Parameters
- protocol: `ProtocolEnum` - Which protocol to use. Possible values are ProtocolEnum_UDP and ProtocolEnum_TCP.

### Returns
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## SetTechnology
```go
func (c Client) SetTechnology(technology pb.TechnologyEnum) error
```
### Description
Sets which technology should be used.

### Parameters
- technology: `TechnologyEnum` - Which technology to use. Possible values are TechnologyEnum_OPENVPN and 
TechnologyEnum_NORDLYNX.

### Returns
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## SetWhitelist
```go
func (c Client) SetWhitelist(req *pb.SetWhitelistRequest) error
```
### Description
Configures the NordVPN daemon's whitelist settings.

### Parameters
- req: [`SetWhitelistRequest`](#setwhitelistrequest) - Contains a list of TCP ports, UDP ports and subnets to whitelist.

### Returns
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## Settings
```go
func (c Client) Settings() (*pb.SettingsResponse, error)
```
### Description
Retrieves the current NordVPN settings.

### Parameters
`Settings` does not take any parameters.

### Returns
- [`SettingsResponse`](#settingsresponse) - Contains current NordVPN settings.
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## SettingsProtocols
```go
func (c Client) SettingsProtocols() (*pb.ProtocolsResponse, error)
```
### Description
Retrieves a list of supported protocols.

### Parameters
`SettingsProtocols` does not take any parameters.

### Returns
- [`ProtocolsResponse`](#protocolsresponse) - A list of supported protocols.
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## SettingsTechnologies
```go
func (c Client) SettingsTechnologies() (*pb.TechnologyResponse, error)
```
### Description
Retrieves a list of supported technologies.

### Parameters
`SettingsTechnologies` does not take any parameters.

### Returns
- [`TechnologyResponse`](#technologyresponse) - A list of supported technologies.
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


## Status
```go
func (c Client) Status() (*pb.StatusResponse, error)
```
### Description
Retrieves the current connection status and telemetry.

### Parameters
`Status` does not take any parameters.

### Returns
- [`StatusResponse`](#statusresponse) - The status of the connection with telemetry.
- error - Contains an `Error` if an error occurred, otherwise error will be `nil`


# Structs
Structs and struct names are generated by `protoc`, so you may notice some have weird names and may contain some 
additional fields other than those detailed below. Only structs that are required for the programmer to use this 
library are listed below. Some structs may be visible to the programmer but are actually used internally.
Additionally, these structs contain some additional gRPC status information, which is omitted from details below.

OpenNord contains the following structures:

- [AccountResponse](#accountresponse)
- [CitiesResponse](#citiesresponse)
- [ConnectRequest](#connectrequest)
- [ConnectResponse](#connectresponse)
- [CountriesResponse](#countriesresponse)
- [FrontendCountriesResponse](#frontendcountriesresponse)
- [GroupsRequest](#groupsrequest)
- [GroupsResponse](#groupsresponse)
- [IsLoggedInResponse](#isloggedinresponse)
- [LoginRequest](#loginrequest)
- [LoginOAuth2Response](#loginoauth2response)
- [Payload](#payload)
- [Plan](#plan)
- [PlansResponse](#plansresponse)
- [Ports](#ports)
- [ProtocolsResponse](#protocolsresponse)
- [RateConnectionRequest](#rateconnectionrequest)
- [SetAutoConnectRequest](#setautoconnectrequest)
- [SetDnsRequest](#setdnsrequest)
- [SetWhitelistRequest](#setwhitelistrequest)
- [Settings](#settings-struct)
- [SettingsResponse](#settingsresponse)
- [StatusResponse](#statusresponse)
- [TechnologyResponse](#technologyresponse)
- [Whitelist](#whitelist)


## AccountResponse
```go
type AccountResponse struct {
	Email     string
	ExpiresAt string
}
```

### Fields
- Email: `string` - The email address of the user (must be logged in).
- ExpiresAt: `string` - The time when the user's subscription will expire formatted as `YYYY-MM-DD HH:MM:SS`.


## CitiesResponse
```go
type CitiesResponse struct {
    Cities []string `protobuf:"bytes,2,rep,name=cities,proto3" json:"cities,omitempty"`
}
```

### Fields
- Cities: `[]string` - List of cities.


## ConnectRequest
```go
type ConnectRequest struct {
    ServerTag string    
    Protocol  ProtocolEnum
    Obfuscate bool
    CyberSec  bool
    Dns       []string
    WhiteList *Whitelist
}
```

### Fields
- ServerTag: `string` - The server tag equal to either a city, country or group, e.g. "united_kingdom", "uk", "europe"
- Protocol: `ProtocolEnum` - The protocol to use. Possible values are TechnologyEnum_OPENVPN and 
TechnologyEnum_NORDLYNX.
- Obfuscate: `bool` - Whether the connection should use obfuscation.
- CyberSec: `bool` - Whether the connection should use CyberSec.
- Dns: `[]string` - A list of up to 3 DNS servers.
- WhiteList `*Whitelist` - The whitelist to use for the connection.


## ConnectResponse
```go
type ConnectResponse struct {
    Messages []string
}
```

### Fields
- ServerTag: `[]string` - List of messages about the current stage in the connection process.


## CountriesResponse
```go
type CountriesResponse struct {
	Countries []string
}
```

### Fields
- Countries: `[]string` - List of countries.


## FrontendCountriesResponse
```go
type FrontendCountriesResponse struct {
    Countries []*FrontendCountry
}
```

### Fields
- Countries: `[]*FrontendCountry` - List of countries with their code e.g. ("united_kingdom", "uk").


## GroupsRequest
```go
type GroupsRequest struct {
    Protocol  ProtocolEnum
    Obfuscate bool
}
```

### Fields
- Protocol: `ProtocolEnum` - The protocol to use, possible values are TechnologyEnum_OPENVPN and
TechnologyEnum_NORDLYNX.
- Obfuscate: `bool` - Whether the group needs to support obfuscation.


## GroupsResponse
```go
type GroupsResponse struct {
    Groups []string
}
```

### Fields
- Groups: `[]string` - List of groups.


## IsLoggedInResponse
```go
type IsLoggedInResponse struct {
    IsLoggedIn bool
}
```

### Fields
- IsLoggedIn: `bool` - Whether the user is logged in to NordVPN.


## LoginRequest
```go
type LoginRequest struct {
    Username string
    Password string
}
```

### Fields
- Username: `string` - The user's username / email address.
- Password: `string` - The user's password.


## LoginOAuth2Response
```go
type LoginOAuth2Response struct {
    Url string
}
```

### Fields
- Url: `string` - The OAuth2 link.


## Payload
```go
type Payload struct {
    Data []string
}
```

### Fields
- Data: `string` - List of messages from the daemon.


## Plan
```go
type Plan struct {
    Name        string 
    Description string
    Cost        string
    Currency    string
}
```

### Fields
- Name: `string` - Name of the plan.
- Description: `string` - Description of the plan.
- Cost: `string` - Cost of the plan.
- Currency: `string` - Currency of the given cost.


## PlansResponse
```go
type PlansResponse struct {
	Plans []*Plan
}
```

### Fields
- Plans: `[]*Plan` - List of [`Plan`](#plan)s.


## RateConnectionRequest
```go
type RateConnectionRequest struct {
    Rating uint32
}
```

### Fields
- Rating: `uint32` - The rating of the connection.


## SetAutoConnectRequest
```go
type SetAutoConnectRequest struct {
    ServerTag   string       
    Protocol    ProtocolEnum 
    CyberSec    bool         
    Obfuscate   bool         
    AutoConnect bool         
    Dns         []string     
    Whitelist   *Whitelist   
}
```

### Fields
- ServerTag: `string` - The server tag. This can be either a country, city or group. e.g. "united_kingdom", "uk", 
"Europe"
- Protocol: `ProtocolEnum` - The protocol to use, possible values are Possible values are TechnologyEnum_OPENVPN and
  TechnologyEnum_NORDLYNX.
- CyberSec: `bool` - Whether the connection should use CyberSec.
- Obfuscate: `bool` - Whether the connection should use obfuscation.
- AutoConnect: `bool` - Whether auto-connect should be enabled or disabled.
- Dns: `[]string` - List of up to 3 DNS servers.
- Whitelist: `*Whitelist` - The whitelist to use when connecting.


## SetDNSRequest
```go
type SetDNSRequest struct {
    Dns      []string 
    CyberSec bool     
}
```

### Fields
- DNS: `[]string` - List of up to 3 DNS servers.
- CyberSec: `bool` - Whether CyberSec should be enabled or disabled.


## SetWhitelistRequest
```go
type SetWhitelistRequest struct {
    Whitelist *Whitelist
}
```

### Fields
- Whitelist: `*Whitelist` - The whitelist to use.


## Ports
```go
type Ports struct {
	Udp []int32
	Tcp []int32
}
```

### Fields
- Udp: `[]int32` - List of whitelisted UDP ports.
- Tcp: `[]int32` - List of whitelisted TCP ports.


## ProtocolsResponse
```go
type ProtocolsResponse struct {
	Protocols []string
}
```

## Fields
- Protocols: `[]string` - List of supported protocols.


## Settings {#struct}
```go
type Settings struct {
    Technology  TechnologyEnum 
    Firewall    bool           
    KillSwitch  bool           
    AutoConnect bool           
    Notify      bool           
    Ipv6        bool           
}
```

## Fields
- technology: `TechnologyEnum` - Which technology to use. Possible values are TechnologyEnum_OPENVPN and 
TechnologyEnum_NORDLYNX.
- Firewall: `bool` - Whether the NordVPN firewall is enabled.
- KillSwitch: `bool` - Whether the NordVPN killswitch is enabled.
- AutoConnect: `bool` - Whether auto-connect is enabled.
- Notify: `bool` - Whether desktop notifications are enabled.
- Ipv6: `bool` - Whether IPv6 is enabled.


## SettingsResponse
```go
type SettingsResponse struct {
	Settings *Settings
}
```

## Fields
- Settings: `*Settings` - The settings currently used by the NordVPN daemon.


## StatusResponse
```go
type StatusResponse struct {
	State      string         
	Technology TechnologyEnum 
	Protocol   ProtocolEnum   
	Ip         string         
	Hostname   string         
	Country    string         
	City       string         
	Download   int64          
	Upload     int64          
	Uptime     int64
}
```

## Fields
- State: `string` - The current connection state.
- Technology: `TechnologyEnum` - Which technology to use. Possible values are TechnologyEnum_OPENVPN and
  TechnologyEnum_NORDLYNX.
- protocol: `ProtocolEnum` - Which protocol to use. Possible values are ProtocolEnum_UDP and ProtocolEnum_TCP.
- Ip: `string` - The current server's IP address.
- Hostname: `string` - The current server's IP hostname.
- Country: `string` - The current server's country.
- City: `string` - The current server's city.
- Download: `string` - The number of bytes downloaded in the current session.
- Upload: `string` - The number of bytes uploaded in the current session.
- Uptime: `string` - The duration of the session in nanoseconds.


## TechnologyResponse
```go
type TechnologyResponse struct {
	Technologies []string
}
```

## Fields
- Technologies: `[]string` - List of supported technologies.


## Whitelist
```go
type Whitelist struct {
	Ports   *Ports
	Subnets []string
}
```

### Fields
- Ports: `*Ports` - The ports to whitelist.
- Subnets: `[]string` - The subnets to whitelist (in CIDR notation). e.g. 10.10.10.0/24



# Interfaces
OpenNord defines only one interface, that is [`Daemon_ConnectClient`](#daemon_connectclient) which is detailed below.

## Daemon_ConnectClient
```go
type Daemon_ConnectClient interface {
	Recv() (*ConnectResponse, error)
	grpc.ClientStream
}
```

### Functions {#daemon}
- Recv: `(*ConnectResponse, error)` - Receives the current stage in the connection process, i.e. connecting, connected.


# Example
Below is a simple example that logs the user in, connects to a server in London, gets the status, disconnects and logs out.

```go
package main

import (
	"github.com/adamdb5/opennord"
	"github.com/adamdb5/opennord/pb"
	"io"
	"log"
	"time"
)

func main() { 
	// Create the client
	client, err := opennord.NewOpenNordClient()
	if err != nil {
		log.Fatalf("Could not create client: %s.", err)
	}
	
	// Login
	err = client.Login(&pb.LoginRequest{
		Username: "email@domain.com",
		Password: "password",
	})
	if err != nil {
		log.Fatalf("Could not log in: %s.", err)
	} 
	log.Printf("Successfully logged in.")
	
	// Connect to London
	daemonClient, _ := client.Connect(&pb.ConnectRequest{
		ServerTag: "London",
		Protocol:  pb.ProtocolEnum_TCP,
	})
	
	// Log the connection progress
	for {
		msg, err := daemonClient.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not connect: %s.", err)
		}
		log.Printf("Connecting to: %s (%s).", msg.GetMessages()[0], msg.GetMessages()[1])
	}
	
	// Check our connection status
	status, err := client.Status()
	if err != nil {
		log.Fatalf("Could not get status: %s.", err)
	}
	log.Printf("Status: %s.", status.GetState())
	
	// Wait for a bit 
	time.Sleep(5 * time.Second)
	
	// Disconnect
	err = client.Disconnect()
	if err != nil {
		log.Fatalf("Could not disconnect: %s.", err)
	}
	log.Printf("Successfully disconnected.")
	
	// Logout
	err = client.Logout()
	if err != nil {
		log.Fatalf("Could not logout: %s.", err)
	}
	log.Printf("Successfully logged out.")
}
```

Output:
```text
2022/01/06 21:19:29 Successfully logged in.
2022/01/06 21:19:30 Connecting to: United Kingdom #2203 (uk2203.nordvpn.com).
2022/01/06 21:19:30 Connecting to: United Kingdom #2203 (uk2203.nordvpn.com).
2022/01/06 21:19:30 Status: Connected.
2022/01/06 21:19:35 Successfully disconnected.
2022/01/06 21:19:35 Successfully logged out.
```