package models

type CiscoS300ShowVersion struct {
	Sw_version	string	`json:"SW_VERSION"`
	Boot_version	string	`json:"BOOT_VERSION"`
	Hw_version	string	`json:"HW_VERSION"`
}