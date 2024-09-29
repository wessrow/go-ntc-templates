package cisco_xr 

type ShowRsvpNeighbors struct {
	Global_neighbor	string	`json:"GLOBAL_NEIGHBOR"`
	Interface_neighbor	string	`json:"INTERFACE_NEIGHBOR"`
	Interface	string	`json:"INTERFACE"`
}

var ShowRsvpNeighbors_Template = `Value Filldown GLOBAL_NEIGHBOR (\S+)
Value Required INTERFACE_NEIGHBOR (\S+)
Value INTERFACE (\S+)

Start
  ^.*Global Neighbor:\s+${GLOBAL_NEIGHBOR}
  ^.*\-+ -> Parse

Parse
  ^\s+${INTERFACE_NEIGHBOR}\s+${INTERFACE} -> Record`