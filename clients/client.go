package clients

type Client struct {
	Name       string
	Age        int
	CPF        string
	Occupation string
}

func CreateClient(name string, age int, cpf string, occupation string) *Client {
	return &Client{
		Name:       name,
		Age:        age,
		CPF:        cpf,
		Occupation: occupation,
	}
}
