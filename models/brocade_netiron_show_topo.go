package models

type BrocadeNetironShowTopo struct {
	Topogroup	string	`json:"topogroup"`
	Topohwindex	string	`json:"topohwindex"`
	Mastervlan	string	`json:"mastervlan"`
	L2proto	string	`json:"l2proto"`
	Vplsvlanpresent	string	`json:"vplsvlanpresent"`
	Membervlans	string	`json:"membervlans"`
	Membergroup	string	`json:"membergroup"`
	Controlports	string	`json:"controlports"`
	Freeports	string	`json:"freeports"`
}

var BrocadeNetironShowTopo_Template = `Value topogroup (\d+)
Value topohwindex (\d+)
Value mastervlan (\d+)
Value l2proto (MRP|STP|RSTP|VSRP|ERP)
Value vplsvlanpresent (TRUE|FALSE)
Value membervlans ([0-9\ to]+)
Value membergroup ([0-9a-zA-Z\/\ ]+)
Value controlports ([0-9\ \/eto]+)
Value freeports ([0-9\ \/eto]+)

Start
  ^Topology Group\s+${topogroup}
  ^Topo HW Index\s+${topohwindex}
  ^Master VLAN\s+:\s+${mastervlan}
  ^L2 Protocol\s+:\s+${l2proto}
  ^VPLS VLAN exist\s+:\s+${vplsvlanpresent}
  ^Member VLAN\s+:\s+${membervlans}
  ^Member Group\s+:\s+${membergroup}
  ^Control Ports\s+:\s+${controlports}
  ^Free -> Record
`