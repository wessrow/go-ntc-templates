package models

type PaloaltoPanosShowHighAvailabilityAll struct {
	Group	string	`json:"GROUP"`
	Mode	string	`json:"MODE"`
	State	string	`json:"STATE"`
	Ha1_link_state	string	`json:"HA1_LINK_STATE"`
	Ha1_link_setting	string	`json:"HA1_LINK_SETTING"`
	Ha2_link_state	string	`json:"HA2_LINK_STATE"`
	Ha2_link_setting	string	`json:"HA2_LINK_SETTING"`
	Priority	string	`json:"PRIORITY"`
	Preemptive	string	`json:"PREEMPTIVE"`
	Sw_version	string	`json:"SW_VERSION"`
	App_content_compatibility	string	`json:"APP_CONTENT_COMPATIBILITY"`
	Anti_virus_compatibility	string	`json:"ANTI_VIRUS_COMPATIBILITY"`
	Threat_content_compatibility	string	`json:"THREAT_CONTENT_COMPATIBILITY"`
	Vpn_client_sw_compatibility	string	`json:"VPN_CLIENT_SW_COMPATIBILITY"`
	Global_protect_client_sw_compatibility	string	`json:"GLOBAL_PROTECT_CLIENT_SW_COMPATIBILITY"`
	State_sync	string	`json:"STATE_SYNC"`
	Link_monitoring_enabled	string	`json:"LINK_MONITORING_ENABLED"`
	Path_monitoring_enabled	string	`json:"PATH_MONITORING_ENABLED"`
	Config_sync	string	`json:"CONFIG_SYNC"`
	Running_configuration	string	`json:"RUNNING_CONFIGURATION"`
}

var PaloaltoPanosShowHighAvailabilityAll_Template = `Value Filldown GROUP (\d+)
Value Required MODE (\S+)
Value STATE (\S+)
Value HA1_LINK_STATE (\S+)
Value HA1_LINK_SETTING (\S+)
Value HA2_LINK_STATE (\S+)
Value HA2_LINK_SETTING (\S+)
Value PRIORITY (\d+)
Value PREEMPTIVE (\w+)
Value SW_VERSION (\S+)
Value APP_CONTENT_COMPATIBILITY (\w+)
Value ANTI_VIRUS_COMPATIBILITY (\w+)
Value THREAT_CONTENT_COMPATIBILITY (\w+)
Value VPN_CLIENT_SW_COMPATIBILITY (\w+)
Value GLOBAL_PROTECT_CLIENT_SW_COMPATIBILITY (\w+)
Value STATE_SYNC (\S+)
Value LINK_MONITORING_ENABLED (\w+)
Value PATH_MONITORING_ENABLED (\w+)
Value CONFIG_SYNC (\w+)
Value RUNNING_CONFIGURATION (\w+)

Start
  ^Group\s+${GROUP}:
  ^\s+Mode:\s+${MODE}
  ^\s+State:\s+${STATE}
  ^\s+Peer\s+Information: -> Continue
  ^\s+HA1\sControl\sLinks\sJoint\sConfiguration: -> HA1
  ^\s+HA2.* -> HA2
  ^\s+Priority:\s+${PRIORITY}
  ^\s+Preemptive:\s+${PREEMPTIVE}
  ^\s+Software\sVersion:\s+${SW_VERSION}
  ^\s+Application\sContent\sCompatibility:\s+${APP_CONTENT_COMPATIBILITY}
  ^\s+Anti-Virus\sCompatibility:\s+${ANTI_VIRUS_COMPATIBILITY}
  ^\s+Threat\sContent\sCompatibility:\s+${THREAT_CONTENT_COMPATIBILITY}
  ^\s+VPN\sClient\sSoftware\sCompatibility:\s+${VPN_CLIENT_SW_COMPATIBILITY}
  ^\s+Global\sProtect\sClient\sSoftware\sCompatibility:\s+${GLOBAL_PROTECT_CLIENT_SW_COMPATIBILITY}
  ^\s+State\sSynchronization:\s${STATE_SYNC};
  ^\s+Peer\sInformation: -> Peer
  ^\s+Link\sMonitoring\sInformation: -> Link
  ^\s+Path\s+Monitoring\s+Information: -> Path
  ^\s+Configuration\s+Synchronization: -> Sync

HA1
  ^\s+Link\sState:\s${HA1_LINK_STATE};\s+Setting:\s+${HA1_LINK_SETTING} -> Start

HA2
  ^\s+Link\s+State:\s+${HA2_LINK_STATE};\s+Setting:\s+${HA2_LINK_SETTING} -> Start

Peer
  ^\s+Initial Monitor.* -> Start

Link
  ^\s+Enabled:\s+${LINK_MONITORING_ENABLED} -> Start

Path
  ^\s+Enabled:\s+${PATH_MONITORING_ENABLED} -> Start

Sync
  ^\s+Enabled:\s+${CONFIG_SYNC} 
  ^\s+Running\sConfiguration:\s+${RUNNING_CONFIGURATION} -> Record

`