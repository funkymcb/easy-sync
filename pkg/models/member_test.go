package models

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGenerateUsername(t *testing.T) {
	type Args struct {
		firstName  string
		familyName string
	}
	type Want struct {
		username string
	}

	tests := []struct {
		name string
		args Args
		want Want
	}{
		{
			name: "- simple firstname.familyname example",
			args: Args{
				firstName:  "Jane",
				familyName: "Doe",
			},
			want: Want{
				username: "jane.doe",
			},
		},
		{
			name: "- double name as firstName",
			args: Args{
				firstName:  "Jane Elisabeth",
				familyName: "Doe",
			},
			want: Want{
				username: "jane.doe",
			},
		},
		{
			name: "- double name as familyName",
			args: Args{
				firstName:  "Jane",
				familyName: "Doe Baker",
			},
			want: Want{
				username: "jane.doe",
			},
		},
		{
			name: "- both names are double names",
			args: Args{
				firstName:  "Jane Elisabeth",
				familyName: "Doe Baker",
			},
			want: Want{
				username: "jane.doe",
			},
		},
		{
			name: "- maaaany firstNames",
			args: Args{
				firstName:  "Jane Elisabeth Edith Cruzelda Martha",
				familyName: "Doe",
			},
			want: Want{
				username: "jane.doe",
			},
		},
		{
			name: "- '/' in the name... why would you do that anyway?",
			args: Args{
				firstName:  "Jane/Elisabeth",
				familyName: "Doe/Baker",
			},
			want: Want{
				username: "jane.doe",
			},
		},
		{
			name: "- umlauts for the people",
			args: Args{
				firstName:  "Jäneß",
				familyName: "Dölüe",
			},
			want: Want{
				username: "jaeness.doeluee",
			},
		},
		{
			name: "- phd in the house?",
			args: Args{
				firstName:  "Dr. Jane",
				familyName: "Doe",
			},
			want: Want{
				username: "jane.doe",
			},
		},
	}

	for _, test := range tests {
		gotUsername := GenerateUsername(
			test.args.firstName,
			test.args.familyName,
		)
		assert.Equal(t, gotUsername, test.want.username, test.name)
	}
}
