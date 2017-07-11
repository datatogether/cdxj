package cdxj

import (
	"fmt"
)

func CompareRecordSlices(a, b []*Record) error {
	if len(a) != len(b) {
		return fmt.Errorf("record slice length mismatch: %d != %d", len(a), len(b))
	}

	for i, ar := range a {
		br := b[i]
		if err := CompareRecords(ar, br); err != nil {
			return fmt.Errorf("record %d mismatch: %s", i, err.Error())
		}
	}

	return nil
}

func CompareRecords(a, b *Record) error {
	if a == nil && b != nil || b == nil && a != nil {
		return fmt.Errorf("nil mistmatch: %s,%s", a, b)
	} else if a == nil && b == nil {
		return nil
	}

	if a.Uri != b.Uri {
		return fmt.Errorf("record uri mismatch: %s != %s", a.Uri, b.Uri)
	}

	if !a.Timestamp.Equal(b.Timestamp) {
		return fmt.Errorf("timestamp mismatch: %s != %s", a.Timestamp.String(), a.Timestamp.String())
	}

	if a.RecordType != b.RecordType {
		return fmt.Errorf("record type mismatch: %s != %s", a.RecordType, b.RecordType)
	}

	// TODO - compare json field

	return nil
}
