package manager

type Index struct {
	Name           string
	Status         string
	DocumentsCount string
	HashCode       string
	Size           string
	Shards         string
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
	Status       string
	Message      string
	ErrorMessage string
}
