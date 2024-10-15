package cisco_xr

type ShowBgpVrfAllIpv4UnicastSummary struct {
	State           string `json:"STATE"`
	Local_as_number string `json:"LOCAL_AS_NUMBER"`
	Vrf             string `json:"VRF"`
	Rd              string `json:"RD"`
}

var ShowBgpVrfAllIpv4UnicastSummary_Template string = `Value Required VRF (\S+)
Value RD (\S+)
Value STATE (\w+)
Value LOCAL_AS_NUMBER (\S+)

Start
  ^BGP\s+VRF\s+${VRF},\s+state:\s+${STATE}
  ^BGP Route Distinguisher:\s+${RD}
  ^.*,\slocal\sAS\snumber\s${LOCAL_AS_NUMBER} -> Record
`
