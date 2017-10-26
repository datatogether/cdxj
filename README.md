# CDXJ

Golang package implementing the CDXJ file format used by OpenWayback
3.0.0 (and later) to index web archive contents (notably in WARC and
ARC files) and make them searchable via a resource resolution service.
The format builds on the CDX file format originally developed by the
Internet Archive for the indexing behind the WaybackMachine. This
specification builds on it by simplifying the primary fields while
adding a flexible JSON 'block' to each record, allowing high
flexiblity in the inclusion of additional data.

<!-- Repo Badges for: Github Project, Slack, License-->

[![GitHub](https://img.shields.io/badge/project-Data_Together-487b57.svg?style=flat-square)](http://github.com/datatogether)
[![Slack](https://img.shields.io/badge/slack-Archivers-b44e88.svg?style=flat-square)](https://archivers-slack.herokuapp.com/)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](./LICENSE) 

[1-3 sentence description of repository contents]

## License & Copyright

Copyright (C) <year> Data Together
This program is free software: you can redistribute it and/or modify it under
the terms of the GNU AFFERO General Public License as published by the Free Software
Foundation, version 3.0.

This program is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE.

See the [`LICENSE`](./LICENSE) file for details.

## Getting Involved

We would love involvement from more people! If you notice any errors or would like to submit changes, please see our [Contributing Guidelines](./.github/CONTRIBUTING.md). 

We use GitHub issues for [tracking bugs and feature requests](https://github.com/datatogether/REPONAME/issues) and Pull Requests (PRs) for [submitting changes](https://github.com/datatogether/REPONAME/pulls)

## ...

## Installation 

Use in any golang package with:

`import "github.com/datatogether/cdxj"`

## Usage

CDXJ provides the following functions and types:

#### func  CanonicalizeURL

```go
func CanonicalizeURL(rawurl string) string
```
Canonicalization is applied to URIs to remove trivial differences in the URIs
that do not reflect that the URI reference different resources. Examples include
removing session ID parameters, unneccessary port declerations (e.g. :80 when
crawling HTTP). OpenWayback implements its own canonicalization process.
Typically, it will be applied to the searchable URIs in CDXJ files. You can,
however, use any canonicalization scheme you care for (including none). You must
simple ensure that the same canonicalization process is applied to the URIs when
performing searches. Otherwise they may not match correctly. TODO - import
github.com/puerkitobio/purell to canonicalize urls

#### func  SURTUrl

```go
func SURTUrl(rawurl string) (string, error)
```
SURTUrl is a transformation applied to URIs which makes their left-to-right
representation better match the natural hierarchy of domain names. A URI
`<scheme://domain.tld/path?query>` has SURT form
`<scheme://(tld,domain,)/path?query>`. Conversion to SURT form also involves
making all characters lowercase, and changing the 'https' scheme to 'http'.
Further, the '/' after a URI authority component -- for example, the third slash
in a regular HTTP URI -- will only appear in the SURT form if it appeared in the
plain URI form.

#### func  UnSURTUrl

```go
func UnSURTUrl(surturl string) (string, error)
```
UnSURTUrl turns a SURT'ed url back into a normal Url TODO - should accept SURT
urls that contain a scheme

#### type Reader

```go
type Reader struct {
}
```

A Reader reads records from a CSV-encoded file.

As returned by NewReader, a Reader expects input conforming to RFC 4180. The
exported fields can be changed to customize the details before the first call to
Read or ReadAll.

#### func  NewReader

```go
func NewReader(r io.Reader) *Reader
```
NewReader returns a new Reader that reads from r.

#### func (*Reader) Read

```go
func (r *Reader) Read() (*Record, error)
```
Read reads a record from the reader err io.EOF will be returned when the last
record is reached

#### func (*Reader) ReadAll

```go
func (r *Reader) ReadAll() ([]*Record, error)
```
ReadAll consumes the entire reader, returning a slice of records

#### type Record

```go
type Record struct {
	// Searchable URI
	// By *searchable*, we mean that the following transformations have been applied to it:
	// 1. Canonicalization - See Appendix A
	// 2. Sort-friendly URI Reordering Transform (SURT)
	// 3. The scheme is dropped from the SURT format
	Uri string
	// should correspond to the WARC-Date timestamp as of WARC 1.1.
	// The timestamp shall represent the instant that data capture for record
	// creation began.
	// All timestamps should be in UTC.
	Timestamp time.Time
	// Indicates what type of record the current line refers to.
	// This field is fully compatible with WARC 1.0 definition of
	// WARC-Type (chapter 5.5 and chapter 6).
	RecordType string
	// This should contain fully valid JSON data. The only limitation, beyond those
	// imposed by JSON encoding rules, is that this may not contain any newline
	// characters, either in Unix (0x0A) or Windows form (0x0D0A).
	// The first occurance of a 0x0A constitutes the end of this field (and the record).
	JSON map[string]interface{}
}
```

Following the header lines, each additional line should represent exactly one
resource in a web archive. Typically in a WARC (ISO 28500) or ARC file, although
the exact storage of the resource is not defined by this specification. Each
such line shall be refered to as a *record*.

#### func (*Record) MarshalCDXJ

```go
func (r *Record) MarshalCDXJ() ([]byte, error)
```
MarshalCDXJ outputs a CDXJ representation of r

#### func (*Record) UnmarshalCDXJ

```go
func (r *Record) UnmarshalCDXJ(data []byte) (err error)
```
UnmarshalCDXJ reads a cdxj record from a byte slice

#### type Records

```go
type Records [][]byte
```

Records implements sortable for a slice marshaled CDXJ byte slices

#### func (Records) Len

```go
func (a Records) Len() int
```

#### func (Records) Less

```go
func (a Records) Less(i, j int) bool
```

#### func (Records) Swap

```go
func (a Records) Swap(i, j int)
```

#### type Writer

```go
type Writer struct {
}
```

Writer writes to an io.Writer, create one with NewWriter You *must* call call
Close to write the record to the specified writer

#### func  NewWriter

```go
func NewWriter(w io.Writer) *Writer
```
NewWriter allocates a new CDXJ Writer

#### func (*Writer) Close

```go
func (w *Writer) Close() error
```
Close dumps the writer to the underlying io.Writer

#### func (*Writer) Write

```go
func (w *Writer) Write(r *Record) error
```
Write a record to the writer

## Development

Coming Soon

