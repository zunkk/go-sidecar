package util

import (
	"bytes"
	"encoding/json"

	"github.com/wundergraph/astjson"
)

func BeautifyJSON(compressedJSON string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(compressedJSON), "", "  ")
	if err != nil {
		return compressedJSON
	}
	return out.String()
}

func MergeJsons(jsons ...[]byte) ([]byte, error) {
	if len(jsons) == 0 {
		return nil, nil
	}
	if len(jsons) == 1 {
		return jsons[0], nil
	}
	finalJson, err := astjson.ParseBytes(jsons[0])
	if err != nil {
		return nil, err
	}
	for _, jsonRaw := range jsons[1:] {
		b, err := astjson.ParseBytes(jsonRaw)
		if err != nil {
			return nil, err
		}
		finalJson, _, err = astjson.MergeValues(finalJson, b)
		if err != nil {
			return nil, err
		}
	}
	return []byte(BeautifyJSON(string(finalJson.MarshalTo(nil)))), nil
}
