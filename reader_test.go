package cdxj

import (
	"strings"
	"testing"
	"time"
)

const eg = `!OpenWayback-CDXJ 1.0
(com,cnn,)/world> 2015-09-03T13:27:52Z response {"a":0,"b":"b","c":false}
(uk,ac,rpms,)/> 2015-09-03T13:27:52Z request {"frequency":241,"spread":3}
(uk,co,bbc,)/images> 2015-09-03T13:27:52Z response {"frequency":725,"spread":1}
`

var parsed = []*Record{
	&Record{"cnn.com/world", time.Date(2015, time.September, 3, 13, 27, 52, 0, time.UTC), "response", map[string]interface{}{"a": 0, "b": "b", "c": false}},
	&Record{"rpms.ac.uk/", time.Date(2015, time.September, 3, 13, 27, 52, 0, time.UTC), "request", map[string]interface{}{"frequency": 241, "spread": 3}},
	&Record{"bbc.co.uk/images", time.Date(2015, time.September, 3, 13, 27, 52, 0, time.UTC), "response", map[string]interface{}{"frequency": 725, "spread": 1}},
}

func TestReader(t *testing.T) {
	records, err := NewReader(strings.NewReader(eg)).ReadAll()
	if err != nil {
		t.Error(err)
		return
	}

	if err := CompareRecordSlices(records, parsed); err != nil {
		t.Error(err)
		return
	}
}
