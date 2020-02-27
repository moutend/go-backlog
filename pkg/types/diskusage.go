package types

type DiskUsage struct {
	Capacity   int `json:"capacity"`
	Issue      int `json:"issue"`
	Wiki       int `json:"wiki"`
	File       int `json:"file"`
	SubVersion int `json:"subversion"`
	Git        int `json:"git"`
	GitLFS     int `json:"gitLFS"`
}

type TotalDiskUsage struct {
	DiskUsage
	Details []DiskUsage `json:"details"`
}
