package api

import "encoding/json"

func (p *pageResponse) MarshalBinary() (data []byte, err error) {
	return json.Marshal(p)
}

func (p *pageResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, p)
}
