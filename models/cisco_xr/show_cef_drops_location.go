package cisco_xr 

type ShowCefDropsLocation struct {
	Location	string	`json:"LOCATION"`
	Unresolved_drops	string	`json:"UNRESOLVED_DROPS"`
	Unsupported_drops	string	`json:"UNSUPPORTED_DROPS"`
	Null0_drops	string	`json:"NULL0_DROPS"`
	No_route_drops	string	`json:"NO_ROUTE_DROPS"`
	No_adjacency_drops	string	`json:"NO_ADJACENCY_DROPS"`
	Checksum_error_drops	string	`json:"CHECKSUM_ERROR_DROPS"`
	Rpf_drops	string	`json:"RPF_DROPS"`
	Rpf_suppressed_drops	string	`json:"RPF_SUPPRESSED_DROPS"`
	Rp_destined_drops	string	`json:"RP_DESTINED_DROPS"`
	Discard_drops	string	`json:"DISCARD_DROPS"`
	Gre_lookup_drops	string	`json:"GRE_LOOKUP_DROPS"`
	Gre_processing_drops	string	`json:"GRE_PROCESSING_DROPS"`
	Lisp_punt_drops	string	`json:"LISP_PUNT_DROPS"`
	Lisp_encap_err_drops	string	`json:"LISP_ENCAP_ERR_DROPS"`
	Lisp_decap_err_drops	string	`json:"LISP_DECAP_ERR_DROPS"`
}

var ShowCefDropsLocation_Template = `Value LOCATION (\S+)
Value UNRESOLVED_DROPS (\d+)
Value UNSUPPORTED_DROPS (\d+)
Value NULL0_DROPS (\d+)
Value NO_ROUTE_DROPS (\d+)
Value NO_ADJACENCY_DROPS (\d+)
Value CHECKSUM_ERROR_DROPS (\d+)
Value RPF_DROPS (\d+)
Value RPF_SUPPRESSED_DROPS (\d+)
Value RP_DESTINED_DROPS (\d+)
Value DISCARD_DROPS (\d+)
Value GRE_LOOKUP_DROPS (\d+)
Value GRE_PROCESSING_DROPS (\d+)
Value LISP_PUNT_DROPS (\d+)
Value LISP_ENCAP_ERR_DROPS (\d+)
Value LISP_DECAP_ERR_DROPS (\d+)

Start
  ^Node: -> Continue.Record
  ^Node:\s+${LOCATION}
  ^\s+Unresolved\s+drops\s+packets\s+:\s+${UNRESOLVED_DROPS}
  ^\s+Unsupported\s+drops\s+packets\s+:\s+${UNSUPPORTED_DROPS}
  ^\s+Null0\s+drops\s+packets\s+:\s+${NULL0_DROPS}
  ^\s+No\s+route\s+drops\s+packets\s+:\s+${NO_ROUTE_DROPS}
  ^\s+No\s+Adjacency\s+drops\s+packets\s+:\s+${NO_ADJACENCY_DROPS}
  ^\s+Checksum\s+error\s+drops\s+packets\s+:\s+${CHECKSUM_ERROR_DROPS}
  ^\s+RPF\s+drops\s+packets\s+:\s+${RPF_DROPS}
  ^\s+RPF\s+suppressed\s+drops\s+packets\s+:\s+${RPF_SUPPRESSED_DROPS}
  ^\s+RP\s+destined\s+drops\s+packets\s+:\s+${RP_DESTINED_DROPS}
  ^\s+Discard\s+drops\s+packets\s+:\s+${DISCARD_DROPS}
  ^\s+GRE\s+lookup\s+drops\s+packets\s+:\s+${GRE_LOOKUP_DROPS}
  ^\s+GRE\s+processing\s+drops\s+packets\s+:\s+${GRE_PROCESSING_DROPS}
  ^\s+LISP\s+punt\s+drops\s+packets\s+:\s+${LISP_PUNT_DROPS}
  ^\s+LISP\s+encap\s+err\s+drops\s+packets\s+:\s+${LISP_ENCAP_ERR_DROPS}
  ^\s+LISP\s+decap\s+err\s+drops\s+packets\s+:\s+${LISP_DECAP_ERR_DROPS}
  ^CEF\s+Drop\s+Statistics\s*$$
  ^\s*$$
  ^.* -> Error "LINE NOT FOUND"
`