package forms

// CreateClientForm ...
type CreateClientForm struct {
	ClientName       string `form:"client_name" json:"clientName" binding:"required,min=3,max=100"`
	Address          string `form:"address" json:"address" binding:"required,min=3,max=200"`
	ArmedUnitPrice   int    `form:"armed_unit_price" json:"armedUnitPrice" binding:"required"`
	ArmedCount       int    `form:"armed_count" json:"armedCount" binding:"required"`
	UnarmedUnitPrice int    `form:"unarmed_unit_price" json:"unarmedUnitPrice" binding:"required"`
	UnarmedCount     int    `form:"unarmed_count" json:"unarmedCount" binding:"required"`
	TotalGuards      *int   `form:"total_guards" json:"totalGuards" binding:"omitempty,min=1"`
	StartDate        string `form:"start_date" json:"startDate" binding:"required"`
	Contact          string `form:"contact" json:"contact" binding:"required,min=3,max=16"`
}

// UpdateClientForm ...
type UpdateClientForm struct {
	ID               string  `form:"id" json:"id" binding:"required,len=6"`
	ClientName       *string `form:"client_name" json:"clientName" binding:"omitempty,min=3,max=100"`
	Address          *string `form:"address" json:"address" binding:"omitempty,min=3,max=200"`
	ArmedUnitPrice   *int    `form:"armed_unit_price" json:"armedUnitPrice" binding:"omitempty"`
	ArmedCount       *int    `form:"armed_count" json:"armedCount" binding:"omitempty"`
	UnarmedUnitPrice *int    `form:"unarmed_unit_price" json:"unarmedUnitPrice" binding:"omitempty"`
	UnarmedCount     *int    `form:"unarmed_count" json:"unarmedCount" binding:"omitempty"`
	TotalGuards      *int    `form:"total_guards" json:"totalGuards" binding:"omitempty,min=1"`
	StartDate        *string `form:"start_date" json:"startDate" binding:"omitempty"`
	Contact          *string `form:"contact" json:"contact" binding:"omitempty,min=3,max=16"`
}

// CreateClientFormMessages defines validation error messages for create client form.
var CreateClientFormMessages = ValidationMessages{
	"ClientName": {
		"required": "clientName is required",
		"min":      "clientName should be gerater than 3 characters",
	},
	"Address": {
		"required": "Please enter the address",
		"min":      "address should be between 3 to 200 characters",
	},
	"ArmedUnitPrice": {
		"required": "Please enter the armedUnitPrice",
		"min":      "armedUnitPrice should be greater than 0",
	},
	"ArmedCount": {
		"required": "Please enter the armedCount",
		"min":      "armedCount should be greater than 0",
	},
	"UnarmedUnitPrice": {
		"required":      "Please enter the unarmedUnitPrice",
		"min										": "unarmedUnitPrice should be greater than 0",
	},
	"UnarmedCount": {
		"required":      "Please enter the unarmedCount",
		"min										": "unarmedCount should be greater than 0",
	},
	"TotalGuards": {
		"required":      "Please enter the totalGuards",
		"min										": "totalGuards should be greater than 0",
	},
	"StartDate": {
		"required": "Please enter the startDate",
	},
	"Contact": {
		"required":      "Please enter the contact",
		"min										": "contact should be between 3 to 16 characters",
		"max										": "contact should be between 3 to 16 characters",
	},
}

// UpdateClientFormMessages defines validation error messages for update client form.
var UpdateClientFormMessages = ValidationMessages{
	"ID": {
		"required": "id is required",
		"len":      "id must be 6 characters",
	},
	"ClientName": {
		"required": "Please enter the clientName",
		"min":      "clientName should be gerater than 3 characters",
	},
	"Address": {
		"required": "Please enter the client address",
		"min":      "address should be between 3 to 200 characters",
	},
	"ArmedUnitPrice": {
		"required": "Please enter the armedUnitPrice",
		"min":      "armedUnitPrice should be greater than 0",
	},
	"ArmedCount": {
		"required": "Please enter the armedCount",
		"min":      "armedCount should be greater than 0",
	},
	"UnarmedUnitPrice": {
		"required":      "Please enter the unarmedUnitPrice",
		"min										": "unarmedUnitPrice should be greater than 0",
	},
	"UnarmedCount": {
		"required":      "Please enter the unarmedCount",
		"min										": "unarmedCount should be greater than 0",
	},
	"TotalGuards": {
		"required":      "Please enter the totalGuards",
		"min										": "totalGuards should be greater than 0",
	},
	"StartDate": {
		"required": "Please enter the startDate",
	},
	"Contact": {
		"required":      "Please enter the contact",
		"min										": "contact should be between 3 to 16 characters",
		"max										": "contact should be between 3 to 16 characters",
	},
}
