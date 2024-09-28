package models

type CiscoIosShowRunningConfigPartitionRouteMap struct {
	Route_map_name	string	`json:"ROUTE_MAP_NAME"`
	Sequence	string	`json:"SEQUENCE"`
	Match_metric_value	string	`json:"MATCH_METRIC_VALUE"`
	Match_prefix_list	string	`json:"MATCH_PREFIX_LIST"`
	Match_community_value	string	`json:"MATCH_COMMUNITY_VALUE"`
	Set_community_value	string	`json:"SET_COMMUNITY_VALUE"`
	Set_local_pref_value	string	`json:"SET_LOCAL_PREF_VALUE"`
	Set_prepend_as_path	string	`json:"SET_PREPEND_AS_PATH"`
}