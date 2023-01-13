package controllers

import (
	"fmt"
	"fortune3-pickaxe/models"
	"net/http"
	"strconv"
	"time"
)

type Goals struct {
	Templates struct {
		New Template
	}
	Goal models.Goal
}

func (g Goals) New(w http.ResponseWriter, r *http.Request) {
	g.Templates.New.Execute(w, r, nil)
}

func (g Goals) Create(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	totalNeeed, _ := strconv.ParseFloat(r.FormValue("total-needed"), 64)
	currentSavings, _ := strconv.ParseFloat(r.FormValue("current-savings"), 64)
	monthlyContribution, _ := strconv.ParseFloat(r.FormValue("monthly-contribution"), 64)
	monthsToGoal, _ := strconv.ParseInt(r.FormValue("months-to-goal"), 0, 64)

	startDate, _ := time.Parse("DD-MM-YYYY", r.FormValue("start-date"))
	endDate, _ := time.Parse("DD-MM-YYYY", r.FormValue("end-date"))

	goal, err := g.Goal.Create(
		name,
		totalNeeed,
		currentSavings,
		monthlyContribution,
		monthsToGoal,
		startDate,
		endDate,
	)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Goal created: %+v\n", goal)
	http.Redirect(w, r, "/", http.StatusFound)
}
