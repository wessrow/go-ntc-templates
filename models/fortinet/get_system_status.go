package fortinet

type GetSystemStatus struct {
	Extreme_db                    string `json:"EXTREME_DB"`
	Ips_engine                    string `json:"IPS_ENGINE"`
	Botnet_db                     string `json:"BOTNET_DB"`
	Log_hard_disk                 string `json:"LOG_HARD_DISK"`
	Securitylevel                 string `json:"SECURITYLEVEL"`
	Signature                     string `json:"SIGNATURE"`
	Ot_detect_db                  string `json:"OT_DETECT_DB"`
	Virtual_domains_status        string `json:"VIRTUAL_DOMAINS_STATUS"`
	Fips_cc_mode                  string `json:"FIPS_CC_MODE"`
	Release_version_information   string `json:"RELEASE_VERSION_INFORMATION"`
	Serial_number                 string `json:"SERIAL_NUMBER"`
	Iot_detect                    string `json:"IOT_DETECT"`
	Operation_mode                string `json:"OPERATION_MODE"`
	Virtual_domain_configuration  string `json:"VIRTUAL_DOMAIN_CONFIGURATION"`
	Cluster_state_change_time     string `json:"CLUSTER_STATE_CHANGE_TIME"`
	Ips_malicious_url_database    string `json:"IPS_MALICIOUS_URL_DATABASE"`
	Ot_patch_db                   string `json:"OT_PATCH_DB"`
	System_part_number            string `json:"SYSTEM_PART_NUMBER"`
	Hostname                      string `json:"HOSTNAME"`
	Bios_version                  string `json:"BIOS_VERSION"`
	Private_encryption            string `json:"PRIVATE_ENCRYPTION"`
	Current_virtual_domain        string `json:"CURRENT_VIRTUAL_DOMAIN"`
	Cluster_uptime                string `json:"CLUSTER_UPTIME"`
	Av_ai_ml_model                string `json:"AV_AI_ML_MODEL"`
	Version                       string `json:"VERSION"`
	Virus_db                      string `json:"VIRUS_DB"`
	Fmwp_db                       string `json:"FMWP_DB"`
	Industrial_db                 string `json:"INDUSTRIAL_DB"`
	Ot_threat_db                  string `json:"OT_THREAT_DB"`
	Max_number_of_virtual_domains string `json:"MAX_NUMBER_OF_VIRTUAL_DOMAINS"`
	Current_ha_mode               string `json:"CURRENT_HA_MODE"`
	Fortios_x86_64                string `json:"FORTIOS_X86_64"`
	Last_reboot_reason            string `json:"LAST_REBOOT_REASON"`
	Ips_db                        string `json:"IPS_DB"`
	Ips_etdb                      string `json:"IPS_ETDB"`
	Branch_point                  string `json:"BRANCH_POINT"`
	System_time                   string `json:"SYSTEM_TIME"`
	Extended_db                   string `json:"EXTENDED_DB"`
	App_db                        string `json:"APP_DB"`
}

