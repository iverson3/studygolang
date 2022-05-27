package db

type DB struct {

}

type BbResult struct {
	data string
}

func (db DB) Exec(args [][]byte) *BbResult {
	return &BbResult{data: "db exec result"}
}

func (r BbResult) ToBytes() []byte {
	return []byte(r.data)
}