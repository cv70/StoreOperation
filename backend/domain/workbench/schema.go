package workbench

type Metric struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Value string `json:"value"`
	Trend string `json:"trend"`
}

type Card struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Reason    string `json:"reason"`
	Action    string `json:"action"`
	Impact    string `json:"impact"`
	Deadline  string `json:"deadline"`
	Priority  string `json:"priority"`
	CardType  string `json:"card_type"`
	OwnerRole string `json:"owner_role"`
	Status    string `json:"status"`
}

type OverviewResp struct {
	Role          string   `json:"role"`
	Headline      string   `json:"headline"`
	Metrics       []Metric `json:"metrics"`
	Opportunities []Card   `json:"opportunities"`
	Risks         []Card   `json:"risks"`
	Focus         []string `json:"focus"`
}

type TransitionReq struct {
	ToState string `json:"to_state" binding:"required"`
	Reason  string `json:"reason"`
}

type ActionEvent struct {
	CardID    string `json:"card_id"`
	FromState string `json:"from_state"`
	ToState   string `json:"to_state"`
	Reason    string `json:"reason"`
	CreatedAt string `json:"created_at,omitempty"`
}
