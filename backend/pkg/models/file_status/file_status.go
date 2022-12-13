package file_status

import (
	"bytes"
	"encoding/json"
)

type FileStatus int

const (
	Pending FileStatus = iota
	InProcess
	Processed
	Failed
)

func (s FileStatus) String() string {
	return toString[s]
}

var toString = map[FileStatus]string{
	Failed:    "Failed",
	InProcess: "InProcess",
	Pending:   "Pending",
	Processed: "Processed",
}

var toId = map[string]FileStatus{
	"Failed":    Failed,
	"InProcess": InProcess,
	"Pending":   Pending,
	"Processed": Processed,
}

func (s FileStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *FileStatus) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	*s = toId[j]
	return nil
}
