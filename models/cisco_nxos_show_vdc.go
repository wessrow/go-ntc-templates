package models

type CiscoNxosShowVdc struct {
	Vdc_id	string	`json:"VDC_ID"`
	Vdc_name	string	`json:"VDC_NAME"`
	State	string	`json:"STATE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Type	string	`json:"TYPE"`
	Lc	string	`json:"LC"`
}