var GetSystemStatus_Template string = `#
# Patrick Marc Preuss, Refried Jello
#
# 2020-01-13: Inital Version
# 2021-01-08: Update for 6.2
#
# FG Version: 5.6, 6.0.7, 6.2
# HW        : FG1500D
# CLUSTER MODE: AP
# VDOMS       : ENABLED
#
# 2023-10-25: Update for 7.4 by Sebastian
# 2022-11-27: Update for 7.0 by Klaus
#
Value Required HOSTNAME (\S+)
Value VERSION (.*)
Value SECURITYLEVEL (.*)
Value SIGNATURE (.*)
Value VIRUS_DB (.*)
Value EXTENDED_DB (.*)
Value EXTREME_DB (.*)
Value IPS_DB (.*)
Value IPS_ETDB (.*)
Value APP_DB (.*)
Value FMWP_DB (.*) 
Value INDUSTRIAL_DB (.*)
Value SERIAL_NUMBER (\S+)
Value IPS_MALICIOUS_URL_DATABASE (.*)
Value IOT_DETECT (.*)
Value OT_DETECT_DB (.*)
Value OT_PATCH_DB (.*)
Value OT_THREAT_DB (.*)
Value IPS_ENGINE (.*)
Value BOTNET_DB (.*)
Value BIOS_VERSION (\S+)
Value SYSTEM_PART_NUMBER (\S+)
Value LOG_HARD_DISK (.*)
Value PRIVATE_ENCRYPTION (\S+)
Value OPERATION_MODE (\S+)
Value CURRENT_VIRTUAL_DOMAIN (\S+)
Value MAX_NUMBER_OF_VIRTUAL_DOMAINS (\d+)
Value VIRTUAL_DOMAINS_STATUS (.*)
Value VIRTUAL_DOMAIN_CONFIGURATION (\S+)
Value FIPS_CC_MODE (\S+)
Value CURRENT_HA_MODE (.*)
Value CLUSTER_UPTIME (.*)
Value CLUSTER_STATE_CHANGE_TIME (.*)
Value BRANCH_POINT (\S+)
Value RELEASE_VERSION_INFORMATION (\S+)
Value FORTIOS_X86_64 (\S+)
Value SYSTEM_TIME (.*)
Value AV_AI_ML_MODEL (.+)
Value LAST_REBOOT_REASON (.+)

Start
  ^Version:\s+${VERSION}
  ^Security Level:\s+${SECURITYLEVEL}
  ^Firmware Signature:\s+${SIGNATURE}
  ^Virus-DB:\s+${VIRUS_DB}
  ^Extended\s+DB:\s+${EXTENDED_DB}
  ^AV\s+AI/ML\s+Model:\s+${AV_AI_ML_MODEL}
  ^Extreme\s+DB:\s+${EXTREME_DB}
  ^IPS-DB:\s+${IPS_DB}
  ^IPS-ETDB:\s+${IPS_ETDB}
  ^APP-DB:\s+${APP_DB}
  ^FMWP-DB:\s+${APP_DB}
  ^INDUSTRIAL-DB:\s+${INDUSTRIAL_DB}
  ^Serial-Number:\s+${SERIAL_NUMBER}
  ^IPS\s+Malicious\s+URL\s+Database:\s+${IPS_MALICIOUS_URL_DATABASE}
  ^IoT-Detect:\s+${APP_DB}
  ^OT-Detect-DB:\s+${APP_DB}
  ^OT-Patch-DB:\s+${APP_DB}
  ^OT-Threat-DB:\s+${APP_DB}
  ^IPS-Engine:\s+${APP_DB}
  ^Botnet\s+DB:\s+${BOTNET_DB}
  ^BIOS\s+version:\s+${BIOS_VERSION}
  ^System\s+Part-Number:\s+${SYSTEM_PART_NUMBER}
  ^Log\s+hard\s+disk:\s+${LOG_HARD_DISK}
  ^Hostname:\s+${HOSTNAME}
  ^Private\s+Encryption:\s+${PRIVATE_ENCRYPTION}
  ^Operation\s+Mode:\s+${OPERATION_MODE}
  ^Current\s+virtual\s+domain:\s+${CURRENT_VIRTUAL_DOMAIN}
  ^Max\s+number\s+of\s+virtual\s+domains:\s+${MAX_NUMBER_OF_VIRTUAL_DOMAINS}
  ^Virtual\s+domains\s+status:\s+${VIRTUAL_DOMAINS_STATUS}
  ^Virtual\s+domain\s+configuration:\s+${VIRTUAL_DOMAIN_CONFIGURATION}
  ^FIPS-CC\s+mode:\s+${FIPS_CC_MODE}
  ^Current\s+HA\s+mode:\s+${CURRENT_HA_MODE}
  ^Cluster\s+uptime:\s+${CLUSTER_UPTIME}
  ^Cluster\s+state\s+change\s+time:\s+${CLUSTER_STATE_CHANGE_TIME}
  ^Branch\s+point:\s+${BRANCH_POINT}
  ^Release\s+Version\s+Information:\s+${RELEASE_VERSION_INFORMATION}
  ^FortiOS\s+x86-64:\s+${FORTIOS_X86_64}
  ^System\s+time:\s+${SYSTEM_TIME}
  ^Last\s+reboot\s+reason:\s+${LAST_REBOOT_REASON}
  ^\s*$$
  ^. -> Error
`
