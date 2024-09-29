package models

type CiscoXrAdminShowControllerFabricHealth struct {
	Fsdb_aggregator	string	`json:"FSDB_AGGREGATOR"`
	Rack_id	string	`json:"RACK_ID"`
	Fsdb_status	string	`json:"FSDB_STATUS"`
	Sfe_status_fc0	string	`json:"SFE_STATUS_FC0"`
	Sfe_status_fc1	string	`json:"SFE_STATUS_FC1"`
	Sfe_status_fc2	string	`json:"SFE_STATUS_FC2"`
	Sfe_status_fc3	string	`json:"SFE_STATUS_FC3"`
	Sfe_status_fc4	string	`json:"SFE_STATUS_FC4"`
	Sfe_status_fc5	string	`json:"SFE_STATUS_FC5"`
	Router_rack_health_total	string	`json:"ROUTER_RACK_HEALTH_TOTAL"`
	Router_rack_health_lcc	string	`json:"ROUTER_RACK_HEALTH_LCC"`
	Router_rack_health_fcc	string	`json:"ROUTER_RACK_HEALTH_FCC"`
	Router_planes_health_up	string	`json:"ROUTER_PLANES_HEALTH_UP"`
	Router_planes_health_mcast_down	string	`json:"ROUTER_PLANES_HEALTH_MCAST_DOWN"`
	Router_planes_health_down	string	`json:"ROUTER_PLANES_HEALTH_DOWN"`
	Router_planes_health_admin_down	string	`json:"ROUTER_PLANES_HEALTH_ADMIN_DOWN"`
	Router_sfe_asics_health_total	string	`json:"ROUTER_SFE_ASICS_HEALTH_TOTAL"`
	Router_sfe_asics_health_up	string	`json:"ROUTER_SFE_ASICS_HEALTH_UP"`
	Router_sfe_asics_health_down	string	`json:"ROUTER_SFE_ASICS_HEALTH_DOWN"`
	Router_fia_asics_health_total	string	`json:"ROUTER_FIA_ASICS_HEALTH_TOTAL"`
	Router_fia_asics_health_up	string	`json:"ROUTER_FIA_ASICS_HEALTH_UP"`
	Router_fia_asics_health_down	string	`json:"ROUTER_FIA_ASICS_HEALTH_DOWN"`
	Plane_id_0_admin_state	string	`json:"PLANE_ID_0_ADMIN_STATE"`
	Plane_id_1_admin_state	string	`json:"PLANE_ID_1_ADMIN_STATE"`
	Plane_id_2_admin_state	string	`json:"PLANE_ID_2_ADMIN_STATE"`
	Plane_id_3_admin_state	string	`json:"PLANE_ID_3_ADMIN_STATE"`
	Plane_id_4_admin_state	string	`json:"PLANE_ID_4_ADMIN_STATE"`
	Plane_id_5_admin_state	string	`json:"PLANE_ID_5_ADMIN_STATE"`
	Plane_id_0_plane_state	string	`json:"PLANE_ID_0_PLANE_STATE"`
	Plane_id_1_plane_state	string	`json:"PLANE_ID_1_PLANE_STATE"`
	Plane_id_2_plane_state	string	`json:"PLANE_ID_2_PLANE_STATE"`
	Plane_id_3_plane_state	string	`json:"PLANE_ID_3_PLANE_STATE"`
	Plane_id_4_plane_state	string	`json:"PLANE_ID_4_PLANE_STATE"`
	Plane_id_5_plane_state	string	`json:"PLANE_ID_5_PLANE_STATE"`
	Plane_id_0_racks_in_issue	string	`json:"PLANE_ID_0_RACKS_IN_ISSUE"`
	Plane_id_1_racks_in_issue	string	`json:"PLANE_ID_1_RACKS_IN_ISSUE"`
	Plane_id_2_racks_in_issue	string	`json:"PLANE_ID_2_RACKS_IN_ISSUE"`
	Plane_id_3_racks_in_issue	string	`json:"PLANE_ID_3_RACKS_IN_ISSUE"`
	Plane_id_4_racks_in_issue	string	`json:"PLANE_ID_4_RACKS_IN_ISSUE"`
	Plane_id_5_racks_in_issue	string	`json:"PLANE_ID_5_RACKS_IN_ISSUE"`
	Plane_id_0_data_drop_err	string	`json:"PLANE_ID_0_DATA_DROP_ERR"`
	Plane_id_1_data_drop_err	string	`json:"PLANE_ID_1_DATA_DROP_ERR"`
	Plane_id_2_data_drop_err	string	`json:"PLANE_ID_2_DATA_DROP_ERR"`
	Plane_id_3_data_drop_err	string	`json:"PLANE_ID_3_DATA_DROP_ERR"`
	Plane_id_4_data_drop_err	string	`json:"PLANE_ID_4_DATA_DROP_ERR"`
	Plane_id_5_data_drop_err	string	`json:"PLANE_ID_5_DATA_DROP_ERR"`
	Rack_planes_health_up	string	`json:"RACK_PLANES_HEALTH_UP"`
	Rack_planes_health_mcast_down	string	`json:"RACK_PLANES_HEALTH_MCAST_DOWN"`
	Rack_planes_health_down	string	`json:"RACK_PLANES_HEALTH_DOWN"`
	Rack_sfe_asics_health_total	string	`json:"RACK_SFE_ASICS_HEALTH_TOTAL"`
	Rack_sfe_asics_health_up	string	`json:"RACK_SFE_ASICS_HEALTH_UP"`
	Rack_sfe_asics_health_down	string	`json:"RACK_SFE_ASICS_HEALTH_DOWN"`
	Rack_fia_asics_health_total	string	`json:"RACK_FIA_ASICS_HEALTH_TOTAL"`
	Rack_fia_asics_health_up	string	`json:"RACK_FIA_ASICS_HEALTH_UP"`
	Rack_fia_asics_health_down	string	`json:"RACK_FIA_ASICS_HEALTH_DOWN"`
	Rack_valid_fab_ids	string	`json:"RACK_VALID_FAB_IDS"`
	Plane_id_0_sfe_asics_total	string	`json:"PLANE_ID_0_SFE_ASICS_TOTAL"`
	Plane_id_0_sfe_asics_up	string	`json:"PLANE_ID_0_SFE_ASICS_UP"`
	Plane_id_0_sfe_asics_down	string	`json:"PLANE_ID_0_SFE_ASICS_DOWN"`
	Plane_id_0_fab_id_reachable	string	`json:"PLANE_ID_0_FAB_ID_REACHABLE"`
	Plane_id_1_sfe_asics_total	string	`json:"PLANE_ID_1_SFE_ASICS_TOTAL"`
	Plane_id_1_sfe_asics_up	string	`json:"PLANE_ID_1_SFE_ASICS_UP"`
	Plane_id_1_sfe_asics_down	string	`json:"PLANE_ID_1_SFE_ASICS_DOWN"`
	Plane_id_1_fab_id_reachable	string	`json:"PLANE_ID_1_FAB_ID_REACHABLE"`
	Plane_id_2_sfe_asics_total	string	`json:"PLANE_ID_2_SFE_ASICS_TOTAL"`
	Plane_id_2_sfe_asics_up	string	`json:"PLANE_ID_2_SFE_ASICS_UP"`
	Plane_id_2_sfe_asics_down	string	`json:"PLANE_ID_2_SFE_ASICS_DOWN"`
	Plane_id_2_fab_id_reachable	string	`json:"PLANE_ID_2_FAB_ID_REACHABLE"`
	Plane_id_3_sfe_asics_total	string	`json:"PLANE_ID_3_SFE_ASICS_TOTAL"`
	Plane_id_3_sfe_asics_up	string	`json:"PLANE_ID_3_SFE_ASICS_UP"`
	Plane_id_3_sfe_asics_down	string	`json:"PLANE_ID_3_SFE_ASICS_DOWN"`
	Plane_id_3_fab_id_reachable	string	`json:"PLANE_ID_3_FAB_ID_REACHABLE"`
	Plane_id_4_sfe_asics_total	string	`json:"PLANE_ID_4_SFE_ASICS_TOTAL"`
	Plane_id_4_sfe_asics_up	string	`json:"PLANE_ID_4_SFE_ASICS_UP"`
	Plane_id_4_sfe_asics_down	string	`json:"PLANE_ID_4_SFE_ASICS_DOWN"`
	Plane_id_4_fab_id_reachable	string	`json:"PLANE_ID_4_FAB_ID_REACHABLE"`
	Plane_id_5_sfe_asics_total	string	`json:"PLANE_ID_5_SFE_ASICS_TOTAL"`
	Plane_id_5_sfe_asics_up	string	`json:"PLANE_ID_5_SFE_ASICS_UP"`
	Plane_id_5_sfe_asics_down	string	`json:"PLANE_ID_5_SFE_ASICS_DOWN"`
	Plane_id_5_fab_id_reachable	string	`json:"PLANE_ID_5_FAB_ID_REACHABLE"`
}

