package cdxj

import (
	"github.com/datatogether/warc"
)

type Index []*Record

func (index Index) AddWARCRecord(rec *warc.Record) (Index, error) {
	cdxjRec, err := CreateRecord(rec)
	if err != nil {
		return index, err
	}
	return index.AddRecord(cdxjRec), nil
}

func (index Index) AddRecord(rec *Record) Index {
	for _, r := range index {
		if r.Uri == rec.Uri {
			return index
		}
	}
	return append(index, rec)
}
