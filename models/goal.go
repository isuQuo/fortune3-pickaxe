package models

import "time"

type Goal struct {
	Name                string
	TotalNeeded         float64
	CurrentSavings      float64
	MonthlyContribution float64
	StartDate           time.Time
	EndDate             time.Time
	MonthsToGoal        int
}

func (g *Goal) CalculateSavings() {
	g.MonthlyContribution = (g.TotalNeeded - g.CurrentSavings) / (float64(g.MonthsToGoal))
	g.StartDate = time.Now()
	g.EndDate = g.StartDate.AddDate(0, g.MonthsToGoal, 0)
}
