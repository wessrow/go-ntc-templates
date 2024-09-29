package models

type MikrotikRouterosIpNeighborPrintDetail struct {
	Number	string	`json:"NUMBER"`
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Ipv4_address	string	`json:"IPV4_ADDRESS"`
	Ipv6_address	string	`json:"IPV6_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Identity	string	`json:"IDENTITY"`
	Platform	string	`json:"PLATFORM"`
	Version	string	`json:"VERSION"`
	Unpack	string	`json:"UNPACK"`
	Age	string	`json:"AGE"`
	Uptime	string	`json:"UPTIME"`
	Software_id	string	`json:"SOFTWARE_ID"`
	Board	string	`json:"BOARD"`
	Ipv6	string	`json:"IPV6"`
	Interface_name	string	`json:"INTERFACE_NAME"`
	System_description	string	`json:"SYSTEM_DESCRIPTION"`
	System_caps	string	`json:"SYSTEM_CAPS"`
	System_caps_enabled	string	`json:"SYSTEM_CAPS_ENABLED"`
}

var MikrotikRouterosIpNeighborPrintDetail_Template = `Value NUMBER (\d+)
Value INTERFACE (\S+)
Value IP_ADDRESS (\S+)
Value IPV4_ADDRESS (\S+)
Value IPV6_ADDRESS (\S+)
Value MAC_ADDRESS ([a-zA-Z0-9]{2}(:[a-zA-Z0-9]{2}){5})
Value IDENTITY (.*?)
Value PLATFORM (.*?)
Value VERSION (.*?)
Value UNPACK (none|simple|uncompressed-headers|uncompressed-all)
Value AGE ([0-9a-z]+)
Value UPTIME ([0-9a-z]+)
Value SOFTWARE_ID (\S+)
Value BOARD (\S+)
Value IPV6 (yes|no)
Value INTERFACE_NAME (.*?(?="))
Value SYSTEM_DESCRIPTION (.*?(?="))
Value SYSTEM_CAPS ([a-zA-Z0-9,-]*)
Value SYSTEM_CAPS_ENABLED ([a-zA-Z0-9,-]*)

Start
  ^\s*\d+ -> Continue.Record
  ^\s*${NUMBER} -> Continue
  ^.*\s+interface=${INTERFACE} -> Continue
  ^.*\s+address=${IP_ADDRESS} -> Continue
  ^.*\s+address4=${IPV4_ADDRESS} -> Continue
  ^.*\s+address6=${IPV6_ADDRESS} -> Continue
  ^.*\s+mac-address=${MAC_ADDRESS} -> Continue
  ^.*\s+identity="${IDENTITY}" -> Continue
  ^.*\s+platform="${PLATFORM}" -> Continue
  ^.*\s+version="${VERSION}" -> Continue
  ^.*\s+unpack=${UNPACK} -> Continue
  ^.*\s+age=${AGE} -> Continue
  ^.*\s+uptime=${UPTIME} -> Continue
  ^.*\s+software-id="${SOFTWARE_ID}" -> Continue
  ^.*\s+board="${BOARD}" -> Continue
  ^.*\s+ipv6=${IPV6} -> Continue
  ^.*\s+interface-name="${INTERFACE_NAME}" -> Continue
  ^.*\s+system-description="${SYSTEM_DESCRIPTION}" -> Continue
  ^.*\s+system-caps="?${SYSTEM_CAPS}"? -> Continue
  ^.*\s+system-caps-enabled="?${SYSTEM_CAPS_ENABLED}"? -> Continue
  ^\s*
  ^. -> Error

`