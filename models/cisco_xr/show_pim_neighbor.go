package cisco_xr

type ShowPimNeighbor struct {
	Pri       string `json:"PRI"`
	Flags     string `json:"FLAGS"`
	Neighbor  string `json:"NEIGHBOR"`
	Interface string `json:"INTERFACE"`
	Uptime    string `json:"UPTIME"`
	Expires   string `json:"EXPIRES"`
	Dr        string `json:"DR"`
}

var ShowPimNeighbor_Template string = `Value NEIGHBOR (\S+)
Value INTERFACE (\S+)
Value UPTIME ((\S+(\s\S+)*))
Value EXPIRES (\S+)
Value DR (\d+)
Value PRI (\S+)
Value FLAGS ((\S+(\s\S+)*))

Start
  ^${NEIGHBOR}(\*)*\s+${INTERFACE}\s+${UPTIME}\s+${EXPIRES}\s+${DR}\s+(\(${PRI}\))?\s+${FLAGS} -> Record`
