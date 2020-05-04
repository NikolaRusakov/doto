package gql

import (
	"context"
	"github.com/google/uuid"
	"github.com/kr/pretty"
	"github.com/mitchellh/mapstructure"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/rusakov/doto/gql/db"
	"log"
	"os"
)

type Resolver struct {
	Goals map[string]*Goal
}

type GoalRes struct {
	id int
	//labels string
	props Goal
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateGoal(ctx context.Context, inputGoal InputGoal) (*Goal, error) {
	var (
		driver  neo4j.Driver
		session neo4j.Session
		goal    *Goal
		err     error
		uid     = uuid.New().String()
	)
	inputGoal.ID = &uid
	goal = new(Goal)
	pretty.Print(inputGoal)
	// Construct a new Driver
	if driver, err = neo4j.NewDriver(
		os.Getenv("NEO_URL"),
		neo4j.BasicAuth(os.Getenv("NEO_NAME"),
			os.Getenv("NEO_PASS"),
			""),
		func(config *neo4j.Config) {
			config.Log = neo4j.ConsoleLogger(neo4j.ERROR)
		}); err != nil {
		log.Fatal(err)
	}
	defer driver.Close()

	// Acquire a Session
	if session, err = driver.Session(neo4j.AccessModeRead); err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer session.Close()
	session, err = db.Run()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	result, err := session.Run(
		`CREATE (n:Goal { 
		ID: $ID,
		Section: $Section,
		Name: $Name,
		Description: $Description,
		Estimation: $Estimation,
		IsActive: $IsActive,
		Timestamp: $Timestamp,
		Tasks: $Tasks,
		ConnectionWith: $ConnectionWith
		}) RETURN properties(n), ID(n) `,
		map[string]interface{}{
			"ID":             inputGoal.ID,
			"Section":        inputGoal.Section,
			"Name":           inputGoal.Name,
			"Description":    inputGoal.Description,
			"Estimation":     inputGoal.Estimation,
			"IsActive":       inputGoal.IsActive,
			"Timestamp":      inputGoal.Timestamp,
			"Tasks":          inputGoal.Tasks,
			"ConnectionWith": inputGoal.ConnectionWith,
		})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for result.Next() {
		err = mapstructure.Decode(result.Record().GetByIndex(0), goal)
	}

	if err = result.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return goal, err
}
func (r *mutationResolver) CreateTask(ctx context.Context, inputTask InputTask) (*Task, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateGoal(ctx context.Context, inputGoal InputGoal) (*Goal, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateTask(ctx context.Context, inputTask InputTask) (*Task, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GoalByID(ctx context.Context, id string) (*Goal, error) {
	var (
		goal    *Goal
		err     error
		session neo4j.Session
	)

	session, err = db.Run()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	result, err := session.Run(
		`Match (n:Goal { ID: $ID }))<-[:IS_TASK_OF]-(task)
		CREATE (task: Task { })
RETURN properties(n), task, subtask)`,
		map[string]interface{}{
			"ID": id,
		})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for result.Next() {
		err = mapstructure.Decode(result.Record().GetByIndex(0), goal)
	}

	if err = result.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return goal, err
}
func (r *queryResolver) GoalsFromSection(ctx context.Context, section Section) ([]*Goal, error) {
	_, _ = pretty.Print(section)
	panic("not implemented")
}
func (r *queryResolver) GoalsInBatch(ctx context.Context, arg SearchArgs) ([]*Goal, error) {
	panic("not implemented")
}
func (r *queryResolver) TaskDetailsFromGoal(ctx context.Context, task InputTask) (*Task, error) {
	panic("not implemented")
}

func NewResolver() Config {
	return Config{
		Resolvers: &Resolver{
			Goals: map[string]*Goal{},
		},
	}
}
