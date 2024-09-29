package mikrotik_routeros 

type IpDhcpServerLeasePrint struct {
	Number	string	`json:"NUMBER"`
	Flags	string	`json:"FLAGS"`
	Address	string	`json:"ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Host_name	string	`json:"HOST_NAME"`
	Server	string	`json:"SERVER"`
	Rate_limit	string	`json:"RATE_LIMIT"`
	Status	string	`json:"STATUS"`
	Last_seen	string	`json:"LAST_SEEN"`
}

var IpDhcpServerLeasePrint_Template = `Value NUMBER (\d+)
Value FLAGS ([XRDB]+)
Value ADDRESS (\S+)
Value MAC_ADDRESS ((?:[0-9a-fA-F]{2}:){5}[0-9a-fA-F]{2})
Value HOST_NAME (\S+)
Value SERVER (\S+)
Value RATE_LIMIT (\S+)
Value STATUS (\S+)
Value LAST_SEEN (\S+)

Start
  ^\s*#\s+ADDRESS\s+MAC-ADDRESS\s+HOST-NAME\s+SERVER\s+RATE-LIMIT\s+STATUS\s+LAST-SEEN\s*$$ -> EntriesTable

EntriesTable
  # Impossible to separate HOST-NAME and SERVER? Have the same regex, can be empty,
  # are next to each other -> use '...print terse...' version of the command
  ^\s*${NUMBER}(?:\s+${FLAGS})?\s+${ADDRESS}(?:\s+${MAC_ADDRESS})?(?:\s+${HOST_NAME})?\s+${SERVER}(?:\s+${RATE_LIMIT})?\s+${STATUS}\s+${LAST_SEEN}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`