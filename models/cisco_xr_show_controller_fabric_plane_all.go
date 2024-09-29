package models

type CiscoXrShowControllerFabricPlaneAll struct {
	Plane_id	string	`json:"PLANE_ID"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Plane_state	string	`json:"PLANE_STATE"`
	Up_dn_counter	string	`json:"UP_DN_COUNTER"`
	Up_mcast_counter	string	`json:"UP_MCAST_COUNTER"`
}

var CiscoXrShowControllerFabricPlaneAll_Template = `Value PLANE_ID (\d+)
Value ADMIN_STATE (\S+)
Value PLANE_STATE (\S+)
Value UP_DN_COUNTER (\d+)
Value UP_MCAST_COUNTER (\d+)


Start
  ^.+UTC
  ^Plane.+up->mcast
  ^Id.+counter\s+counter
  ^-.+ 
  ^${PLANE_ID}\s+${ADMIN_STATE}\s+${PLANE_STATE}\s+${UP_DN_COUNTER}\s+${UP_MCAST_COUNTER} -> Record
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
  
EOF
`