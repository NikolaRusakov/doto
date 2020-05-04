package goal_impl

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/jmcvetta/neoism"
	"github.com/mitchellh/mapstructure"
	"github.com/rusakov/doto/server/api/models"
	"github.com/rusakov/doto/server/api/restapi/operations/goal"
	"github.com/rusakov/doto/server/neo"
)

func GoalsSection(params goal.GoalsSectionParams) middleware.Responder {
	s := params.Section
	println(s)
	res0 := []struct {
		N neoism.Node // Column "n" gets automagically unmarshalled into field N
	}{}

	//s := strings.Split(sections[0], ",")
	println("multi")
	for i, v := range s {
		println(i, v)
	}
	cq1 := neoism.CypherQuery{
		Statement:  `MATCH (n:GOAL) WHERE n.section IN {sections} RETURN n`,
		Parameters: neoism.Props{"sections": &s},
		Result:     &res0,
	}
	err := neo.Db.Cypher(&cq1)

	//output, err := json.Marshal(res0)
	if err != nil {
		//http.Error(w, err.Error(), 500)
		//return
	}
	goals := make([]*models.Goal, len(res0))
	for k, v := range res0 {
		println(v.N.Data)
		var g *models.Goal
		err := mapstructure.Decode(v.N.Data, &g)
		if err != nil {
			//http.Error(w, err.Error(), 500)
			//return
		}
		goals[k] = g
	}

	//conv, _ := json.Marshal(goals)
	return goal.NewGoalsSectionOK().WithPayload(goals)
}