package linux

type IpLinkShow struct {
	Id          string `json:"ID"`
	Mtu         string `json:"MTU"`
	State       string `json:"STATE"`
	Mode        string `json:"MODE"`
	Interface   string `json:"INTERFACE"`
	Flags       string `json:"FLAGS"`
	Type        string `json:"TYPE"`
	Alias       string `json:"ALIAS"`
	Qdisc       string `json:"QDISC"`
	Group       string `json:"GROUP"`
	Qlen        string `json:"QLEN"`
	Master      string `json:"MASTER"`
	Mac_address string `json:"MAC_ADDRESS"`
	Broadcast   string `json:"BROADCAST"`
}

var IpLinkShow_Template string = `Value Required ID (\d+)
Value Required INTERFACE ([^:]+)
Value Required FLAGS (\S+)
Value Required MTU (\d+)
Value Required QDISC (\S+)
Value Required STATE (\S+)
Value MODE (\S+)
Value GROUP (\S+)
Value QLEN (\d+)
Value MASTER (\S+)
Value Required TYPE (\S+)
Value MAC_ADDRESS ([a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2})
Value BROADCAST ([a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2})
Value ALIAS (.+)

Start
  ^\d+ -> Continue.Record
  ^${ID}: ${INTERFACE}: <${FLAGS}> mtu ${MTU} qdisc ${QDISC}(?: master ${MASTER})? state ${STATE}(?: mode ${MODE})?(?: group ${GROUP})?(?: qlen ${QLEN})?\s*$$ -> Continue
  # Link type and MAC address
  ^\s+link/${TYPE}\s*$$
  ^\s+link/${TYPE} ${MAC_ADDRESS} brd ${BROADCAST}\s*$$
  ^\s+alias\s+${ALIAS}$$
  ^\d+
  ^\s+
  ^. -> Error

`
