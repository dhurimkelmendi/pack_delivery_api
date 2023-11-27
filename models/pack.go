package models

// PackSize is a struct that represents a db row of the pack_sizes table
type PackSize struct {
	tableName struct{} `pg:"pack_sizes"`
	Size      int      `json:"size"`
}
