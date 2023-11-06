package types

type DBFilter map[string]any

type ReadQuery struct {
	By 		string
	Item 	any
}

type DeleteRequestParams struct {
	ItemID string `json:"id"`
} 

type ReadRequestParams struct {
	ReadBy string
	Data   any
	//add in pagenation and limits
}