package api

type EmailRequest struct {
	Source                string        `json:"Source" binding:"required"`
	Destination           Destination   `json:"Destination" binding:"required"`
	Message               Message       `json:"Message" binding:"required"`
	ReplyToAddresses      []string      `json:"ReplyToAddresses"`
	ReturnPath            string        `json:"ReturnPath"`
	SourceArn             string        `json:"SourceArn"`
	ReturnPathArn         string        `json:"ReturnPathArn"`
	Tags                  []Tag         `json:"Tags"`
	ConfigurationSetName  string        `json:"ConfigurationSetName"`
}

type Destination struct {
	ToAddresses  []string `json:"ToAddresses" binding:"required"`
	CcAddresses  []string `json:"CcAddresses"`
	BccAddresses []string `json:"BccAddresses"`
}

type Message struct {
	Subject Content `json:"Subject" binding:"required"`
	Body    Body    `json:"Body" binding:"required"`
}

type Content struct {
	Data    string `json:"Data" binding:"required"`
	Charset string `json:"Charset"`
}

type Body struct {
	Text Content `json:"Text"`
	Html Content `json:"Html"`
}

type Tag struct {
	Name  string `json:"Name" binding:"required"`
	Value string `json:"Value" binding:"required"`
}
