package email

type SyncResult struct {
	Saved     int `json:"saved"`
	Duplicate int `json:"duplicate"`
	Skipped   int `json:"skipped"`
	Failed    int `json:"failed"`
}