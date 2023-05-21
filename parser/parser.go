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
		MonthlyFee       float32 `json:"monthly_fee"`
		Schedule         struct {
			StartDate time.Time `json:"startDate"`
		} `json:"schedule"`
		Region struct {
			Name     string `json:"nome"`
			Priority int    `json:"prioridade"`
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
	return &response, err
}

func (p *PlanData) Priority() {
	l := p.Plans[0]
	for _, d := range p.Plans {
		if l.Region.Priority > d.Region.Priority {
			l.Region.Priority = d.Region.Priority
		}
	}
}

func (p *PlanData) AdjustPlanDate() {
	for _, p := range p.Plans {
		p.Schedule.StartDate = time.Now().Add(time.Duration(time.Now().Month()))
	}
}
