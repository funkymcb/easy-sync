package models

import (
	"fmt"
	"strings"
)

// Member represents the members of our club
// change this model to your needs
type Member struct {
	Salution     string       `json:"salution"`     // Anrede;
	FirstName    string       `json:"firstName"`    // Vorname;
	FamilyName   string       `json:"familyName"`   // Nachname;
	Company      string       `json:"companyName"`  // Firma;
	Street       string       `json:"street"`       // Strasse;
	PostalCode   string       `json:"zip"`          // PLZ;
	Town         string       `json:"city"`         // Ort;
	Telefone     string       `json:"privatePhone"` // Telefon;
	IBAN         string       `json:"iban"`         // IBAN;
	BIC          string       `json:"bic"`          // BIC;
	DateOfBirth  string       `json:"dateOfBirth"`  // Geburtsdatum;
	Mobile       string       `json:"mobilePhone"`  // Mobil;
	EMail        string       `json:"privateEmail"` // E-Mail;
	Title        string       `json:"nameAffix"`    // Titel
	CustomFields CustomFields `json:"customFields"`
}

// CustomFields stuct represents the CustomFields
// which we created for our members.
// Change to your needs
type CustomFields struct {
	Gender           string `json:"gender"`
	BankName         string `json:"bankName"`
	Username         string `json:"username"`
	Fax              string `json:"fax"`
	TypeOfMembership string `json:"typeOfMembership"`
	Department       string `json:"department"`
}

// GenerateUsername creates a username based on firstName and familyName
// username will look like: firstname.familyname
func GenerateUsername(firstName, familyName string) string {
	firstName = strings.ToLower(firstName)
	familyName = strings.ToLower(familyName)
	// list all replacements here
	replacer := strings.NewReplacer(
		"/", " ",
		"ä", "ae",
		"ö", "oe",
		"ü", "ue",
		"ß", "ss",
		"dr. ", "",
		"(", " ",
	)
	firstName = replacer.Replace(firstName)
	familyName = replacer.Replace(familyName)

	// just use 'first' firstName in case of multiple names
	firstName = strings.Split(firstName, " ")[0]
	// just use first familyName in case of multiple names
	familyName = strings.Split(familyName, " ")[0]

	username := fmt.Sprintf("%s.%s", firstName, familyName)
	return username
}
