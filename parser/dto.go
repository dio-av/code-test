package parser

import "time"

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
