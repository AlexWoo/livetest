package flv

import (
	"io"
	"media"
)

type FLVScriptData struct {
	ScriptData []byte
}

func (data *FLVScriptData) Parser(r io.Reader, b []byte) (int, error) {
	data.ScriptData = b

	return media.OK, nil
}

func (data *FLVScriptData) Sprint() string {
	return ""
}
