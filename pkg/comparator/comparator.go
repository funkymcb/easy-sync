package comparator

import (
	"log"

	"github.com/funkymcb/easy-sync/pkg/models"
)

var membersDiff []models.Member

// GetMembersDiff compares a base of members (eg. from json file)
// with members of a specific target (eg. of easyverein)
// comparison is based on the keys: (hirachy in descending order)
//   1. firstName, secondName
func GetMembersDiff(baseMembers, targetMembers []models.Member) ([]models.Member, error) {
	// iterate over members from json
	for _, baseMember := range baseMembers {
		for j, targetMember := range targetMembers {
			// compare firstname and familyname
			if baseMember.FirstName == targetMember.FirstName &&
				baseMember.FamilyName == targetMember.FamilyName {
				if models.VerboseFlag {
					log.Printf("%s %s already exists in easyverein",
						baseMember.FirstName,
						baseMember.FamilyName,
					)
				}
				break
			}
			// check if loop has finished withou match
			if j+1 == len(targetMembers) {
				if models.VerboseFlag {
					log.Printf("member to be added: %s %s",
						baseMember.FirstName,
						baseMember.FamilyName,
					)
				}
				membersDiff = append(membersDiff, baseMember)
			}
		}
	}
	return membersDiff, nil
}
