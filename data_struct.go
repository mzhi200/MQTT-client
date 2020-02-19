package main

type TpData struct {
	V uint32 `json:"v,omitempty"`
}
type PowerData struct {
	V float32 `json:"v,omitempty"`
}
type DpType struct {
	Tp    []TpData `json:"temperatrue,omitempty"`
	Power []PowerData `json:"power,omitempty"`
}
type date struct {
	Id uint32 `json:"id"`
	Dp DpType `json:"dp,omitempty"`
}
