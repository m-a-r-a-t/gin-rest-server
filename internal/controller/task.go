package controller

import (
	"context"
	"fmt"
	resp "gin-rest-server/internal/lib/api/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

type TaskSolver interface {
	SolveTask1(ctx context.Context, nums []int) (int, error)
	SolveTask2(ctx context.Context, str1, str2 string) (bool, error)
	SolveTask3(ctx context.Context, s string, numRows int) (string, error)
}

type TasksController struct {
	log          *slog.Logger
	taskSolverUc TaskSolver
}

func NewTaskController(log *slog.Logger, taskSolver TaskSolver) *TasksController {
	return &TasksController{log, taskSolver}

}

func (tc *TasksController) Register(router *gin.RouterGroup) {
	r := router.Group("/tasks")
	r.POST("/task1", tc.SolveTask1)
	r.POST("/task2", tc.SolveTask2)
	r.POST("/task3", tc.SolveTask3)
}

type Task1Input struct {
	Nums []int `json:"nums" binding:"required,min=2"`
}

type Task1Response struct {
	Value int `json:"value"`
}

func (tc *TasksController) SolveTask1(c *gin.Context) {
	var inputData Task1Input

	if err := c.ShouldBindJSON(&inputData); err != nil {
		validateErr := fmt.Errorf("validation error: %w", err)
		tc.log.Error(validateErr.Error())

		c.JSON(http.StatusBadRequest, resp.Error(validateErr.Error()))
		return
	}

	tc.log.Debug(fmt.Sprintf("input nums: %v", inputData.Nums))

	value, err := tc.taskSolverUc.SolveTask1(c.Request.Context(), inputData.Nums)

	if err != nil {
		tc.log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	tc.log.Debug(fmt.Sprintf("result: %d", value))

	c.JSON(http.StatusOK, Task1Response{Value: value})
}

type Task2Input struct {
	Str1 string `json:"str1" binding:"required"`
	Str2 string `json:"str2" binding:"required"`
}

type Task2Response struct {
	Value bool `json:"value"`
}

func (tc *TasksController) SolveTask2(c *gin.Context) {
	var inputData Task2Input

	if err := c.ShouldBindJSON(&inputData); err != nil {
		validateErr := fmt.Errorf("validation error: %w", err)
		tc.log.Error(validateErr.Error())

		c.JSON(http.StatusBadRequest, resp.Error(validateErr.Error()))
		return
	}

	tc.log.Debug(fmt.Sprintf("input strings: %s , %s", inputData.Str1, inputData.Str2))

	value, err := tc.taskSolverUc.SolveTask2(c.Request.Context(), inputData.Str1, inputData.Str2)

	if err != nil {
		tc.log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	tc.log.Debug(fmt.Sprintf("result: %v", value))

	c.JSON(http.StatusOK, Task2Response{Value: value})
}

type Task3Input struct {
	Str     string `json:"str" binding:"required,min=1"`
	NumRows int    `json:"num_rows" binding:"required,min=1"`
}

type Task3Response struct {
	Value string `json:"value"`
}

func (tc *TasksController) SolveTask3(c *gin.Context) {
	var inputData Task3Input

	if err := c.ShouldBindJSON(&inputData); err != nil {
		validateErr := fmt.Errorf("validation error: %w", err)
		tc.log.Error(validateErr.Error())

		c.JSON(http.StatusBadRequest, resp.Error(validateErr.Error()))
		return
	}

	tc.log.Debug(fmt.Sprintf("input str and num rows: %s , %d", inputData.Str, inputData.NumRows))

	value, err := tc.taskSolverUc.SolveTask3(c.Request.Context(), inputData.Str, inputData.NumRows)

	if err != nil {
		tc.log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	tc.log.Debug(fmt.Sprintf("result: %v", value))

	c.JSON(http.StatusOK, Task3Response{Value: value})
}
