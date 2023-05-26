package entity

type DataType uint8

const (
	RAW DataType = iota + 1
	CRED
	FILE
	CARD
)
