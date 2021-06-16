package main

import (
	"encoding/json"
	"fmt"
)

// Check is the struct representation of the output of this app
type Check struct {
	Company string `json:"company"`
	Found   bool   `json:"found"`
}

// PrettyString creates a pretty string of the Check that we'll use as
// output
func (c *Check) prettyString() string {
	p := fmt.Sprintf(
		"Company: %s\nFound: %t\n",
		c.Company, c.Found,
	)
	return p
}

// JSON converts the Check struct to JSON, we'll use the JSON string as
// output
func (c *Check) JSON() string {
	cJSON, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(cJSON)
}
