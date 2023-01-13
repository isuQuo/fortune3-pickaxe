package models

import (
	"fmt"
	"time"
)

type Goal struct {
	Name                string
	TotalNeeded         float64
	CurrentSavings      float64
	MonthlyContribution float64
	StartDate           time.Time
	EndDate             time.Time
	MonthsToGoal        int
}

func (g *Goal) Create(name string, totalNeeded, currentSavings, monthlyContribution float64, monthsToGoal int64, startDate, endDate time.Time) (*Goal, error) {
	return &Goal{
		Name:                name,
		CurrentSavings:      currentSavings,
		TotalNeeded:         totalNeeded,
		MonthlyContribution: monthlyContribution,
		StartDate:           startDate,
		EndDate:             endDate,
		MonthsToGoal:        g.MonthsToGoal,
	}, nil
}

func (g *Goal) CalculateSavings() string {
	g.MonthlyContribution = (g.TotalNeeded - g.CurrentSavings) / (float64(g.MonthsToGoal))
	g.StartDate = time.Now()
	g.EndDate = g.StartDate.AddDate(0, g.MonthsToGoal, 0)

	return fmt.Sprintf("Name: %s\nCurrent Savings: $%.2f\nTotal Needed: $%.2f\nMonthly Contribution: $%.2f\nStart Date: %s\nEnd Date: %s\n",
		g.Name, g.CurrentSavings, g.TotalNeeded, g.MonthlyContribution, g.StartDate, g.EndDate)
}
