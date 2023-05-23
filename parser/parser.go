package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type PlanData struct {
	Device struct {
		Name string `json:"name"`
	} `json:"device"`
	Plans []struct {
		ID               int     `json:"id"`
		Type             string  `json:"type"`
		Name             string  `json:"name"`
		PhonePrice       int     `json:"phonePrice"`
		PhonePriceOnPlan int     `json:"phonePriceOnPlan"`
		Installments     int     `json:"installments"`
		MonthlyFee       float32 `json:"monthlyFee"`
		Schedule         struct {
			StartDate time.Time `json:"startDate"`
		} `json:"schedule"`
		Region struct {
			Name     string `json:"name"`
			Priority int    `json:"priority"`
		} `json:"region"`
	} `json:"plans"`
}

func (p *PlanData) GetJson(path string) (*PlanData, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		err := fmt.Errorf("error openning data json file. %w", err)
		return nil, err
	}

	var response PlanData
	err = json.Unmarshal(file, &response)
	if err != nil {
		err := fmt.Errorf("error unmarshalling json from file. %w", err)
		return nil, err
	}
	return &response, nil
}

// FIX: each plan must have it´s particular priority
func (p *PlanData) Priority() map[string]int {
	l := p.Plans[0].Region.Priority
	t := make(map[string]int)
	for _, d := range p.Plans {
		if l > d.Region.Priority {
			t[d.Name] = d.Region.Priority
		} else {
			t[d.Name] = l
		}
	}
	return t
}

func (p *PlanData) AdjustPlanDate() {
	for i, d := range p.Plans {
		d.Schedule.StartDate = time.Now().Add(time.Hour * 24 * 30)
		p.Plans[i].Schedule.StartDate = d.Schedule.StartDate
	}
}
