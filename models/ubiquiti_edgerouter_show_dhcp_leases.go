package models

type UbiquitiEdgerouterShowDhcpLeases struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Expiration	string	`json:"EXPIRATION"`
	Pool	string	`json:"POOL"`
	Host	string	`json:"HOST"`
}

var UbiquitiEdgerouterShowDhcpLeases_Template = `Value IP_ADDRESS ((?:[0-9]{1,3}\.){3}[0-9]{1,3})
Value MAC_ADDRESS ((?:[0-9a-fA-F]{2}\:){5}[0-9a-fA-F]{2})
Value EXPIRATION (.+\S)
Value POOL (\S+)
Value HOST (\S+)

Start
  ^IP\saddress\s+Hardware\sAddress\s+Lease\sexpiration\s+Pool\s+Client\sName\s*$$
  ^-+\s+-+\s+-+\s+-+\s+-+\s*$$ -> Entries

Entries
  ^${IP_ADDRESS}\s+${MAC_ADDRESS}\s+${EXPIRATION}\s+${POOL}\s+${HOST}\s*$$ -> Record
  ^. -> Error

`