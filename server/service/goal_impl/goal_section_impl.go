package goal_impl

import (
	//"encoding/json"
	"fmt"
	"github.com/go-openapi/runtime/middleware"

	//"github.com/google/uuid"
	//"github.com/rusakov/doto/server/api/models"
	"github.com/rusakov/doto/server/api/restapi/operations/goal"
	"github.com/rusakov/doto/server/neo"
)

func GoalPost(params goal.GoalPost) middleware.Responder {
	//var goal *models.Goal
	//goal.ID = uuid.New().String()
//params.Handler.Handle()
	//e := json.NewDecoder(params).Decode(&goal)
	//if e != nil {
	//	http.Error(w, e.Error(), 500)
	//	return
	//}

	var inode map[string]interface{}
	//unmarsh, _ := json.Marshal(goal)
	//
	//er := json.Unmarshal(unmarsh, &inode)
	//if er != nil {
	//	http.Error(params..Error(), 500)
	//}

	for field, val := range inode {
		fmt.Println("KV Pair: ", field, " : ", val)
	}

	n, err := neo.Db.CreateNode(inode)
	if err != nil {
		//http.Error(w, err.Error(), 500)
		//return err.Error()
		println(err.Error())
	}
	n.AddLabel("GOAL")

	//output, err := json.Marshal(goal)
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	return goal.().WithPayload(goals)
}