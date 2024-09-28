package models

type CiscoNxosDir struct {
	Size	string	`json:"SIZE"`
	Date_time	string	`json:"DATE_TIME"`
	Item_name	string	`json:"ITEM_NAME"`
	Total_size	string	`json:"TOTAL_SIZE"`
	Total_free	string	`json:"TOTAL_FREE"`
	Total_used	string	`json:"TOTAL_USED"`
	File_system	string	`json:"FILE_SYSTEM"`
}