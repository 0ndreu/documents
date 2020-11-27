package models

type Document struct {
	ID           int64  `json:"id"`
	Document     []byte `json:"document"`
	ClassifierID int32  `json:"classifierID"`
}
