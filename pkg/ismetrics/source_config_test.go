package ismetrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSourceConfigList(t *testing.T) {
	tests := []struct {
		in  []byte
		out SourceConfigList
		err bool
	}{
		{
			in: []byte(`
				[
					{
						"key_id": "123",
						"key_filename": "a.pem"
					}
				]
			`),
			out: SourceConfigList{
				{
					KeyID:       "123",
					KeyFilename: "a.pem",
					KeyData:     "",
				},
			},
			err: false,
		},
		{
			in:  []byte(""),
			out: nil,
			err: true,
		},
		{
			in:  []byte("[}"),
			out: nil,
			err: true,
		},
	}

	for _, test := range tests {
		o, err := getSourceConfigList(test.in)
		if test.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}

		assert.Equal(t, test.out, o)
	}
}
