package api

type ErrorResponse struct {
	Code  int
	Error string
}

type Demo struct {
	Name        string
	Type        int
	Description string
}

var Demos = []Demo{
	{
		Name:        "Test01",
		Type:        1,
		Description: "Testing 1",
	},
	{
		Name:        "Test02",
		Type:        1,
		Description: "Testing 2",
	},
	{
		Name:        "Test03",
		Type:        2,
		Description: "Testing 3",
	},
	{
		Name:        "Test04",
		Type:        3,
		Description: "Testing 4",
	},
}
