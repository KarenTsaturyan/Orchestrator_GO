package main

import (
	"fmt"
	"orchestrator/manager"
	"orchestrator/task"
	"orchestrator/worker"
	"os"
	"strconv"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

func main() {
	// API_WORKER_HOST=localhost API_WORKER_PORT=5555 API_MANAGER_HOST=localhost API_MANAGER_PORT=5556 go run main.g
	whost := os.Getenv("API_WORKER_HOST")
	wport, _ := strconv.Atoi(os.Getenv("API_WORKER_PORT"))

	mhost := os.Getenv("API_MANAGER_HOST")
	mport, _ := strconv.Atoi(os.Getenv("API_MANAGER_PORT"))

	fmt.Println("Starting Orchestrator worker")

	w := worker.Worker{
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]*task.Task),
	}
	wapi := worker.Api{Address: whost, Port: wport, Worker: &w}

	go w.RunTasks()
	go w.CollectStats()
	go wapi.Start()

	fmt.Println("Starting Orchestrator manager")

	workers := []string{fmt.Sprintf("%s:%d", whost, wport)}
	m := manager.New(workers)
	mapi := manager.Api{Address: mhost, Port: mport, Manager: m}

	go m.ProcessTasks()
	go m.UpdateTasks()

	mapi.Start()

}
