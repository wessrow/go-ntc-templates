package cisco_nxos 

type ShowSpanningTreeRoot struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Priority	string	`json:"PRIORITY"`
	Root_id	string	`json:"ROOT_ID"`
	Root_cost	string	`json:"ROOT_COST"`
	Hello_time	string	`json:"HELLO_TIME"`
	Max_age	string	`json:"MAX_AGE"`
	Fwd_dly	string	`json:"FWD_DLY"`
	Root_port	string	`json:"ROOT_PORT"`
}

var ShowSpanningTreeRoot_Template = `Value Required VLAN_ID (\S+)
Value Required PRIORITY (\S+)
Value Required ROOT_ID (\S+)
Value Required ROOT_COST (\d+)
Value Required HELLO_TIME (\d+)
Value Required MAX_AGE (\d+)
Value Required FWD_DLY (\d+)
Value Required ROOT_PORT (.*)

Start
  ^\s*Root\s*Hello\s*Max\s*Fwd$$
  ^\s*Vlan\s*Root\sID\s*Cost\s*Time\s*Age\s*Dly\s*Root\s*Port$$
  ^\s*- 
  ^\s*${VLAN_ID}\s+${PRIORITY}\s+${ROOT_ID}+\s+${ROOT_COST}+\s+${HELLO_TIME}+\s+${MAX_AGE}+\s+${FWD_DLY}+\s+${ROOT_PORT} -> Record
 ^. -> Error`