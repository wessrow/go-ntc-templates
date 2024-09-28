package models

type CiscoIosShowPowerAvailable struct {
	Sys_used_watts	string	`json:"SYS_USED_WATTS"`
	Sys_avail_watts	string	`json:"SYS_AVAIL_WATTS"`
	Inline_used_watts	string	`json:"INLINE_USED_WATTS"`
	Inline_avail_watts	string	`json:"INLINE_AVAIL_WATTS"`
	Backplane_used_watts	string	`json:"BACKPLANE_USED_WATTS"`
	Backplane_avail_watts	string	`json:"BACKPLANE_AVAIL_WATTS"`
	Total_used_watts	string	`json:"TOTAL_USED_WATTS"`
	Max_avail_watts	string	`json:"MAX_AVAIL_WATTS"`
}