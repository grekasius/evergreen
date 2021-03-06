// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"fmt"
	"io"
	"strconv"

	"github.com/evergreen-ci/evergreen/rest/model"
)

type GroupedFiles struct {
	TaskName *string          `json:"taskName"`
	Files    []*model.APIFile `json:"files"`
}

type GroupedProjects struct {
	Name     string                   `json:"name"`
	Projects []*model.UIProjectFields `json:"projects"`
}

type PatchDuration struct {
	Makespan  *string    `json:"makespan"`
	TimeTaken *string    `json:"timeTaken"`
	Time      *PatchTime `json:"time"`
}

type PatchTime struct {
	Started     *string `json:"started"`
	Finished    *string `json:"finished"`
	SubmittedAt string  `json:"submittedAt"`
}

type Projects struct {
	Favorites     []*model.UIProjectFields `json:"favorites"`
	OtherProjects []*GroupedProjects       `json:"otherProjects"`
}

type TaskResult struct {
	ID           string `json:"id"`
	DisplayName  string `json:"displayName"`
	Version      string `json:"version"`
	Status       string `json:"status"`
	BaseStatus   string `json:"baseStatus"`
	BuildVariant string `json:"buildVariant"`
}

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "ASC"
	SortDirectionDesc SortDirection = "DESC"
)

var AllSortDirection = []SortDirection{
	SortDirectionAsc,
	SortDirectionDesc,
}

func (e SortDirection) IsValid() bool {
	switch e {
	case SortDirectionAsc, SortDirectionDesc:
		return true
	}
	return false
}

func (e SortDirection) String() string {
	return string(e)
}

func (e *SortDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortDirection", str)
	}
	return nil
}

func (e SortDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TaskSortCategory string

const (
	TaskSortCategoryName       TaskSortCategory = "NAME"
	TaskSortCategoryStatus     TaskSortCategory = "STATUS"
	TaskSortCategoryBaseStatus TaskSortCategory = "BASE_STATUS"
	TaskSortCategoryVariant    TaskSortCategory = "VARIANT"
)

var AllTaskSortCategory = []TaskSortCategory{
	TaskSortCategoryName,
	TaskSortCategoryStatus,
	TaskSortCategoryBaseStatus,
	TaskSortCategoryVariant,
}

func (e TaskSortCategory) IsValid() bool {
	switch e {
	case TaskSortCategoryName, TaskSortCategoryStatus, TaskSortCategoryBaseStatus, TaskSortCategoryVariant:
		return true
	}
	return false
}

func (e TaskSortCategory) String() string {
	return string(e)
}

func (e *TaskSortCategory) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskSortCategory(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskSortCategory", str)
	}
	return nil
}

func (e TaskSortCategory) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TestSortCategory string

const (
	TestSortCategoryStatus   TestSortCategory = "STATUS"
	TestSortCategoryDuration TestSortCategory = "DURATION"
	TestSortCategoryTestName TestSortCategory = "TEST_NAME"
)

var AllTestSortCategory = []TestSortCategory{
	TestSortCategoryStatus,
	TestSortCategoryDuration,
	TestSortCategoryTestName,
}

func (e TestSortCategory) IsValid() bool {
	switch e {
	case TestSortCategoryStatus, TestSortCategoryDuration, TestSortCategoryTestName:
		return true
	}
	return false
}

func (e TestSortCategory) String() string {
	return string(e)
}

func (e *TestSortCategory) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TestSortCategory(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TestSortCategory", str)
	}
	return nil
}

func (e TestSortCategory) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
