package controllers

import (
	"fortune3-pickaxe/models"
	"html/template"
)

type GoalController struct {
	GoalView  template.Template
	GoalModel *models.Goal
}

func NewGoalController(gm *models.Goal) *GoalController {
	return &GoalController{
		GoalView:  template.Template{}, //views.Must(views.ParseFS(templates.FS, "templates/goal.gohtml")),
		GoalModel: gm,
	}
}

// func (gc *GoalController) CalculateSavings(w http.ResponseWriter, r *http.Request) {
func (gc *GoalController) CalculateSavings() {
	gc.GoalModel.CalculateSavings()
	// get correct data and supply to template
	/* var data struct {
		...
	}
	data. = r.FormValue("...") */
	//gc.GoalView.Execute(w, r, GoalController{})
}
