package models

type CienaSaosSoftwareShow struct {
	Version_installed	string	`json:"VERSION_INSTALLED"`
	Version_running	string	`json:"VERSION_RUNNING"`
	Kernel	string	`json:"KERNEL"`
}