var CiscoXrAdminShowControllerFabricHealth_Template = `Value FSDB_AGGREGATOR (\w+)
Value RACK_ID (\d+)
Value FSDB_STATUS (\S+)
Value SFE_STATUS_FC0 (\S+)
Value SFE_STATUS_FC1 (\S+)
Value SFE_STATUS_FC2 (\S+)
Value SFE_STATUS_FC3 (\S+)
Value SFE_STATUS_FC4 (\S+)
Value SFE_STATUS_FC5 (\S+)
Value ROUTER_RACK_HEALTH_TOTAL (\d+)
Value ROUTER_RACK_HEALTH_LCC (\d+)
Value ROUTER_RACK_HEALTH_FCC (\d+)
Value ROUTER_PLANES_HEALTH_UP (\d+)
Value ROUTER_PLANES_HEALTH_MCAST_DOWN (\d+)
Value ROUTER_PLANES_HEALTH_DOWN (\d+)
Value ROUTER_PLANES_HEALTH_ADMIN_DOWN (\d+)
Value ROUTER_SFE_ASICS_HEALTH_TOTAL (\d+)
Value ROUTER_SFE_ASICS_HEALTH_UP (\d+)
Value ROUTER_SFE_ASICS_HEALTH_DOWN (\d+)
Value ROUTER_FIA_ASICS_HEALTH_TOTAL (\d+)
Value ROUTER_FIA_ASICS_HEALTH_UP (\d+)
Value ROUTER_FIA_ASICS_HEALTH_DOWN (\d+)
Value PLANE_ID_0_ADMIN_STATE (\S+)
Value PLANE_ID_1_ADMIN_STATE (\S+)
Value PLANE_ID_2_ADMIN_STATE (\S+)
Value PLANE_ID_3_ADMIN_STATE (\S+)
Value PLANE_ID_4_ADMIN_STATE (\S+)
Value PLANE_ID_5_ADMIN_STATE (\S+)
Value PLANE_ID_0_PLANE_STATE (\S+)
Value PLANE_ID_1_PLANE_STATE (\S+)
Value PLANE_ID_2_PLANE_STATE (\S+)
Value PLANE_ID_3_PLANE_STATE (\S+)
Value PLANE_ID_4_PLANE_STATE (\S+)
Value PLANE_ID_5_PLANE_STATE (\S+)
Value PLANE_ID_0_RACKS_IN_ISSUE (\d+)
Value PLANE_ID_1_RACKS_IN_ISSUE (\d+)
Value PLANE_ID_2_RACKS_IN_ISSUE (\d+)
Value PLANE_ID_3_RACKS_IN_ISSUE (\d+)
Value PLANE_ID_4_RACKS_IN_ISSUE (\d+)
Value PLANE_ID_5_RACKS_IN_ISSUE (\d+)
Value PLANE_ID_0_DATA_DROP_ERR (\S+)
Value PLANE_ID_1_DATA_DROP_ERR (\S+)
Value PLANE_ID_2_DATA_DROP_ERR (\S+)
Value PLANE_ID_3_DATA_DROP_ERR (\S+)
Value PLANE_ID_4_DATA_DROP_ERR (\S+)
Value PLANE_ID_5_DATA_DROP_ERR (\S+)
Value RACK_PLANES_HEALTH_UP (\d+)
Value RACK_PLANES_HEALTH_MCAST_DOWN (\d+)
Value RACK_PLANES_HEALTH_DOWN (\d+)
Value RACK_SFE_ASICS_HEALTH_TOTAL (\d+)
Value RACK_SFE_ASICS_HEALTH_UP (\d+)
Value RACK_SFE_ASICS_HEALTH_DOWN (\d+)
Value RACK_FIA_ASICS_HEALTH_TOTAL (\d+)
Value RACK_FIA_ASICS_HEALTH_UP (\d+)
Value RACK_FIA_ASICS_HEALTH_DOWN (\d+)
Value RACK_VALID_FAB_IDS (\d+)
Value PLANE_ID_0_SFE_ASICS_TOTAL (\S+)
Value PLANE_ID_0_SFE_ASICS_UP (\S+)
Value PLANE_ID_0_SFE_ASICS_DOWN (\S+)
Value PLANE_ID_0_FAB_ID_REACHABLE (\S+)
Value PLANE_ID_1_SFE_ASICS_TOTAL (\S+)
Value PLANE_ID_1_SFE_ASICS_UP (\S+)
Value PLANE_ID_1_SFE_ASICS_DOWN (\S+)
Value PLANE_ID_1_FAB_ID_REACHABLE (\S+)
Value PLANE_ID_2_SFE_ASICS_TOTAL (\S+)
Value PLANE_ID_2_SFE_ASICS_UP (\S+)
Value PLANE_ID_2_SFE_ASICS_DOWN (\S+)
Value PLANE_ID_2_FAB_ID_REACHABLE (\S+)
Value PLANE_ID_3_SFE_ASICS_TOTAL (\S+)
Value PLANE_ID_3_SFE_ASICS_UP (\S+)
Value PLANE_ID_3_SFE_ASICS_DOWN (\S+)
Value PLANE_ID_3_FAB_ID_REACHABLE (\S+)
Value PLANE_ID_4_SFE_ASICS_TOTAL (\S+)
Value PLANE_ID_4_SFE_ASICS_UP (\S+)
Value PLANE_ID_4_SFE_ASICS_DOWN (\S+)
Value PLANE_ID_4_FAB_ID_REACHABLE (\S+)
Value PLANE_ID_5_SFE_ASICS_TOTAL (\S+)
Value PLANE_ID_5_SFE_ASICS_UP (\S+)
Value PLANE_ID_5_SFE_ASICS_DOWN (\S+)
Value PLANE_ID_5_FAB_ID_REACHABLE (\S+)


Start
  ^.+UTC
  ^Fabric.+Health
  ^-.+
  ^Flags.+Down\s?
  ^\s+L.+Yes\s+?
  ^\s+F.+Ok
  ^\s+V.+Valid,\s+? -> Collaborator
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
  
Collaborator
  ^Collaborator.+
  ^-.+ 
  ^\s+FSDB Aggregator:\s+${FSDB_AGGREGATOR}
  ^\s+\+.+\+?
  ^\s+\|Rack id\s+\|\s?${RACK_ID}\|
  ^\s+\|FSDB status\|${FSDB_STATUS}\|
  ^\s+\|FC Location.+
  ^\s+\|SFE status\s+\|\s+${SFE_STATUS_FC0}\s+\|\s+${SFE_STATUS_FC1}\s+\|\s+${SFE_STATUS_FC2}\s+\|\s+${SFE_STATUS_FC3}\s+\|\s+${SFE_STATUS_FC4}\s+\|\s+${SFE_STATUS_FC5}\s+\| -> RouterHealth 
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"

RouterHealth
  ^\s+\+.+\+?
  ^Router\s+Health:
  ^-.+
  ^\s+Rack.+Asics\s+?
  ^\s+T.+\s+?
  ^\s+-.+
  ^\s+${ROUTER_RACK_HEALTH_TOTAL}/${ROUTER_RACK_HEALTH_LCC}/${ROUTER_RACK_HEALTH_FCC}\s+${ROUTER_PLANES_HEALTH_UP}/${ROUTER_PLANES_HEALTH_MCAST_DOWN}/${ROUTER_PLANES_HEALTH_DOWN}/${ROUTER_PLANES_HEALTH_ADMIN_DOWN}\s+${ROUTER_SFE_ASICS_HEALTH_TOTAL}/${ROUTER_SFE_ASICS_HEALTH_UP}/${ROUTER_SFE_ASICS_HEALTH_DOWN}\s+${ROUTER_FIA_ASICS_HEALTH_TOTAL}/${ROUTER_FIA_ASICS_HEALTH_UP}/${ROUTER_FIA_ASICS_HEALTH_DOWN}
  ^\s+Plane.+\s+?
  ^\s+id.+
  ^\s+0\s+${PLANE_ID_0_ADMIN_STATE}\s+${PLANE_ID_0_PLANE_STATE}\s+${PLANE_ID_0_RACKS_IN_ISSUE}\s+${PLANE_ID_0_DATA_DROP_ERR}\s+?
  ^\s+1\s+${PLANE_ID_1_ADMIN_STATE}\s+${PLANE_ID_1_PLANE_STATE}\s+${PLANE_ID_1_RACKS_IN_ISSUE}\s+${PLANE_ID_1_DATA_DROP_ERR}\s+?
  ^\s+2\s+${PLANE_ID_2_ADMIN_STATE}\s+${PLANE_ID_2_PLANE_STATE}\s+${PLANE_ID_2_RACKS_IN_ISSUE}\s+${PLANE_ID_2_DATA_DROP_ERR}\s+?
  ^\s+3\s+${PLANE_ID_3_ADMIN_STATE}\s+${PLANE_ID_3_PLANE_STATE}\s+${PLANE_ID_3_RACKS_IN_ISSUE}\s+${PLANE_ID_3_DATA_DROP_ERR}\s+?
  ^\s+4\s+${PLANE_ID_4_ADMIN_STATE}\s+${PLANE_ID_4_PLANE_STATE}\s+${PLANE_ID_4_RACKS_IN_ISSUE}\s+${PLANE_ID_4_DATA_DROP_ERR}\s+?
  ^\s+5\s+${PLANE_ID_5_ADMIN_STATE}\s+${PLANE_ID_5_PLANE_STATE}\s+${PLANE_ID_5_RACKS_IN_ISSUE}\s+${PLANE_ID_5_DATA_DROP_ERR}\s+? -> RackHealth
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
  
RackHealth
  ^\s+-.+
  ^Rack\s+Health:
  ^-.+ 
  ^\s+Rack.+LCC
  ^\s+SFE.+Valid\s+?
  ^\s+T.+
  ^\s+${RACK_SFE_ASICS_HEALTH_TOTAL}/${RACK_SFE_ASICS_HEALTH_UP}/${RACK_SFE_ASICS_HEALTH_DOWN}\s+${RACK_FIA_ASICS_HEALTH_TOTAL}/${RACK_FIA_ASICS_HEALTH_UP}/${RACK_FIA_ASICS_HEALTH_DOWN}\s+${RACK_PLANES_HEALTH_UP}/${RACK_PLANES_HEALTH_MCAST_DOWN}/${RACK_PLANES_HEALTH_DOWN}\s+${RACK_VALID_FAB_IDS}
  ^\s+Plane.+ids\s+?
  ^\s+id.+
  ^\s+0\s+\S+\s+${PLANE_ID_0_SFE_ASICS_TOTAL}/${PLANE_ID_0_SFE_ASICS_UP}/${PLANE_ID_0_SFE_ASICS_DOWN}\s+${PLANE_ID_0_FAB_ID_REACHABLE}\s+?
  ^\s+1\s+\S+\s+${PLANE_ID_1_SFE_ASICS_TOTAL}/${PLANE_ID_1_SFE_ASICS_UP}/${PLANE_ID_1_SFE_ASICS_DOWN}\s+${PLANE_ID_1_FAB_ID_REACHABLE}\s+?
  ^\s+2\s+\S+\s+${PLANE_ID_2_SFE_ASICS_TOTAL}/${PLANE_ID_2_SFE_ASICS_UP}/${PLANE_ID_2_SFE_ASICS_DOWN}\s+${PLANE_ID_2_FAB_ID_REACHABLE}\s+?
  ^\s+3\s+\S+\s+${PLANE_ID_3_SFE_ASICS_TOTAL}/${PLANE_ID_3_SFE_ASICS_UP}/${PLANE_ID_3_SFE_ASICS_DOWN}\s+${PLANE_ID_3_FAB_ID_REACHABLE}\s+?
  ^\s+4\s+\S+\s+${PLANE_ID_4_SFE_ASICS_TOTAL}/${PLANE_ID_4_SFE_ASICS_UP}/${PLANE_ID_4_SFE_ASICS_DOWN}\s+${PLANE_ID_4_FAB_ID_REACHABLE}\s+?
  ^\s+5\s+\S+\s+${PLANE_ID_5_SFE_ASICS_TOTAL}/${PLANE_ID_5_SFE_ASICS_UP}/${PLANE_ID_5_SFE_ASICS_DOWN}\s+${PLANE_ID_5_FAB_ID_REACHABLE}\s+? -> Record
  ^\s+-.+
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
  
EOF

`