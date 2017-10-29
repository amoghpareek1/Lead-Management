package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	UUID     string
}

// Response ...
type Response struct {
	Success bool
	Data    interface{}
}

// ConnectionType ...
type ConnectionType uint

const (
	// Salesforce ...
	Salesforce ConnectionType = iota + 1
	// MySQL ...
	MySQL
)

// Connection ...
type Connection struct {
	gorm.Model
	SalesforceConnection
	MySQLConnection
	OwnerID uint
	Type    ConnectionType
}

// MarshalJSON ...
func (c *Connection) MarshalJSON() ([]byte, error) {
	output := make(map[string]interface{})

	output["ID"] = c.ID

	switch int(c.Type) {
	case 1:
		output["Type"] = "Salesforce"
	case 2:
		output["Type"] = "MySQL"
	}

	b, _ := json.Marshal(&output)

	return b, nil
}

// UnmarshalJSON ...
func (c *Connection) UnmarshalJSON(data []byte) error {
	log.Println("ok")
	input := make(map[string]interface{})
	_ = json.Unmarshal(data, &input)

	c.SalesforceName = input["SalesforceName"].(string)
	c.SalesforceUsername = input["SalesforceUsername"].(string)
	c.SalesforcePassword = input["SalesforcePassword"].(string)
	c.SalesforceSecurityToken = input["SalesforceSecurityToken"].(string)
	c.SalesforceAPIVersion = input["SalesforceAPIVersion"].(string)

	c.MySQLName = input["MySQLName"].(string)
	c.MySQLServer = input["MySQLServer"].(string)
	c.MySQLPort = input["MySQLPort"].(string)
	c.MySQLUserID = input["MySQLUserID"].(string)
	c.MySQLPassword = input["MySQLPassword"].(string)
	c.MySQLDatabase = input["MySQLDatabase"].(string)

	c.OwnerID = input["OwnerID"].(uint)

	switch input["Type"].(string) {
	case "Salesforce":
		c.Type = 1
	case "MySQL":
		c.Type = 2
	default:
		return errors.New("UNKNOWN CONNECTION TYPE VALUE")
	}

	return nil
}

// SalesforceConnection ...
type SalesforceConnection struct {
	SalesforceName          string
	SalesforceUsername      string
	SalesforcePassword      string
	SalesforceSecurityToken string
	SalesforceAPIVersion    string
}

// MySQLConnection ...
type MySQLConnection struct {
	MySQLName     string
	MySQLServer   string
	MySQLPort     string
	MySQLUserID   string
	MySQLPassword string
	MySQLDatabase string
}

// Job ...
type Job struct {
	gorm.Model
	Name                 string
	SourceConnectionID   uint
	SourceConnectionType string
	TargetConnectionID   uint
	TargetConnectionType string
}
