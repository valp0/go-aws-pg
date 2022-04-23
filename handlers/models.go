package handlers

// Used to send the response in a JSON format to the client.
type Data struct {
	Data interface{} `json:"data"`
}

// Used when the response contains one or more items in an array.
type Items struct {
	Items interface{} `json:"items"`
}

// Used for those cases when a message needs to be displayed inside a JSON.
type Message struct {
	Message interface{} `json:"message"`
}

// Used for the token_details response when /auth/get-token/ is callen.
type Token struct {
	TokenDetails interface{} `json:"token_details"`
}
