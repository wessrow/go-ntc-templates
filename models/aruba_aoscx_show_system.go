package models

type ArubaAoscxShowSystem struct {
	Hostname	string	`json:"HOSTNAME"`
	Contact	string	`json:"CONTACT"`
	Location	string	`json:"LOCATION"`
	Vendor	string	`json:"VENDOR"`
	Product	string	`json:"PRODUCT"`
	Serial	string	`json:"SERIAL"`
	Base_mac	string	`json:"BASE_MAC"`
	Version	string	`json:"VERSION"`
	Time_zone	string	`json:"TIME_ZONE"`
	Uptime_weeks	string	`json:"UPTIME_WEEKS"`
	Uptime_days	string	`json:"UPTIME_DAYS"`
	Uptime_hours	string	`json:"UPTIME_HOURS"`
	Uptime_minutes	string	`json:"UPTIME_MINUTES"`
}