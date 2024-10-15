package brocade_fastiron

type ShowLagBrief struct {
	Deployed         string `json:"deployed"`
	Trunkid          string `json:"trunkid"`
	Primaryinterface string `json:"primaryinterface"`
	Portlist         string `json:"portlist"`
	Portlistcont     string `json:"portlistcont"`
	Lagnameshort     string `json:"lagnameshort"`
	Lagtype          string `json:"lagtype"`
}

var ShowLagBrief_Template string = `Value lagnameshort ([0-9\/\sA-Za-z\-\.]+)
Value lagtype (dynamic|static)
Value deployed (Y|N)
Value trunkid (\d+)
Value primaryinterface ([0-9\/]+)
Value portlist ([eto0-9\/\s\r\n]+)
Value portlistcont ([eto0-9\/\s\r\n]+)

Start
  ^\S -> Continue.Record
  ^${lagnameshort}${lagtype}\s+${deployed}\s+${trunkid}\s+${primaryinterface}\s+${portlist}
  ^\s+${portlistcont}`
