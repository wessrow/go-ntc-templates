package linux_ip 

type LinkShow struct {
	Id	string	`json:"ID"`
	Interface	string	`json:"INTERFACE"`
	Flags	string	`json:"FLAGS"`
	Mtu	string	`json:"MTU"`
	Qdisc	string	`json:"QDISC"`
	State	string	`json:"STATE"`
	Mode	string	`json:"MODE"`
	Group	string	`json:"GROUP"`
	Qlen	string	`json:"QLEN"`
	Master	string	`json:"MASTER"`
	Type	string	`json:"TYPE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Broadcast	string	`json:"BROADCAST"`
	Alias	string	`json:"ALIAS"`
}

var LinkShow_Template = `Value Required ID (\d+)
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