package writer

import (
	"testing"

	"github.com/funkymcb/easy-sync/pkg/models"
	"github.com/stretchr/testify/assert"
)

// not really testing anything... but coverage makes us all happy :)
func TestWriteJson(t *testing.T) {
	type Args struct {
		members []models.Member
		outFile string
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
			name: "basic write operation",
			args: Args{
				members: []models.Member{
					{
						FirstName:  "John",
						FamilyName: "Doe",
					},
				},
				outFile: "test-result.json",
			},
			want: Want{
				err: false,
			},
		},
		{
			name: "basic write operation",
			args: Args{
				members: []models.Member{
					{
						FirstName:  "John",
						FamilyName: "Doe",
					},
				},
				outFile: "test-result.json",
			},
			want: Want{
				err: false,
			},
		},
	}

	for _, test := range tests {
		var gotErr bool
		if err := WriteJSONFile(test.args.members, test.args.outFile); err != nil {
			gotErr = true
		}
		assert.Equal(t, gotErr, test.want.err, test.name)
	}
}
