package brocade_fastiron

type ShowTrunk struct {
	State7 string `json:"STATE7"`
	Link1  string `json:"LINK1"`
	Link3  string `json:"LINK3"`
	Link5  string `json:"LINK5"`
	Link8  string `json:"LINK8"`
	State2 string `json:"STATE2"`
	State4 string `json:"STATE4"`
	Port2  string `json:"PORT2"`
	Port7  string `json:"PORT7"`
	Port8  string `json:"PORT8"`
	Link6  string `json:"LINK6"`
	Link7  string `json:"LINK7"`
	State1 string `json:"STATE1"`
	State5 string `json:"STATE5"`
	Port1  string `json:"PORT1"`
	Link2  string `json:"LINK2"`
	Link4  string `json:"LINK4"`
	Port5  string `json:"PORT5"`
	Port6  string `json:"PORT6"`
	State3 string `json:"STATE3"`
	State6 string `json:"STATE6"`
	State8 string `json:"STATE8"`
	Id     string `json:"ID"`
	Port3  string `json:"PORT3"`
	Port4  string `json:"PORT4"`
}

var ShowTrunk_Template string = `Value Filldown ID ([0-9\/]+)
Value Required PORT1 ([0-9\/]+)
Value PORT2 ([0-9\/]+)
Value PORT3 ([0-9\/]+)
Value PORT4 ([0-9\/]+)
Value PORT5 ([0-9\/]+)
Value PORT6 ([0-9\/]+)
Value PORT7 ([0-9\/]+)
Value PORT8 ([0-9\/]+)
Value LINK1 (\w+)
Value LINK2 (\w+)
Value LINK3 (\w+)
Value LINK4 (\w+)
Value LINK5 (\w+)
Value LINK6 (\w+)
Value LINK7 (\w+)
Value LINK8 (\w+)
Value STATE1 (\w+)
Value STATE2 (\w+)
Value STATE3 (\w+)
Value STATE4 (\w+)
Value STATE5 (\w+)
Value STATE6 (\w+)
Value STATE7 (\w+)
Value STATE8 (\w+)

Start
  ^Operational trunks: -> BODY

BODY
  ^$$ -> Continue.Record
  ^Trunk ID: ${ID}
  ^Active Ports: 2 -> CASE2
  ^Active Ports: 3 -> CASE3
  ^Active Ports: 4 -> CASE4
  ^Active Ports: 5 -> CASE5
  ^Active Ports: 6 -> CASE6
  ^Active Ports: 7 -> CASE7
  ^Active Ports: 8 -> CASE8

CASE2
  ^Ports\s+Link_Status -> ICX
  ^Ports\s+${PORT1}\s+${PORT2}
  ^Link_Status\s+${LINK1}\s+${LINK2}
  ^port_state\s+${STATE1}\s+${STATE2} -> BODY

CASE3
  ^Ports\s+Link_Status -> ICX    
  ^Ports\s+${PORT1}\s+${PORT2}\s+${PORT3}
  ^Link_Status\s+${LINK1}\s+${LINK2}\s+${LINK3}
  ^port_state\s+${STATE1}\s+${STATE2}\s+${STATE3} -> BODY

CASE4
  ^Ports\s+Link_Status -> ICX    
  ^Ports\s+${PORT1}\s+${PORT2}\s+${PORT3}\s+${PORT4}
  ^Link_Status\s+${LINK1}\s+${LINK2}\s+${LINK3}\s+${LINK4}
  ^port_state\s+${STATE1}\s+${STATE2}\s+${STATE3}\s+${STATE4} -> BODY

CASE5
  ^Ports\s+Link_Status -> ICX    
  ^Ports\s+${PORT1}\s+${PORT2}\s+${PORT3}\s+${PORT4}\s+${PORT5}
  ^Link_Status\s+${LINK1}\s+${LINK2}\s+${LINK3}\s+${LINK4}\s+${LINK5}
  ^port_state\s+${STATE1}\s+${STATE2}\s+${STATE3}\s+${STATE4}\s+${STATE5} -> BODY

CASE6
  ^Ports\s+Link_Status -> ICX    
  ^Ports\s+${PORT1}\s+${PORT2}\s+${PORT3}\s+${PORT4}\s+${PORT5}\s+${PORT6}
  ^Link_Status\s+${LINK1}\s+${LINK2}\s+${LINK3}\s+${LINK4}\s+${LINK5}\s+${LINK6}
  ^port_state\s+${STATE1}\s+${STATE2}\s+${STATE3}\s+${STATE4}\s+${STATE5}\s+${STATE6} -> BODY

CASE7
  ^Ports\s+Link_Status -> ICX    
  ^Ports\s+${PORT1}\s+${PORT2}\s+${PORT3}\s+${PORT4}\s+${PORT5}\s+${PORT6}\s+${PORT7}
  ^Link_Status\s+${LINK1}\s+${LINK2}\s+${LINK3}\s+${LINK4}\s+${LINK5}\s+${LINK6}\s+${LINK7}
  ^port_state\s+${STATE1}\s+${STATE2}\s+${STATE3}\s+${STATE4}\s+${STATE5}\s+${STATE6}\s+${STATE7} -> BODY

CASE8
  ^Ports\s+Link_Status -> ICX    
  ^Ports\s+${PORT1}\s+${PORT2}\s+${PORT3}\s+${PORT4}\s+${PORT5}\s+${PORT6}\s+${PORT7}\s+${PORT8}
  ^Link_Status\s+${LINK1}\s+${LINK2}\s+${LINK3}\s+${LINK4}\s+${LINK5}\s+${LINK6}\s+${LINK7}\s+${LINK8}
  ^port_state\s+${STATE1}\s+${STATE2}\s+${STATE3}\s+${STATE4}\s+${STATE5}\s+${STATE6}\s+${STATE7}\s+${STATE8} -> BODY

ICX
  ^${PORT1}\s+${LINK1}\s+${STATE1} -> ICX2

ICX2
  ^${PORT2}\s+${LINK2}\s+${STATE2} -> ICX3

ICX3
  ^${PORT3}\s+${LINK3}\s+${STATE3} -> ICX4

ICX4
  ^${PORT4}\s+${LINK4}\s+${STATE4} -> ICX5

ICX5
  ^${PORT5}\s+${LINK5}\s+${STATE5} -> ICX6

ICX6 
  ^${PORT6}\s+${LINK6}\s+${STATE6} -> ICX7

ICX7
  ^${PORT7}\s+${LINK7}\s+${STATE7} -> ICX8

ICX8
  ^${PORT8}\s+${LINK8}\s+${STATE8} -> BODY

`
