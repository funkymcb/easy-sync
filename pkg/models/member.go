package models

// Member represents the members of our club
// change this model to your needs
type Member struct {
	Ausw              string `json:"ausw"`               // Ausw; TODO what is this field?
	MemberID          string `json:"member_id"`          // MitglNr;
	Salution          string `json:"salution"`           // Anrede;
	FirstName         string `json:"first_name"`         // Vorname;
	SecondName        string `json:"second_name"`        // Nachname;
	Company           string `json:"company"`            // Firma;
	ShippingMethod    string `json:"shipping_method"`    // Versandart;
	AddressSupplement string `json:"address_supplement"` // Adresszusatz;
	Street            string `json:"street"`             // Strasse;
	PostalCode        string `json:"postal_code"`        // PLZ;
	Town              string `json:"town"`               // Ort;
	Telefone          string `json:"telefone"`           // Telefon;
	Country           string `json:"country"`            // Landname;
	IBAN              string `json:"iban"`               // IBAN;
	BIC               string `json:"bic"`                // BIC;
	BirthDate         string `json:"birth_date"`         // Geburtsdatum;
	BusinessTelefone  string `json:"business_telefone"`  // Telefon (gesch.);
	Fax               string `json:"fax"`                // Fax; really?
	Mobile            string `json:"mobile"`             // Mobil;
	EMail             string `json:"e_mail"`             // E-Mail;
	Nationality       string `json:"nationality"`        // Staatsangehörigkeit;
	Gender            string `json:"gender"`             // Geschlecht;
	MartialStatus     string `json:"martial_status"`     // Familienstand;
	Job               string `json:"job"`                // Beruf;
	Status            string `json:"status"`             // Status; TODO what status?
	BankDesignation   string `json:"bank_designation"`   // Bankbezeichnung;
	Entry             string `json:"entry"`              // Estringritt;
	Termination       string `json:"termination"`        // Austritt;
	Department        string `json:"department"`         // Abteilung;
	Function          string `json:"function"`           // Funktionen;
	MandateNumber     string `json:"mandate_number"`     // MandatsNr;
	Title             string `json:"title"`              // Titel
}
