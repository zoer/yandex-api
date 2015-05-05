package direct

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCampaigns_List(t *testing.T) {
	setup()
	defer teardown()

	a := assert.New(t)

	c := Campaign{
		SumAvailableForTransfer: 123.12,
		CampaignID:              1,
		Login:                   "test",
		Name:                    "John",
		StartDate:               "2011-01-01",
		StrategyName:            "strategy",
		ContextStrategyName:     "other strategy",
		Sum:                     12.3,
		Rest:                    0.3,
		Shows:                   1432,
		Clicks:                  8523,
		Status:                  "archived",
		StatusShow:              "closed",
		StatusArchive:           "ok",
		StatusActivating:        "fail",
		StatusModerate:          "done",
		IsActive:                "no",
		ManagerName:             "Eric",
		AgencyName:              "Magazin",
		CampaignCurrency:        "rub",
		SourceCampaignID:        452,
		DayBudgetEnabled:        "sometimes",
		EnableRelatedKeywords:   "never",
	}

	mux.HandleFunc(client.BaseURL.RequestURI(), func(w http.ResponseWriter, r *http.Request) {
		b, _ := json.Marshal(c)
		fmt.Fprintf(w, `{"data":[%s]}`, string(b))
	})

	campaigns, err := client.Campaigns.GetList()
	a.NoError(err)
	a.Equal(campaigns[0], c)
}

func TestCampaigns_CreateOrUpdate(t *testing.T) {
	setup()
	defer teardown()

	a := assert.New(t)

	c := CampaignParams{
		Login:                      "my-login",
		CampaignID:                 765,
		Name:                       "Eric",
		FIO:                        "Eric Cartman",
		StartDate:                  "2012-04-12",
		Currency:                   "rub",
		AdditionalMetrikaCounters:  []int{1, 23, 5},
		ClickTrackingEnabled:       "yes",
		StatusBehavior:             "super behavior",
		StatusContextStop:          "nothing",
		ContextLimit:               "one thing",
		ContextLimitSum:            823,
		ContextPricePercent:        12,
		AutoOptimization:           "no optimization",
		StatusMetricaControl:       "meter",
		DisabledDomains:            "a,b",
		DisabledIps:                "10.0.0.1",
		StatusOpenStat:             "free",
		ConsiderTimeTarget:         "limitless",
		MinusKeywords:              []string{"1", "as"},
		AddRelevantPhrases:         "boo",
		RelevantPhrasesBudgetLimit: 12353,
		MobileBidAdjustment:        8923,
		EnableRelatedKeywords:      "yes,no,mb",
	}
	c.SmsNotification.MetricaSms = "call me 1"
	c.SmsNotification.ModerateResultSms = "call me 2"
	c.SmsNotification.MoneyInSms = "call me 3 "
	c.SmsNotification.MoneyOutSms = "call me 4"
	c.SmsNotification.SmsTimeFrom = "call me 5"
	c.SmsNotification.SmsTimeTo = "call me 6"
	c.EmailNotification.Email = "foo@moo.ru"
	c.EmailNotification.WarnPlaceInterval = 235
	c.EmailNotification.MoneyWarningValue = 8523
	c.EmailNotification.SendAccNews = "cuba"
	c.EmailNotification.SendWarn = "chuba"
	c.TimeTarget.ShowOnHolidays = "always"
	c.TimeTarget.HolidayShowFrom = 1
	c.TimeTarget.HolidayShowTo = 7
	c.TimeTarget.DaysHours = []DayHours{
		DayHours{[]int{1, 2}, []int{1, 5}, []int{1, 3}},
	}
	c.TimeTarget.TimeZone = "europe/moscow"
	c.TimeTarget.WorkingHolidays = "never"
	c.ContextStrategy.StrategyName = "eagle"
	c.ContextStrategy.ContextLimit = "nothing limited"
	c.ContextStrategy.ContextLimitSum = 523
	c.ContextStrategy.ContextPricePercent = 923
	c.ContextStrategy.MaxPrice = 42.5
	c.ContextStrategy.AveragePrice = 52.5
	c.ContextStrategy.AverageCPA = 58.2
	c.ContextStrategy.WeeklySumLimit = 8295
	c.ContextStrategy.ClicksPerWeek = 5235
	c.ContextStrategy.GoalID = 215
	c.Strategy.StrategyName = "bad name"
	c.Strategy.MaxPrice = 124.3
	c.Strategy.AveragePrice = 12.3
	c.Strategy.WeeklySumLimit = 523.3
	c.Strategy.AverageCPA = 12.4
	c.Strategy.ClicksPerWeek = 843
	c.Strategy.GoalID = 3849
	c.DayBudget.Amount = 42.1
	c.DayBudget.SpendMode = "some meat"

	mux.HandleFunc(client.BaseURL.RequestURI(), func(w http.ResponseWriter, r *http.Request) {
		bc := new(CampaignsCreateOrUpdateRequest)
		json.NewDecoder(r.Body).Decode(bc)
		a.Equal("CreateOrUpdateCampaign", bc.Method)
		a.Equal(c, bc.Param)

		fmt.Fprintf(w, `{"data":%d}`, c.CampaignID)
	})

	id, err := client.Campaigns.CreateOrUpdate(&c)
	a.NoError(err)
	a.Equal(id, c.CampaignID)
}

