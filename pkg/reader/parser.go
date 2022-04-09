package reader

import "github.com/funkymcb/easy-sync/pkg/models"

// parseMembers reads the csv records and parses them to the members struct
func parseMembers(records [][]string) []models.Member {
	var members []models.Member
	// iterate over each record and parse it into member struct
	for _, record := range records {
		username := models.GenerateUsername(record[3], record[4])

		member := models.Member{
			Salution:    record[2],
			FirstName:   record[3],
			FamilyName:  record[4],
			Company:     record[5],
			Street:      record[8],
			PostalCode:  record[9],
			Town:        record[10],
			Telefone:    record[11],
			IBAN:        record[13],
			BIC:         record[14],
			DateOfBirth: record[15],
			Mobile:      record[18],
			EMail:       record[19],
			Title:       record[31],
			CustomFields: models.CustomFields{
				Gender:     record[21],
				BankName:   record[25],
				Username:   username,
				Fax:        record[17],
				Department: record[28],
			},
			// Ausw:			  record[0],
			// MemberID:          record[1],
			// ShippingMethod:    record[6],
			// AddressSupplement: record[7],
			// Country:           record[12],
			// BusinessTelefone:  record[16],
			// Nationality:       record[20],
			// MartialStatus:     record[22],
			// Job:               record[23],
			// Status:            record[24],
			// Entry:             record[26],
			// Termination:       record[27],
			// Department:        record[28],
			// Function:          record[29],
			// MandateNumber:     record[30],
		}
		members = append(members, member)
	}
	return members
}
