package cisco_ios 

type ShowRunningConfigPartitionRouteMap struct {
	Route_map_name	string	`json:"ROUTE_MAP_NAME"`
	Sequence	string	`json:"SEQUENCE"`
	Match_metric_value	string	`json:"MATCH_METRIC_VALUE"`
	Match_prefix_list	string	`json:"MATCH_PREFIX_LIST"`
	Match_community_value	string	`json:"MATCH_COMMUNITY_VALUE"`
	Set_community_value	string	`json:"SET_COMMUNITY_VALUE"`
	Set_local_pref_value	string	`json:"SET_LOCAL_PREF_VALUE"`
	Set_prepend_as_path	string	`json:"SET_PREPEND_AS_PATH"`
}

var ShowRunningConfigPartitionRouteMap_Template = `Value Required ROUTE_MAP_NAME (\w+)
Value SEQUENCE (\d+)
Value MATCH_METRIC_VALUE (\d+)
Value MATCH_PREFIX_LIST (\w+)
Value MATCH_COMMUNITY_VALUE (\d+.*\d+|\w+.*\w+|\d+)
Value SET_COMMUNITY_VALUE (\d+.*\d+.*\w+|\w+.*\w+)
Value SET_LOCAL_PREF_VALUE (\d+)
Value SET_PREPEND_AS_PATH (\w+.*\w+.*\w+|\w+)

Start
  ^route-map -> Continue.Record
  ^route-map\s+${ROUTE_MAP_NAME}\s+permit\s+${SEQUENCE}
  ^\smatch\smetric\s${MATCH_METRIC_VALUE}
  ^\sset\s+community\s+${SET_COMMUNITY_VALUE}
  ^\smatch\s+ip\s+address\s+prefix-list\s+${MATCH_PREFIX_LIST}
  ^\sset\s+local-preference\s+${SET_LOCAL_PREF_VALUE}
  ^\smatch\s+community\s+${MATCH_COMMUNITY_VALUE}
  ^\sset\s+as-path\s+prepend\s+${SET_PREPEND_AS_PATH}
  ^!
  # Catch all unuseful raw data
  ^(!\s*|$$|Building configuration.*|Current configuration.*|Configuration.*|end.*)
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  # Error out if raw data does not match any above rules.
  ^.* -> Error "Could not parse line:"

 `