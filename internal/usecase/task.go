package usecase

import (
	"context"
	"fmt"
)

type TaskSolver interface {
	SolveTask1(nums []int) (int, error)
	SolveTask2(str1, str2 string) bool
	SolveTask3(s string, numRows int) (string, error)
}

type TaskUseCase struct {
	taskSolverLib TaskSolver
}

func NewTaskUseCase(taskSolverLib TaskSolver) *TaskUseCase {
	return &TaskUseCase{taskSolverLib}
}

func (uc *TaskUseCase) SolveTask1(ctx context.Context, nums []int) (int, error) {
	v, err := uc.taskSolverLib.SolveTask1(nums)

	if err != nil {
		return 0, fmt.Errorf("solver lib error: %w", err)
	}

	return v, nil
}

func (uc *TaskUseCase) SolveTask2(ctx context.Context, str1, str2 string) (bool, error) {
	v := uc.taskSolverLib.SolveTask2(str1, str2)

	return v, nil
}

func (uc *TaskUseCase) SolveTask3(ctx context.Context, s string, numRows int) (string, error) {
	v, err := uc.taskSolverLib.SolveTask3(s, numRows)

	if err != nil {
		return "", fmt.Errorf("solver lib error: %w", err)
	}

	return v, nil
}
