package njalla

type NjallaRecord struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Domain  string `json:"domain"`
	Name    string `json:"name"`
	TTL     int    `json:"ttl"`
	Type    string `json:"type"`
}

type ListRecords struct {
	Method string `json:"method"`
	Params struct {
		Domain string `json:"domain"`
	} `json:"params"`
}

type CreateRecord struct {
	Method string `json:"method"`
	Params struct {
		Domain  string `json:"domain"`
		Name    string `json:"name"`
		Content string `json:"content"`
		TTL     int    `json:"ttl"`
		Type    string `json:"type"`
	} `json:"params"`
}

type EditRecord struct {
	Method string `json:"method"`
	Params struct {
		Domain  string `json:"domain"`
		ID      string `json:"id"`
		Content string `json:"content"`
	} `json:"params"`
}

type RemoveRecord struct {
	Method string `json:"method"`
	Params struct {
		Domain string `json:"domain"`
		ID     string `json:"id"`
	} `json:"params"`
}
