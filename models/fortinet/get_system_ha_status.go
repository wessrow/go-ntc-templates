package fortinet

type GetSystemHaStatus struct {
	Ha_health                  []string `json:"HA_HEALTH"`
	Ha_session_pickup_status   string   `json:"HA_SESSION_PICKUP_STATUS"`
	Ha_slave_unit_serial       string   `json:"HA_SLAVE_UNIT_SERIAL"`
	Cluster_uptime             string   `json:"CLUSTER_UPTIME"`
	Ha_override_status         string   `json:"HA_OVERRIDE_STATUS"`
	Ha_master_unit_name        string   `json:"HA_MASTER_UNIT_NAME"`
	Ha_master_unit_serial      string   `json:"HA_MASTER_UNIT_SERIAL"`
	Ha_master_unit_index       string   `json:"HA_MASTER_UNIT_INDEX"`
	Ha_slave_unit_index        string   `json:"HA_SLAVE_UNIT_INDEX"`
	Model                      string   `json:"MODEL"`
	Ha_group                   string   `json:"HA_GROUP"`
	Cluster_state_changed_time string   `json:"CLUSTER_STATE_CHANGED_TIME"`
	Ha_session_pickup_delay    string   `json:"HA_SESSION_PICKUP_DELAY"`
	Ha_slave_unit_name         string   `json:"HA_SLAVE_UNIT_NAME"`
	Ha_mode                    string   `json:"HA_MODE"`
}

var GetSystemHaStatus_Template string = `#
# FG Version: 5.6, 6.0, 6.2, 6.4, 7.0
# HW        : varied
#
Value List HA_HEALTH (\S+(?:.*))
Value MODEL (\S+)
Value HA_MODE ([\S\s]+)
Value HA_GROUP (\S+)
Value CLUSTER_UPTIME ([\S\s]+)
Value CLUSTER_STATE_CHANGED_TIME ([\S\s]+)
Value HA_SESSION_PICKUP_STATUS (\S+)
Value HA_SESSION_PICKUP_DELAY (\S+)
Value HA_OVERRIDE_STATUS (\S+)
Value HA_MASTER_UNIT_NAME (\S+)
Value HA_SLAVE_UNIT_NAME (\S+)
Value HA_MASTER_UNIT_SERIAL (\S+)
Value HA_SLAVE_UNIT_SERIAL (\S+)
Value HA_MASTER_UNIT_INDEX (\S+)
Value HA_SLAVE_UNIT_INDEX (\S+)

Start
  ^HA\s+Health\s+Status:\s+${HA_HEALTH}
  ^HA\s+Health\s+Status: -> UnhealthyStatus
  ^Model:\s+${MODEL}
  ^Mode:\s+${HA_MODE}
  ^Group:\s+${HA_GROUP}
  ^Debug:\s+\d+
  ^Cluster\s+Uptime:\s+${CLUSTER_UPTIME}
  ^Cluster\s+state\s+change\s+time:\s+${CLUSTER_STATE_CHANGED_TIME}
  ^(Master|Primary)\s+selected\s+using:
  ^\s*\<\S+
  ^ses_pickup:\s+${HA_SESSION_PICKUP_STATUS},\s+ses_pickup_delay=${HA_SESSION_PICKUP_DELAY}
  ^override:\s+${HA_OVERRIDE_STATUS}
  ^Configuration\s+Status: -> Configuration_Status
  # Catch old 6.0_noha with no "Configuraton Status"
  ^System\s+Usage\s+stats: ->  System_Usage_stats
  ^. -> Error "in-Start"

UnhealthyStatus
  # semicolon necessary to anchor
  ^\s+${HA_HEALTH};\s*$$
  ^Model:\s+${MODEL} -> Start

Configuration_Status
  ^System\s+Usage\s+stats: ->  System_Usage_stats
  ^\s*\S+\([\S\s]+\):\s\S+$$
  ^. -> Error "in-Configuration_Status"

System_Usage_stats
  ^HBDEV\s+stats: -> HBDEV_MONDEV_stats
  ^\s*\S+\([\S\s]+\):$$
  #^\s*\S+:\s+
  ^\s*sessions=
  ^. -> Error "in-System_Usage_stats"

HBDEV_MONDEV_stats
  # Combine stats, no MONDEV in older FW's
  ^\s*\S+\([\S\s]+\):$$
  ^\s*\S+:\s.+rx.+tx.+$$
  ^MONDEV\s+stats:
  ^(Master|Primary)\s*:\s+${HA_MASTER_UNIT_NAME}\s*,\s+${HA_MASTER_UNIT_SERIAL},\s+(HA\s+cluster\s+index|cluster\s+index)\s+=\s+${HA_MASTER_UNIT_INDEX}
  ^(Slave|Secondary)\s*:\s+${HA_SLAVE_UNIT_NAME}\s*,\s+${HA_SLAVE_UNIT_SERIAL},\s+(|HA)\s*cluster\s+index\s+=\s+${HA_SLAVE_UNIT_INDEX}
  ^number\s+of\s+vcluster:\s+\d+
  ^vcluster\s+\d+:
  ^(Master|Slave|Primary|Secondary)\s*:\s+\S+,\s+(operating\s+cluster\s+index|HA\s+operating\s+index)\s+=\s+\d+ -> Record
  ^\s*$$
  ^. -> Error "in-HBDEV_MONDEV_stats"
`
