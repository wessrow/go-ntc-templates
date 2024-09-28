package models

type AristaEosDir struct {
	File_system	string	`json:"FILE_SYSTEM"`
	Permissions	string	`json:"PERMISSIONS"`
	Size	string	`json:"SIZE"`
	Date_time	string	`json:"DATE_TIME"`
	Name	string	`json:"NAME"`
	Total_size	string	`json:"TOTAL_SIZE"`
	Total_free	string	`json:"TOTAL_FREE"`
}