package reader

import (
	"testing"

	"github.com/funkymcb/easy-sync/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	type Args struct {
		inFile    string
		outFile   string
		delimiter string
	}
	type Want struct {
		err bool
	}

	tests := []struct {
		name string
		args Args
		want Want
	}{
		{
			name: "everything should be fine",
			args: Args{
				inFile:    "test-files/basic-wvc.csv",
				outFile:   "test-result.json",
				delimiter: ";",
			},
			want: Want{
				err: false,
			},
		},
		{
			name: "invalid delimiter",
			args: Args{
				inFile:    "test-files/basic-wvc.csv",
				outFile:   "test-result.json",
				delimiter: "needs to be single char",
			},
			want: Want{
				err: true,
			},
		},
		{
			name: "inFile does not exist",
			args: Args{
				inFile:    "test-files/does-not-exist.csv",
				outFile:   "test-result.json",
				delimiter: ";",
			},
			want: Want{
				err: true,
			},
		},
	}

	for _, test := range tests {
		var gotErr bool
		err := ReadCSV(test.args.inFile, test.args.outFile, test.args.delimiter)
		if err != nil {
			gotErr = true
		}
		assert.Equal(t, gotErr, test.want.err, test.name)
	}
}

func TestReadDaata(t *testing.T) {
	type Args struct {
		fileName  string
		delimiter string
	}
	type Want struct {
		data [][]string
		err  bool
	}

	tests := []struct {
		name string
		args Args
		want Want
	}{
		{
			name: "basic file",
			args: Args{
				fileName:  "test-files/basic.csv",
				delimiter: ",",
			},
			want: Want{
				data: [][]string{
					{"John", "Doe", "Boston"},
					{"Jane", "Doe", "Baltimore"},
				},
				err: false,
			},
		},
		{
			name: "different delimiter",
			args: Args{
				fileName:  "test-files/semicolon.csv",
				delimiter: ";",
			},
			want: Want{
				data: [][]string{
					{"John", "Doe", "Boston"},
					{"Jane", "Doe", "Baltimore"},
				},
				err: false,
			},
		},
		{
			name: "file does not exist",
			args: Args{
				fileName:  "test-files/does-not-exist.csv",
				delimiter: ";",
			},
			want: Want{
				data: [][]string{},
				err:  true,
			},
		},
		{
			name: "invalid csv file",
			args: Args{
				fileName:  "test-files/corrupted.csv",
				delimiter: ";",
			},
			want: Want{
				data: [][]string{},
				err:  true,
			},
		},
		{
			name: "empty file",
			args: Args{
				fileName:  "test-files/empty.csv",
				delimiter: ";",
			},
			want: Want{
				data: [][]string{},
				err:  true,
			},
		},
	}

	for _, test := range tests {
		var gotErr bool
		gotData, err := readCSVData(test.args.fileName, test.args.delimiter)
		if err != nil {
			gotErr = true
		}

		assert.Equal(t, gotData, test.want.data, test.name)
		assert.Equal(t, gotErr, test.want.err, test.name)

	}
}

func TestParseMembers(t *testing.T) {
	type Args struct {
		records [][]string
	}
	type Want struct {
		members []models.Member
	}

	tests := []struct {
		name string
		args Args
		want Want
	}{
		{
			name: "valid records",
			args: Args{
				records: [][]string{
					{"", "", "Mr.", "John", "Doe", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
					{"", "", "Mrs.", "Jane", "Doe", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
				},
			},
			want: Want{
				members: []models.Member{
					{
						Salution:    "Mr.",
						FirstName:   "John",
						FamilyName:  "Doe",
						Company:     "",
						Street:      "",
						PostalCode:  "",
						Town:        "",
						Telefone:    "",
						IBAN:        "",
						BIC:         "",
						DateOfBirth: "",
						Mobile:      "",
						EMail:       "",
						Title:       "",
						CustomFields: models.CustomFields{
							Gender:     "",
							BankName:   "",
							Username:   "john.doe",
							Fax:        "",
							Department: "",
						},
					},
					{
						Salution:    "Mrs.",
						FirstName:   "Jane",
						FamilyName:  "Doe",
						Company:     "",
						Street:      "",
						PostalCode:  "",
						Town:        "",
						Telefone:    "",
						IBAN:        "",
						BIC:         "",
						DateOfBirth: "",
						Mobile:      "",
						EMail:       "",
						Title:       "",
						CustomFields: models.CustomFields{
							Gender:     "",
							BankName:   "",
							Username:   "jane.doe",
							Fax:        "",
							Department: "",
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		gotMembers := parseMembers(test.args.records)
		assert.Equal(t, gotMembers, test.want.members, test.name)
	}
}
