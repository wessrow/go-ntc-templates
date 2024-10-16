package cisco_nxos

type ShowBgpVrfAllIpv4UnicastNeighborsRoutes struct {
	Status         string `json:"STATUS"`
	Route_source   string `json:"ROUTE_SOURCE"`
	Weight         string `json:"WEIGHT"`
	Next_hop       string `json:"NEXT_HOP"`
	Metric         string `json:"METRIC"`
	Local_pref     string `json:"LOCAL_PREF"`
	As_path        string `json:"AS_PATH"`
	Origin         string `json:"ORIGIN"`
	Path_selection string `json:"PATH_SELECTION"`
	Prefix         string `json:"PREFIX"`
	Prefix_length  string `json:"PREFIX_LENGTH"`
}

var ShowBgpVrfAllIpv4UnicastNeighborsRoutes_Template string = `Value STATUS ([sxSdh*])
Value PATH_SELECTION ([> |&|])
Value ROUTE_SOURCE ([ieclarI])
Value Filldown PREFIX ([A-Fa-f0-9:\.]+)
Value Filldown PREFIX_LENGTH (\d+)
Value NEXT_HOP ([A-Fa-f0-9:\.]+)
Value METRIC (\d+)
Value LOCAL_PREF (\d+)
Value WEIGHT (\d+)
Value AS_PATH ((\d+|\{\d+)(\.\d+?|\s.*?)*)
Value ORIGIN ([ie\?\|&])

Start
  ^Can't\s+find\s+neighbor
  ^Peer\s+
  ^BGP\s+(routing|table)\s 
  ^Status:
  ^Path type:
  ^Origin codes:
  ^\s*Network\s+Next Hop\s+Metric\s+LocPrf\s+Weight\s+Path -> Bgp_table
  ^. -> Error

Bgp_table
  # Match first when there is no network, since previous line had it already (compliment and filldown below)
  # Older table format, with blank values parsing not reliable. (cisco_nxos_show_ip_bgp.raw)
  #             Network            Next Hop            Metric     LocPrf     Weight Path
  # Example: *>e10.10.101.4/30     10.10.2.1                                      0 64102 i
  # Example: *>e                   10.10.1.1                                      0 64101 64002 i
  ^${STATUS}${PATH_SELECTION}${ROUTE_SOURCE}${PREFIX}/${PREFIX_LENGTH}(\s{0,12}${NEXT_HOP}\s{4}|\s{20})(\s{2,14}${METRIC}|\s{8})(\s{1,18}${LOCAL_PREF}|\s{8})(\s{1,22}${WEIGHT})(\s${AS_PATH})?\s${ORIGIN}$$ -> Record
  ^${STATUS}${PATH_SELECTION}${ROUTE_SOURCE}\s\s\s{16}(\s{0,12}${NEXT_HOP}\s{4}|\s{20})(\s{2,14}${METRIC}|\s{8})(\s{1,18}${LOCAL_PREF}|\s{8})(\s{1,22}${WEIGHT})(\s${AS_PATH})?\s${ORIGIN}$$ -> Record
  #
  # Match newer format (cisco_nxos_show_ip_bgp1.raw)
  # Example: * e0.0.0.0/0 1.2.3.4 4294967295 0 12345.102 
  ^${STATUS}${PATH_SELECTION}${ROUTE_SOURCE}${PREFIX}/${PREFIX_LENGTH}\s${NEXT_HOP}\s${METRIC}\s${LOCAL_PREF}\s${WEIGHT}(\s${AS_PATH})?\s${ORIGIN}$$ -> Record
  ^${STATUS}${PATH_SELECTION}${ROUTE_SOURCE}\s${NEXT_HOP}\s${METRIC}\s${LOCAL_PREF}\s${WEIGHT}(\s${AS_PATH})?\s${ORIGIN}$$ -> Record
  ^. -> Error

EOF
`
