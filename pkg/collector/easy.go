package collector

import (
	"fmt"
	"log"
	"strconv"

	"github.com/funkymcb/easy-sync/pkg/models"
	"github.com/go-resty/resty/v2"
)

var (
	client      *resty.Client
	page        int
	easyMembers []models.Member
)

// GetEasyMembersResponse represents the responce from the easyverein /member/ endpoint
type GetEasyMembersResponse struct {
	Next    string          `json:"next"`
	Members []models.Member `json:"results"`
}

// GetEasyMembers queries easyverein API and gets member data
func GetEasyMembers() ([]models.Member, error) {
	var easyMemberResponse *GetEasyMembersResponse
	page += 1

	config := models.GetConfig()

	url := fmt.Sprintf("%s%s%s",
		config.Easyverein.URL,
		config.Easyverein.Endpoint,
		"contact-details",
	)

	log.Printf("GET %s?page=%d", url, page)
	_, err := client.R().
		SetQueryParams(map[string]string{
			"limit": "100",
			"page":  strconv.Itoa(page),
		}).
		SetHeader("Authorization",
			fmt.Sprintf(
				"Token %s",
				config.Easyverein.Token,
			),
		).
		SetResult(&easyMemberResponse).
		Get(url)
	if err != nil {
		return nil, err
	}

	// iterate over pages recursively
	// if next is empty, all pages have been requested
	if easyMemberResponse.Next != "" {
		easyMembers = append(easyMembers, easyMemberResponse.Members...)
		GetEasyMembers()
	} else {
		// append members of last page
		easyMembers = append(easyMembers, easyMemberResponse.Members...)
	}

	return easyMembers, nil
}

func init() {
	client = resty.New()
}
