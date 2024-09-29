package juniper_junos 

type ShowChassisClusterStatus struct {
	Cluster_id	string	`json:"CLUSTER_ID"`
	Redundancy_group_id	string	`json:"REDUNDANCY_GROUP_ID"`
	Failover_count	string	`json:"FAILOVER_COUNT"`
	Node	[]string	`json:"NODE"`
	Priority	[]string	`json:"PRIORITY"`
	Status	[]string	`json:"STATUS"`
	Preempt	[]string	`json:"PREEMPT"`
	Manual	[]string	`json:"MANUAL"`
	Monitor_failures	[]string	`json:"MONITOR_FAILURES"`
}

var ShowChassisClusterStatus_Template = `Value Filldown CLUSTER_ID (\d+)
Value REDUNDANCY_GROUP_ID (\d+)
Value FAILOVER_COUNT (\d+)
Value List NODE (\S+)
Value List PRIORITY (\d+)
Value List STATUS (\S+)
Value List PREEMPT (\S+)
Value List MANUAL (\S+)
Value List MONITOR_FAILURES (\S+)

Start
  ^[Mm]onitor\s+[Ff]ailure\s+[Cc]odes:\s*$$
  ^\S{1,3}\s+
  ^\s*$$
  ^Cluster\s+ID:\s+${CLUSTER_ID}\s*$$
  ^Node\s+Priority\s+Status\s+Preempt\s+Manual\s+Monitor-failures\s*$$
  ^Redundancy\s+group:\s+${REDUNDANCY_GROUP_ID}\s*,\s*Failover\s+count:\s+${FAILOVER_COUNT}\s*$$ -> RedGroup
  ^. -> Error

RedGroup
  ^${NODE}\s+${PRIORITY}\s+${STATUS}\s+${PREEMPT}\s+${MANUAL}\s+${MONITOR_FAILURES}\s*$$
  ^Redundancy\s+group -> Continue.Record
  ^Redundancy\s+group:\s+${REDUNDANCY_GROUP_ID}\s*,\s*Failover\s+count:\s+${FAILOVER_COUNT}\s*$$
  ^. -> Error
`