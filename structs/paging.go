package structs

type Paging struct {
	Cursors Cursors `json:"cursors"`
}

type Cursors struct {
	After  string `json:"after"`
	Before string `json:"before"`
}
