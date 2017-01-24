package servicecontext

import (
	"fmt"
	"net/http"

	"github.com/evergreen-ci/evergreen/apiv3"
	"github.com/evergreen-ci/evergreen/model/task"
)

// TaskConnector is an interface which abstracts the service layer's fetching
// and modifactions of tasks.
type TaskConnector interface {

	// FindTaskById is a method to find a specific task given its ID.
	FindTaskById(string) (*task.Task, error)

	// FindTasksByBuildId is a method to find a set of tasks which all have the same
	// BuildId. It takes the buildId being queried for as its parameter and
	// returns a list of tasks which match.
	FindTasksByBuildId(string, string, int) ([]task.Task, error)
}

// DBTaskConnector is a struct that implements the TaskConnector interface
// through interactions with the backing database.
type DBTaskConnector struct{}

// FindTaskById uses the service layer's task type to query the backing database for
// the task with the given taskId.
func (tc *DBTaskConnector) FindTaskById(taskId string) (*task.Task, error) {
	t, err := task.FindOne(task.ById(taskId))
	if err != nil {
		return nil, err
	}
	if t == nil {
		return nil, &apiv3.APIError{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("task with id %s not found", taskId),
		}
	}
	return t, nil
}

// FindTasksByBuildId uses the service layer's task type to query the backing database for a
// list of task that matches buildId. It accepts the startTaskId and a limit
// to allow for pagination of the queries. It returns results sorted by taskId.
func (tc *DBTaskConnector) FindTasksByBuildId(buildId, startTaskId string, limit int) ([]task.Task, error) {
	var ts []task.Task
	var err error
	// If we have specified a taskId to start the iteration from, then search
	// for it and fetch the list of tasks starting there.
	if startTaskId != "" {
		ts, err = task.Find(task.ByBuildIdAfterTaskId(buildId, startTaskId).Limit(limit))
		if err != nil {
			return nil, err
		}
		// Otherwise, begin the iteration from the beginning of the list of tasks.
	} else {
		ts, err = task.Find(task.ByBuildId(buildId).Sort([]string{"+" + task.IdKey}).Limit(limit))
		if err != nil {
			return nil, err
		}
	}
	if len(ts) == 0 {
		return nil, &apiv3.APIError{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("tasks with buildId %v not found", buildId),
		}
	}
	return ts, nil
}

// MockTaskConnector stores a cached set of tasks that are queried against by the
// implementations of the TaskConnector interface's functions.
type MockTaskConnector struct {
	CachedTasks []task.Task
}

// FindTaskById provides a mock implementation of a TaskConnector that does not
// need to use a database. It returns results based on the cached tasks
// in the MockTaskConnector.
func (mdf *MockTaskConnector) FindTaskById(taskId string) (*task.Task, error) {
	return &mdf.CachedTasks[0], nil
}

// FindTasksByBuildId provides a mock implementation of a TaskConnector that does not
// need to use a database. It returns results based on the cached tasks
// in the MockTaskConnector.
func (mdf *MockTaskConnector) FindTasksByBuildId(buildId, startTaskId string, limit int) ([]task.Task, error) {
	return mdf.CachedTasks, nil
}
