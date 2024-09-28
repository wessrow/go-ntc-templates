package models

type CiscoIosDir struct {
	File_system	string	`json:"FILE_SYSTEM"`
	Id	string	`json:"ID"`
	Permissions	string	`json:"PERMISSIONS"`
	Size	string	`json:"SIZE"`
	Total_size	string	`json:"TOTAL_SIZE"`
	Total_free	string	`json:"TOTAL_FREE"`
	Date_time	string	`json:"DATE_TIME"`
	Name	string	`json:"NAME"`
}