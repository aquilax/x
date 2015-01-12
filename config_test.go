package x

import (
	"reflect"
	"strings"
	"testing"
)

func TestConfig(t *testing.T) {
	var testCases = []struct {
		yaml     string
		expected *Config
		err      string
	}{
		{
			yaml: `android-tablets:
  sizes:
    - max_width: 500
    - max: 1000
  formats:
    - jpg
    - png
  store:
    - local: /mnt/gallery`,
			expected: &Config{
				"android-tablets": Handler{
					[]Size{
						Size{MaxWidth, 500},
						Size{Max, 1000},
					},
					[]Format{JPG, PNG},
					[]storeConfig{
						{
							"local": "/mnt/gallery",
						},
					},
				},
			},
			err: "",
		},
		{
			yaml:     `asd: 123`,
			expected: &Config{},
			err:      "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!int `123` into x.Handler",
		},
	}
	for _, tc := range testCases {
		c := NewConfig()
		err := c.LoadStream(strings.NewReader(tc.yaml))
		if !reflect.DeepEqual(tc.expected, c) {
			t.Errorf("Expected %q, got %q for %q", tc.expected, c, tc.yaml)
		}
		if (err != nil) && (err.Error() != tc.err) {
			t.Errorf("Expected error %q, got %q for %q", tc.err, err.Error(), tc.yaml)
		}
	}
}
