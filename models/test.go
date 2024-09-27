package models

type Cisco_ios_show_version struct {
	SOFTWARE_IMAGE  string   `json:"SOFTWARE_IMAGE"`
	VERSION         string   `json:"VERSION"`
	RELEASE         string   `json:"RELEASE"`
	ROMMON          string   `json:"ROMMON"`
	HOSTNAME        string   `json:"HOSTNAME"`
	UPTIME          string   `json:"UPTIME"`
	UPTIME_YEARS    string   `json:"UPTIME_YEARS"`
	UPTIME_WEEKS    string   `json:"UPTIME_WEEKS"`
	UPTIME_DAYS     string   `json:"UPTIME_DAYS"`
	UPTIME_HOURS    string   `json:"UPTIME_HOURS"`
	UPTIME_MINUTES  string   `json:"UPTIME_MINUTES"`
	RELOAD_REASON   string   `json:"RELOAD_REASON"`
	RUNNING_IMAGE   string   `json:"RUNNING_IMAGE"`
	HARDWARE        []string `json:"HARDWARE"`
	SERIAL          []string `json:"SERIAL"`
	CONFIG_REGISTER string   `json:"CONFIG_REGISTER"`
	MAC_ADDRESS     []string `json:"MAC_ADDRESS"`
	RESTARTED       string   `json:"RESTARTED"`
}
