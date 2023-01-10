package main

import (
	"fortune3-pickaxe/controllers"
	"fortune3-pickaxe/models"

	"fmt"
)

func main() {
	goal := &models.Goal{
		Name:           "Down Payment on House",
		TotalNeeded:    100000.00,
		CurrentSavings: 25000.00,
		MonthsToGoal:   24,
	}

	g := controllers.NewGoalController(goal)
	g.CalculateSavings()

	fmt.Printf("Name: %s\nTotal Needed: $%.2f\nCurrent Savings: $%.2f\nMonthly Contribution: $%.2f\nStart Date: %s\nEnd Date: %s\n",
		goal.Name, goal.TotalNeeded, goal.CurrentSavings, goal.MonthlyContribution, goal.StartDate.Format("2006-01-02"), goal.EndDate.Format("2006-01-02"))
}