func TestCampaigns_Archive(t *testing.T) {
	setup()
	defer teardown()

	a := assert.New(t)

	id := 321

	mux.HandleFunc(client.BaseURL.RequestURI(), func(w http.ResponseWriter, r *http.Request) {
		bc := new(CampaignsActionRequest)
		json.NewDecoder(r.Body).Decode(bc)
		a.Equal("ArchiveCampaign", bc.Method)
		a.Equal(id, bc.Param.CampaignID)

		fmt.Fprintf(w, `{"data":%d}`, 1)
	})

	err := client.Campaigns.Archive(id)
	a.NoError(err)
}

func TestCampaigns_UnArchive(t *testing.T) {
	setup()
	defer teardown()

	a := assert.New(t)

	id := 321

	mux.HandleFunc(client.BaseURL.RequestURI(), func(w http.ResponseWriter, r *http.Request) {
		bc := new(CampaignsActionRequest)
		json.NewDecoder(r.Body).Decode(bc)
		a.Equal("UnArchiveCampaign", bc.Method)
		a.Equal(id, bc.Param.CampaignID)

		fmt.Fprintf(w, `{"data":%d}`, 1)
	})

	err := client.Campaigns.UnArchive(id)
	a.NoError(err)
}

func TestCampaigns_Delete(t *testing.T) {
	setup()
	defer teardown()

	a := assert.New(t)

	id := 321

	mux.HandleFunc(client.BaseURL.RequestURI(), func(w http.ResponseWriter, r *http.Request) {
		bc := new(CampaignsActionRequest)
		json.NewDecoder(r.Body).Decode(bc)
		a.Equal("DeleteCampaign", bc.Method)
		a.Equal(id, bc.Param.CampaignID)

		fmt.Fprintf(w, `{"data":%d}`, 1)
	})

	err := client.Campaigns.Delete(id)
	a.NoError(err)
}

func TestCampaigns_Resume(t *testing.T) {
	setup()
	defer teardown()

	a := assert.New(t)

	id := 321

	mux.HandleFunc(client.BaseURL.RequestURI(), func(w http.ResponseWriter, r *http.Request) {
		bc := new(CampaignsActionRequest)
		json.NewDecoder(r.Body).Decode(bc)
		a.Equal("ResumeCampaign", bc.Method)
		a.Equal(id, bc.Param.CampaignID)

		fmt.Fprintf(w, `{"data":%d}`, 1)
	})

	err := client.Campaigns.Resume(id)
	a.NoError(err)
}

func TestCampaigns_Stop(t *testing.T) {
	setup()
	defer teardown()

	a := assert.New(t)

	id := 321

	mux.HandleFunc(client.BaseURL.RequestURI(), func(w http.ResponseWriter, r *http.Request) {
		bc := new(CampaignsActionRequest)
		json.NewDecoder(r.Body).Decode(bc)
		a.Equal("StopCampaign", bc.Method)
		a.Equal(id, bc.Param.CampaignID)

		fmt.Fprintf(w, `{"data":%d}`, 1)
	})

	err := client.Campaigns.Stop(id)
	a.NoError(err)
}
