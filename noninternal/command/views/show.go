package views

import (
	"fmt"

	"github.com/kevinklinger/open_terraform/noninternal/command/arguments"
	"github.com/kevinklinger/open_terraform/noninternal/command/format"
	"github.com/kevinklinger/open_terraform/noninternal/command/jsonplan"
	"github.com/kevinklinger/open_terraform/noninternal/command/jsonstate"
	"github.com/kevinklinger/open_terraform/noninternal/configs"
	"github.com/kevinklinger/open_terraform/noninternal/plans"
	"github.com/kevinklinger/open_terraform/noninternal/states/statefile"
	"github.com/kevinklinger/open_terraform/noninternal/terraform"
	"github.com/kevinklinger/open_terraform/noninternal/tfdiags"
)

type Show interface {
	// Display renders the plan, if it is available. If plan is nil, it renders the statefile.
	Display(config *configs.Config, plan *plans.Plan, stateFile *statefile.File, schemas *terraform.Schemas) int

	// Diagnostics renders early diagnostics, resulting from argument parsing.
	Diagnostics(diags tfdiags.Diagnostics)
}

func NewShow(vt arguments.ViewType, view *View) Show {
	switch vt {
	case arguments.ViewJSON:
		return &ShowJSON{view: view}
	case arguments.ViewHuman:
		return &ShowHuman{view: view}
	default:
		panic(fmt.Sprintf("unknown view type %v", vt))
	}
}

type ShowHuman struct {
	view *View
}

var _ Show = (*ShowHuman)(nil)

func (v *ShowHuman) Display(config *configs.Config, plan *plans.Plan, stateFile *statefile.File, schemas *terraform.Schemas) int {
	if plan != nil {
		renderPlan(plan, schemas, v.view)
	} else {
		if stateFile == nil {
			v.view.streams.Println("No state.")
			return 0
		}

		v.view.streams.Println(format.State(&format.StateOpts{
			State:   stateFile.State,
			Color:   v.view.colorize,
			Schemas: schemas,
		}))
	}
	return 0
}

func (v *ShowHuman) Diagnostics(diags tfdiags.Diagnostics) {
	v.view.Diagnostics(diags)
}

type ShowJSON struct {
	view *View
}

var _ Show = (*ShowJSON)(nil)

func (v *ShowJSON) Display(config *configs.Config, plan *plans.Plan, stateFile *statefile.File, schemas *terraform.Schemas) int {
	if plan != nil {
		jsonPlan, err := jsonplan.Marshal(config, plan, stateFile, schemas)

		if err != nil {
			v.view.streams.Eprintf("Failed to marshal plan to json: %s", err)
			return 1
		}
		v.view.streams.Println(string(jsonPlan))
	} else {
		// It is possible that there is neither state nor a plan.
		// That's ok, we'll just return an empty object.
		jsonState, err := jsonstate.Marshal(stateFile, schemas)
		if err != nil {
			v.view.streams.Eprintf("Failed to marshal state to json: %s", err)
			return 1
		}
		v.view.streams.Println(string(jsonState))
	}
	return 0
}

// Diagnostics should only be called if show cannot be executed.
// In this case, we choose to render human-readable diagnostic output,
// primarily for backwards compatibility.
func (v *ShowJSON) Diagnostics(diags tfdiags.Diagnostics) {
	v.view.Diagnostics(diags)
}
