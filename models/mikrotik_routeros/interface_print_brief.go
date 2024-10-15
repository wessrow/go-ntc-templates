package mikrotik_routeros

type InterfacePrintBrief struct {
	Actual_mtu  string `json:"ACTUAL_MTU"`
	L2mtu       string `json:"L2MTU"`
	Max_l2mtu   string `json:"MAX_L2MTU"`
	Status      string `json:"STATUS"`
	Dynamic     string `json:"DYNAMIC"`
	Slave       string `json:"SLAVE"`
	Name        string `json:"NAME"`
	Type        string `json:"TYPE"`
	Mac_address string `json:"MAC_ADDRESS"`
	Id          string `json:"ID"`
}

var InterfacePrintBrief_Template string = `Value ID (\d+)
Value DYNAMIC (D)
Value STATUS (X|R)
Value SLAVE (S)
Value NAME (\S+)
Value TYPE (\S+)
Value ACTUAL_MTU (\d+)
Value L2MTU (\d+)
Value MAX_L2MTU (\d+)
Value MAC_ADDRESS ([0-9a-fA-F]{2}(?::[0-9a-fA-F]{2}){5})

Start
  ^\s*Flags:\s+D\s+-\s+dynamic,\s+X\s+-\s+disabled,\s+R\s+-\s+running,\s+S\s+-\s+slave\s*$$
  ^\s*#\s+NAME\s+TYPE\s+ACTUAL-MTU\s+L2MTU\s+MAX-L2MTU\s+MAC-ADDRESS\s*$$ -> IfaceTable
  ^\s*$$
  ^. -> Error

IfaceTable
  ^\s*${ID}(?:\s+${DYNAMIC})?(?:\s*${STATUS})?(?:\s*${SLAVE})?\s+${NAME}\s+${TYPE}\s+${ACTUAL_MTU}(?:\s+${L2MTU})?(?:\s+${MAX_L2MTU})?(?:\s+${MAC_ADDRESS})?\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
