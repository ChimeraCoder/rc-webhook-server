package main

type Event struct {
	Action string `json:"action"`
	Issue  struct {
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"issue"`
	Repository struct {
		FullName string `json:"full_name"`
		ID       int    `json:"id"`
		Owner    struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
		} `json:"owner"`
	} `json:"repository"`
	Sender struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
	} `json:"sender"`
}
