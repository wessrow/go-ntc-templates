package mikrotik_routeros

type IpDhcpServerLeasePrintWithoutPaging struct {
	Index       string `json:"INDEX"`
	Flags       string `json:"FLAGS"`
	Ip_address  string `json:"IP_ADDRESS"`
	Mac_address string `json:"MAC_ADDRESS"`
	Hostname    string `json:"HOSTNAME"`
	Rate_limit  string `json:"RATE_LIMIT"`
	Server      string `json:"SERVER"`
	Status      string `json:"STATUS"`
	Last_seen   string `json:"LAST_SEEN"`
}

var IpDhcpServerLeasePrintWithoutPaging_Template string = `Value Key INDEX (\d+)
Value FLAGS ([XRDB]+)
Value IP_ADDRESS ([\w.:]+)
Value MAC_ADDRESS (([0-9a-fA-F]{2}[:]){5}([0-9a-fA-F]{2}))
Value HOSTNAME (\S+)
Value SERVER (\w+?)
Value RATE_LIMIT (\d+)
Value STATUS (waiting|testing|authorizing|busy|offered|bound)
Value LAST_SEEN (\w+)

Start
  ^\s*#\s+ADDRESS\s+MAC-ADDRESS\s+HOST-NAME\s+SERVER\s+RATE-LIMIT\s+STATUS\s+LAST-SEEN\s+$$ -> EntriesTable

EntriesTable
  ^\s*${INDEX}\s+(${RATE_LIMIT})?\s+(${STATUS})?\s+(${LAST_SEEN})?\s*$$ -> Record
  ^\s*${INDEX}\s(${FLAGS})?\s{,4}${IP_ADDRESS}\s+(${RATE_LIMIT})?\s+${STATUS}\s+${LAST_SEEN}\s*$$ -> Record
  ^\s*${INDEX}\s(${FLAGS})?\s{,4}(${IP_ADDRESS})?\s+${MAC_ADDRESS}\s+(${RATE_LIMIT})?\s+${STATUS}\s+${LAST_SEEN}\s*$$ -> Record
  ^\s*${INDEX}\s(${FLAGS})?\s{,4}(${IP_ADDRESS})?\s+(${MAC_ADDRESS}|\s{17})\s(${HOSTNAME})?\s+(${SERVER})?\s+(${RATE_LIMIT})?\s+${STATUS}\s+${LAST_SEEN}\s*$$ -> Record
  ^\s*(?:\d{2}:){2}\d{2}\s+echo:\s*.*$$ -> Next
  ^. -> Error
`
