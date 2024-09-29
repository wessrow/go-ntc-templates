package oneaccess_oneos 

type ShowProductInfoArea struct {
	Mac	string	`json:"MAC"`
	Motherboard	string	`json:"MOTHERBOARD"`
	Manufacturing_date	string	`json:"MANUFACTURING_DATE"`
	Serial	[]string	`json:"SERIAL"`
	Product_name	string	`json:"PRODUCT_NAME"`
	Commercial_name	string	`json:"COMMERCIAL_NAME"`
}

var ShowProductInfoArea_Template = `Value MAC ([\w:]+)
Value MOTHERBOARD (\w.*\w)
Value MANUFACTURING_DATE ([^ ]+)
Value List SERIAL ([A-Z]\d+)
Value PRODUCT_NAME (\w.+\w)
Value COMMERCIAL_NAME (\w.+\w)


Start
  ^\|\Wmac0\W+\|\W${MAC}\W+\|$$
  ^\|\WMotherboard\W[Tt]ype\W+\|\W${MOTHERBOARD}\W+\|$$
  ^\|\WManufacturing\W[Dd]ate\W+\|\W${MANUFACTURING_DATE}\W+\|$$
  ^\|\WSerial\W[Nn]umber\W+\|\W${SERIAL}\W+\|$$
  ^\|\WProduct\W[Nn]ame\W+\|\W${PRODUCT_NAME}\W+\|$$
  ^\|\WCommercial\W[Nn]ame\W+\|\W${COMMERCIAL_NAME}\W+\|$$
  ^\+-+(\+-+)?\+\s*$$
  ^\|\s+Product\s+Info\s+Area\s+\|
  ^\|\s+Key\s+\|\s+Value\s+\|\s*$$
  ^\|\s+(mac\d+|Manufacturing\s+(File\s+Reference|Location)|Mreturn\d+|(Model|PCB|HW)\s+[Rr]evision|Last\s+Testing\s+Date|Sales\s+Code|Mib-\d+\s+system\s+\S+|Software\s+compatibility\s+code|SCAid)\s+\|
  ^\s*$$
  ^. -> Error
`