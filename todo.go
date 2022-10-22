package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Item struct {
	Task        string
	Description string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type TodoList []Item

func (t *TodoList) AddTask(task, desc string) {
	it := Item{
		Task:        task,
		Description: desc,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, it)
}

func (t *TodoList) Complete(i int) error {
	ls := *t
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()
	return nil
}

func (t *TodoList) Delete(i int) error {
	ls := *t
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	*t = append(ls[:i-1], ls[i:]...)
	return nil
}

func (t *TodoList) Save(filename string) error {
	js, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, js, 0644)

}

func (t *TodoList) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, t)
}
