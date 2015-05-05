package direct

import (
	"fmt"
)

type CampaignsService struct {
	client *Client
}

type Campaign struct {
	CampaignID              int     `json:",omitempty"`
	Login                   string  `json:",omitempty"`
	Name                    string  `json:",omitempty"`
	StartDate               string  `json:",omitempty"`
	StrategyName            string  `json:",omitempty"`
	ContextStrategyName     string  `json:",omitempty"`
	Sum                     float32 `json:",omitempty"`
	Rest                    float32 `json:",omitempty"`
	SumAvailableForTransfer float32 `json:",omitempty"`
	Shows                   int     `json:",omitempty"`
	Clicks                  int     `json:",omitempty"`
	Status                  string  `json:",omitempty"`
	StatusShow              string  `json:",omitempty"`
	StatusArchive           string  `json:",omitempty"`
	StatusActivating        string  `json:",omitempty"`
	StatusModerate          string  `json:",omitempty"`
	IsActive                string  `json:",omitempty"`
	ManagerName             string  `json:",omitempty"`
	AgencyName              string  `json:",omitempty"`
	CampaignCurrency        string  `json:",omitempty"`
	SourceCampaignID        int     `json:",omitempty"`
	DayBudgetEnabled        string  `json:",omitempty"`
	EnableRelatedKeywords   string  `json:",omitempty"`
}

type CampaignsListRequest struct {
	Method string   `json:"method"`
	Param  []string `json:"param"`
}

type CampaignsListResponse struct {
	Data []Campaign `json:"data"`
}

type CampaignsCreateOrUpdateRequest struct {
	Method string         `json:"method"`
	Param  CampaignParams `json:"param"`
}

type CampaignsCreateOrUpdateResponse struct {
	Data int `json:"data"`
}

type CampaignsArchiveRequest struct {
	Method string `json:"method"`
	Param  struct {
		CampaignID int
	} `json:"param"`
}

type CampaignsArchiveResponse struct {
	Data int `json:"data"`
}

type CampaignsDeleteRequest struct {
	Method string `json:"method"`
	Param  struct {
		CampaignID int
	} `json:"param"`
}

type CampaignsDeleteResponse struct {
	Data int `json:"data"`
}

type CampaignsActionRequest struct {
	Method string `json:"method"`
	Param  struct {
		CampaignID int
	} `json:"param"`
}

type CampaignsActionResponse struct {
	Data int `json:"data"`
}

