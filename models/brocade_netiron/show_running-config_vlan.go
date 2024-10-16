package brocade_netiron

type ShowRunningConfigVlan struct {
	Untaggedports    string `json:"untaggedports"`
	Metroname2       string `json:"metroname2"`
	Metroconfdiag3   string `json:"metroconfdiag3"`
	Metroname4       string `json:"metroname4"`
	Metroid2         string `json:"metroid2"`
	Metroconfmaster2 string `json:"metroconfmaster2"`
	Metroname3       string `json:"metroname3"`
	Metroconfmaster4 string `json:"metroconfmaster4"`
	Metroid3         string `json:"metroid3"`
	Metroconfenable4 string `json:"metroconfenable4"`
	Metroconfdiag4   string `json:"metroconfdiag4"`
	Vlanname         string `json:"vlanname"`
	Metroconfdiag1   string `json:"metroconfdiag1"`
	Metroconfenable3 string `json:"metroconfenable3"`
	Metroconfmaster3 string `json:"metroconfmaster3"`
	Metroconfmaster1 string `json:"metroconfmaster1"`
	Metrointerfaces1 string `json:"metrointerfaces1"`
	Metroname1       string `json:"metroname1"`
	Metroconfdiag2   string `json:"metroconfdiag2"`
	Metrointerfaces4 string `json:"metrointerfaces4"`
	Vlanid           string `json:"vlanid"`
	Metroid1         string `json:"metroid1"`
	Metroconfenable1 string `json:"metroconfenable1"`
	Metrointerfaces2 string `json:"metrointerfaces2"`
	Metroconfenable2 string `json:"metroconfenable2"`
	Taggedports      string `json:"taggedports"`
	Associatedve     string `json:"associatedve"`
	Metrointerfaces3 string `json:"metrointerfaces3"`
	Metroid4         string `json:"metroid4"`
}

var ShowRunningConfigVlan_Template string = `Value vlanid (\d+)
Value vlanname ([a-zA-Z0-9\-\"\ ]+)
Value untaggedports ([0-9\/\ eto]+)
Value taggedports ([0-9\/\ eto]+)
Value associatedve (\d+)
Value metroid1 (\d+)
Value metrointerfaces1 ([0-9\/\ etho]+)
Value metroconfenable1 (enable)
Value metroconfmaster1 (master)
Value metroconfdiag1 (diagnostics)
Value metroname1 ([a-zA-Z0-9\-\"\ ]+)
Value metroid2 (\d+)
Value metrointerfaces2 ([0-9\/\ etho]+)
Value metroconfenable2 (enable)
Value metroconfmaster2 (master)
Value metroconfdiag2 (diagnostics)
Value metroname2 ([a-zA-Z0-9\-\"\ ]+)
Value metroid3 (\d+)
Value metrointerfaces3 ([0-9\/\ etho]+)
Value metroconfenable3 (enable)
Value metroconfmaster3 (master)
Value metroconfdiag3 (diagnostics)
Value metroname3 ([a-zA-Z0-9\-\"\ ]+)
Value metroid4 (\d+)
Value metrointerfaces4 ([0-9\/\ etho]+)
Value metroconfenable4 (enable)
Value metroconfmaster4 (master)
Value metroconfdiag4 (diagnostics)
Value metroname4 ([a-zA-Z0-9\-\"\ ]+)

Start
  ^\w -> Continue.Record
  ^vlan ${vlanid} -> Continue
  ^vlan ${vlanid} name ${vlanname} -> Continue
  ^\s+tagged\s+${taggedports} -> Continue
  ^\s+untagged\s+${untaggedports} -> Continue
  ^\s+router-interface\s+ve\s+${associatedve} -> Continue
  ^! -> Start
  ^\s+metro-ring\s+${metroid1} -> METRO1

METRO1
  ^\s+ring-interfaces\s+${metrointerfaces1} -> Continue
  ^\s+${metroconfenable1} -> Continue
  ^\s+${metroconfmaster1} -> Continue
  ^\s+${metroconfdiag1} -> Continue
  ^\s+name\s+${metroname1} -> Continue
  ^! -> Start
  ^\s+metro-ring\s+${metroid2} -> METRO2

METRO2
  ^\s+ring-interfaces\s+${metrointerfaces2} -> Continue
  ^\s+${metroconfenable2} -> Continue
  ^\s+${metroconfmaster2} -> Continue
  ^\s+${metroconfdiag2} -> Continue
  ^\s+name\s+${metroname2} -> Continue
  ^! -> Start
  ^\s+metro-ring\s+${metroid3} -> METRO3

METRO3
  ^\s+ring-interfaces\s+${metrointerfaces3} -> Continue
  ^\s+${metroconfenable3} -> Continue
  ^\s+${metroconfmaster3} -> Continue
  ^\s+${metroconfdiag3} -> Continue
  ^\s+name\s+${metroname3} -> Continue
  ^! -> Start
  ^\s+metro-ring\s+${metroid4} -> METRO4

METRO4
  ^\s+ring-interfaces\s+${metrointerfaces4} -> Continue
  ^\s+${metroconfenable4} -> Continue
  ^\s+${metroconfmaster4} -> Continue
  ^\s+${metroconfdiag4} -> Continue
  ^\s+name\s+${metroname4} -> Continue
`
