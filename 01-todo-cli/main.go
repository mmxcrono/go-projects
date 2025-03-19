package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type taskAction string

// Define constants
const (
	AddAction    taskAction = "add"
	DeleteAction taskAction = "delete"
	ListAction   taskAction = "list"
	UpdateAction taskAction = "update"
)

// Store constants in a slice
var taskActions = []taskAction{AddAction, DeleteAction, ListAction, UpdateAction}

type Task struct {
	Id            uint32
	Desc          string
	TaskStartDate string
	TaskStartTime string
	TaskEndDate   string
	TaskEndTime   string
}

type DatabaseTable string

const (
	TaskTable DatabaseTable = "task"
)

var DatabaseNextIds = map[DatabaseTable]uint32{
	TaskTable: 2,
}

var Database = map[uint32]Task{
	1: {
		Id:   1,
		Desc: "Existing Task",
	},
}

var (
	availableActionsCache string
	once                  sync.Once
	dbMutex               sync.RWMutex
	cacheMutex            sync.RWMutex
)

func main() {
	seedDatabase()
	availableActions := getAvailableActions()

	// Clean up initial app name from args
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: tasks action args")
		return
	}

	inputAction := args[0]

	switch inputAction {
	case "list":
		fmt.Println("List Tasks")

		printTasks()
	case "delete":
		fmt.Println("Delete Task")

		if len(args) <= 1 {
			fmt.Printf("Usage: tasks delete TASK_ID_NUMBER\n")
			return
		}

		taskId, ok := convertTaskId(args[1])

		if !ok {
			fmt.Println("Error: Invalid task ID. Must be a positive integer.")
			return
		}

		_, ok = getTask(taskId)

		if !ok {
			fmt.Printf("Task ID %v does not exist\n", taskId)
			printTasks()
			return
		}
		deleteTask(taskId)
		printTasks()
	case "update":
		fmt.Println("Update Task")

		if len(args) <= 2 {
			fmt.Printf("Usage: tasks update TASK_ID_NUMBER NEW_DESCRIPTION\n")
			return
		}

		taskId, ok := convertTaskId(args[1])

		if !ok {
			fmt.Println("Error: Invalid task ID. Must be a positive integer.")
			return
		}

		task, ok := getTask(taskId)

		if !ok {
			fmt.Printf("Task ID %v does not exist\n", taskId)
			printTasks()
			return
		}
		newDesc := args[2]
		task.Desc = newDesc
		updateTask(taskId, &task)
		printTasks()
	case "add":
		if len(args) <= 1 {
			fmt.Printf("Usage: tasks add \"Your todo description\"\n")
			return
		}
		description := args[1]

		newTask := Task{
			Desc: description,
		}

		// TODO Handle start/end dates
		// TODO Handle start/end times
		// Enter date - default to today
		// All Day? Y/N
		// Enter time - default to now
		// Enter end date
		// Enter end time
		// Enter time zone - default to PST

		insertTask(&newTask)
		printTasks()
	default:
		fmt.Printf("Unsupported action specified: %v\n", inputAction)
		fmt.Printf("Supported actions: %v\n", availableActions)
	}
}

func seedDatabase() {
	var wg sync.WaitGroup
	// Insert some tasks
	wg.Add(1)
	go func() {
		defer wg.Done()
		insertTask(&Task{
			Desc: "Buy groceries",
		})
		insertTask(&Task{
			Desc: "Complete report",
		})
	}()

	// Read tasks concurrently
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 50) // Simulate processing delay
			task, exists := getTask(uint32(i))
			if exists {
				fmt.Printf("Read Task %d: %s\n", task.Id, task.Desc)
			} else {
				fmt.Printf("Task %d not found.\n", i)
			}
		}(i)
	}

	// Delete a task concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 100) // Simulate delay
		deleteTask(1)
	}()

	wg.Wait()
}

func getAvailableActions() string {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if availableActionsCache != "" {
		return availableActionsCache
	}

	once.Do(func() {
		actionStrings := make([]string, len(taskActions))

		for i, action := range taskActions {
			actionStrings[i] = string(action)
		}

		availableActionsCache = strings.Join(actionStrings, " ")
	})
	return availableActionsCache
}

func printTasks() {
	if len(Database) == 0 {
		fmt.Println("Task list is empty")
		return
	}
	fmt.Println("Task list")
	for _, task := range Database {
		fmt.Printf("ID: %v, Description: %v\n", task.Id, task.Desc)
	}
}

func getTask(id uint32) (Task, bool) {
	dbMutex.RLock()
	defer dbMutex.RUnlock()

	task, exists := Database[id]
	return task, exists
}

func insertTask(task *Task) {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	nextId := DatabaseNextIds[TaskTable]
	task.Id = nextId
	Database[nextId] = *task

	DatabaseNextIds[TaskTable]++
}

func deleteTask(id uint32) {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	if _, exists := Database[id]; exists {
		delete(Database, id)
		fmt.Printf("Task %d deleted.\n", id)
	} else {
		fmt.Printf("Task %d not found.\n", id)
	}
}

func updateTask(id uint32, task *Task) {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	if _, exists := Database[id]; exists {
		Database[id] = *task
		fmt.Printf("Task %d deleted.\n", id)
	} else {
		fmt.Printf("Task %d not found.\n", id)
	}
}

func convertTaskId(input string) (uint32, bool) {
	taskId, err := strconv.ParseUint(input, 10, 32)

	if err != nil {
		return 0, false
	}

	return uint32(taskId), true
}