type CampaignParams struct {
	Login      string `json:",omitempty"`
	CampaignID int    `json:",omitempty"`
	Name       string `json:",omitempty"`
	FIO        string `json:",omitempty"`
	StartDate  string `json:",omitempty"`
	Currency   string `json:",omitempty"`
	Strategy   struct {
		StrategyName   string  `json:",omitempty"`
		MaxPrice       float32 `json:",omitempty"`
		AveragePrice   float32 `json:",omitempty"`
		WeeklySumLimit float32 `json:",omitempty"`
		AverageCPA     float32 `json:",omitempty"`
		ClicksPerWeek  int     `json:",omitempty"`
		GoalID         int     `json:",omitempty"`
	}
	ContextStrategy struct {
		StrategyName        string  `json:",omitempty"`
		ContextLimit        string  `json:",omitempty"`
		ContextLimitSum     int     `json:",omitempty"`
		ContextPricePercent int     `json:",omitempty"`
		MaxPrice            float32 `json:",omitempty"`
		AveragePrice        float32 `json:",omitempty"`
		AverageCPA          float32 `json:",omitempty"`
		WeeklySumLimit      float32 `json:",omitempty"`
		ClicksPerWeek       int     `json:",omitempty"`
		GoalID              int     `json:",omitempty"`
	}
	AdditionalMetrikaCounters []int  `json:",omitempty"`
	ClickTrackingEnabled      string `json:",omitempty"`
	SmsNotification           struct {
		MetricaSms        string `json:",omitempty"`
		ModerateResultSms string `json:",omitempty"`
		MoneyInSms        string `json:",omitempty"`
		MoneyOutSms       string `json:",omitempty"`
		SmsTimeFrom       string `json:",omitempty"`
		SmsTimeTo         string `json:",omitempty"`
	}
	EmailNotification struct {
		Email             string `json:",omitempty"`
		WarnPlaceInterval int    `json:",omitempty"`
		MoneyWarningValue int    `json:",omitempty"`
		SendAccNews       string `json:",omitempty"`
		SendWarn          string `json:",omitempty"`
	}
	StatusBehavior string `json:",omitempty"`
	TimeTarget     struct {
		ShowOnHolidays  string     `json:",omitempty"`
		HolidayShowFrom int        `json:",omitempty"`
		HolidayShowTo   int        `json:",omitempty"`
		DaysHours       []DayHours `json:",omitempty"`
		TimeZone        string     `json:",omitempty"`
		WorkingHolidays string     `json:",omitempty"`
	}
	StatusContextStop          string   `json:",omitempty"`
	ContextLimit               string   `json:",omitempty"`
	ContextLimitSum            int      `json:",omitempty"`
	ContextPricePercent        int      `json:",omitempty"`
	AutoOptimization           string   `json:",omitempty"`
	StatusMetricaControl       string   `json:",omitempty"`
	DisabledDomains            string   `json:",omitempty"`
	DisabledIps                string   `json:",omitempty"`
	StatusOpenStat             string   `json:",omitempty"`
	ConsiderTimeTarget         string   `json:",omitempty"`
	MinusKeywords              []string `json:",omitempty"`
	AddRelevantPhrases         string   `json:",omitempty"`
	RelevantPhrasesBudgetLimit int      `json:",omitempty"`
	DayBudget                  struct {
		Amount    float32 `json:",omitempty"`
		SpendMode string  `json:",omitempty"`
	}
	MobileBidAdjustment   int    `json:",omitempty"`
	EnableRelatedKeywords string `json:",omitempty"`
}

type DayHours struct {
	Hours    []int
	Days     []int
	BidCoefs []int
}

func (c CampaignsService) GetList() ([]Campaign, error) {
	data := &CampaignsListRequest{Method: "GetCampaignsList"}
	req, err := c.client.NewRequest(data)
	if err != nil {
		return nil, err
	}

	list := new(CampaignsListResponse)
	_, err = c.client.Do(req, list)
	if err != nil {
		return nil, err
	}

	return list.Data, nil
}

func (c CampaignsService) CreateOrUpdate(camp *CampaignParams) (int, error) {
	data := &CampaignsCreateOrUpdateRequest{
		Method: "CreateOrUpdateCampaign",
		Param:  *camp,
	}
	req, err := c.client.NewRequest(data)
	if err != nil {
		return 0, err
	}

	result := new(CampaignsCreateOrUpdateResponse)
	_, err = c.client.Do(req, result)
	if err != nil {
		return 0, err
	}

	return result.Data, nil
}

func (c CampaignsService) Archive(id int) error {
	return c.actionRequest("archive", id)
}

func (c CampaignsService) UnArchive(id int) error {
	return c.actionRequest("unarchive", id)
}

func (c CampaignsService) Delete(id int) error {
	return c.actionRequest("delete", id)
}

func (c CampaignsService) Resume(id int) error {
	return c.actionRequest("resume", id)
}

func (c CampaignsService) Stop(id int) error {
	return c.actionRequest("stop", id)
}

func (c CampaignsService) actionRequest(action string, id int) error {
	data := &CampaignsActionRequest{}
	data.Param.CampaignID = id
	switch action {
	case "delete":
		data.Method = "DeleteCampaign"
	case "archive":
		data.Method = "ArchiveCampaign"
	case "unarchive":
		data.Method = "UnArchiveCampaign"
	case "resume":
		data.Method = "ResumeCampaign"
	case "stop":
		data.Method = "StopCampaign"
	default:
		panic(fmt.Sprintf("Action %q not found!", action))
	}
	req, err := c.client.NewRequest(data)
	if err != nil {
		return err
	}

	result := new(CampaignsActionResponse)
	_, err = c.client.Do(req, result)
	if err != nil {
		return err
	}

	if result.Data != 1 {
		return fmt.Errorf("Something goes wrong!")
	}
	return nil
}
