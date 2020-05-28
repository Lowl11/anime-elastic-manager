package common

type Index struct {
	Name           string `json:"name"`
	Status         string `json:"status"`
	DocumentsCount string `json:"documents_count"`
	HashCode       string `json:"hash_code"`
	Size           string `json:"size"`
	Shards         string `json:"shards"`
}

func (i *Index) String() string {
	var output string
	output += "Name: " + i.Name + "\n"
	output += "Status: " + i.Status + "\n"
	output += "Documents count: " + i.DocumentsCount + "\n"
	output += "Hash code: " + i.HashCode + "\n"
	output += "Shards: " + i.Shards + "\n\n"
	return output
}

type JsonResult struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorMessage string `json:"error_message"`
}
