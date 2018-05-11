package volleynet

type Client struct {
	PostUrl     string
	ApiUrl      string
	AmateurPath string
	Cookie      string
}

func DefaultClient() *Client {
	return &Client{
		PostUrl:     "https://beach.volleynet.at/Admin/formular",
		ApiUrl:      "http://www.volleynet.at/api//",
		AmateurPath: "beach/bewerbe/AMATEUR%20TOUR/phase/ABV%20Tour%20AMATEUR%201/sex/M/saison/2018",
	}
}
