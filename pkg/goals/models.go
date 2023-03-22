package goals

import (
	"fmt"
	"strconv"
	"time"
)

type Goal struct {
	Name                string  `db:"name"`
	TotalNeeded         float64 `db:"total_needed"`
	CurrentSavings      float64 `db:"current_savings"`
	MonthlyContribution float64 `db:"monthly_contribution"`
	StartDate           string  `db:"start_date"`
	EndDate             string  `db:"end_date"`
	MonthsToGoal        int     `db:"months_to_goal"`
}

// Validate parses and validates the goal fields and returns a Goal struct or an error
func Validate(name, totalNeeded, currentSavings, monthlyContribution, startDate, endDate, monthsToGoal string) (Goal, error) {
	var errs []error

	if name == "" || totalNeeded == "" || currentSavings == "" || monthlyContribution == "" || monthsToGoal == "" || startDate == "" || endDate == "" {
		return Goal{}, fmt.Errorf("all fields are required")
	}

	totalNeededP, err := strconv.ParseFloat(totalNeeded, 64)
	if err != nil {
		errs = append(errs, fmt.Errorf("total needed must be a number"))
	}
	if totalNeededP <= 0 {
		errs = append(errs, fmt.Errorf("total needed must be greater than 0"))
	}

	currentSavingsP, err := strconv.ParseFloat(currentSavings, 64)
	if err != nil {
		errs = append(errs, fmt.Errorf("current savings must be a number"))
	}
	if currentSavingsP <= 0 {
		errs = append(errs, fmt.Errorf("current savings must be greater than 0"))
	}

	monthlyContributionP, err := strconv.ParseFloat(monthlyContribution, 64)
	if err != nil {
		errs = append(errs, fmt.Errorf("monthly contribution must be a number"))
	}
	if monthlyContributionP <= 0 {
		errs = append(errs, fmt.Errorf("monthly contribution must be greater than 0"))
	}

	monthsToGoalP, err := strconv.Atoi(monthsToGoal)
	if err != nil {
		errs = append(errs, fmt.Errorf("months to goal must be a number"))
	}
	if monthsToGoalP <= 0 {
		errs = append(errs, fmt.Errorf("months to goal must be greater than 0"))
	}

	startDateP, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		errs = append(errs, fmt.Errorf("start date must be in the format DD-MM-YYYY"))
	}
	startDateF := startDateP.Format("02-01-2006")

	endDateP, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		errs = append(errs, fmt.Errorf("end date must be in the format DD-MM-YYYY"))
	}
	endDateF := endDateP.Format("02-01-2006")

	if len(errs) > 0 {
		return Goal{}, fmt.Errorf("validation errors: %v", errs)
	}

	return Goal{
		Name:                name,
		TotalNeeded:         totalNeededP,
		CurrentSavings:      currentSavingsP,
		MonthlyContribution: monthlyContributionP,
		StartDate:           startDateF,
		EndDate:             endDateF,
		MonthsToGoal:        monthsToGoalP,
	}, nil
}

func (g *Goal) CalculateSavings() string {
	g.MonthlyContribution = (g.TotalNeeded - g.CurrentSavings) / (float64(g.MonthsToGoal))

	return fmt.Sprintf("Name: %s\nCurrent Savings: $%.2f\nTotal Needed: $%.2f\nMonthly Contribution: $%.2f\nStart Date: %s\nEnd Date: %s\n",
		g.Name, g.CurrentSavings, g.TotalNeeded, g.MonthlyContribution, g.StartDate, g.EndDate)
}
