package models

type CiscoNxosShowVersion struct {
	Uptime	string	`json:"UPTIME"`
	Last_reboot_reason	string	`json:"LAST_REBOOT_REASON"`
	Bios	string	`json:"BIOS"`
	Os	string	`json:"OS"`
	Boot_image	string	`json:"BOOT_IMAGE"`
	Platform	string	`json:"PLATFORM"`
	Hostname	string	`json:"HOSTNAME"`
	Serial	string	`json:"SERIAL"`
}