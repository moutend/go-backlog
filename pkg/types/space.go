package types

type Space struct {
	SpaceKey          string `json:"spaceKey"`
	Name              string `json:"name"`
	OwnerId           uint64 `json:"ownerId"`
	Lang              string `json:"lang"`
	ReportSendTime    string `json:"reportSendTime"`
	TextFormatingRule string `json:"textFormattingRule"`
	Timezone          string `json:"timezone"`
	Created           *Date  `json:"created"`
	Updated           *Date  `json:"updated"`
}
