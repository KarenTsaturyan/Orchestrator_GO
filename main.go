package main

import (
	"fmt"
	"orchestrator/task"
	"orchestrator/worker"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

// create task implementation
// func createContainer() (*task.Docker, *task.DockerResult) {
// 	c := task.Config{
// 		Name:  "test-container-1",
// 		Image: "postgres:13",
// 		Env: []string{
// 			"POSTGRES_USER=cube",
// 			"POSTGRES_PASSWORD=secret",
// 		},
// 	}

// 	dc, _ := client.NewClientWithOpts(client.FromEnv)
// 	d := task.Docker{
// 		Client: dc,
// 		Config: c,
// 	}

// 	result := d.Run()
// 	if result.Error != nil {
// 		fmt.Printf("%v\n", result.Error)
// 		return nil, nil
// 	}

// 	fmt.Printf("Container %s is running with config %v\n", result.ContainerId, c)
// 	return &d, &result
// }

// func stopContainer(d *task.Docker, id string) *task.DockerResult {
// 	result := d.Stop(id)
// 	if result.Error != nil {
// 		fmt.Printf("%v\n", result.Error)
// 		return nil
// 	}

// 	fmt.Printf("Container %s has been stopped and removed\n", result.ContainerId)
// 	return &result
// }

func main() {
	db := make(map[uuid.UUID]*task.Task)
	w1 := worker.Worker{
		Queue: *queue.New(),
		Db:    db,
	}

	// w2 := worker.Worker{
	// 	Queue: *queue.New(),
	// 	Db:    db,
	// }

	// w3 := worker.Worker{
	// 	Queue: *queue.New(),
	// 	Db:    db,
	// }

	t1 := task.Task{
		ID:    uuid.New(),
		Name:  "test-container-1",
		State: task.Scheduled,
		Image: "strm/helloworld-http",
	}

	// t2 := task.Task{
	// 	ID:    uuid.New(),
	// 	Name:  "test-container-1",
	// 	State: task.Scheduled,
	// 	Image: "strm/helloworld-http",
	// }

	// t3 := task.Task{
	// 	ID:    uuid.New(),
	// 	Name:  "test-container-1",
	// 	State: task.Scheduled,
	// 	Image: "strm/helloworld-http",
	// }

	// first time the worker will see the task
	fmt.Println("starting task")
	w1.AddTask(t1)
	result := w1.RunTask()
	if result.Error != nil {
		panic(result.Error)
	}
	t1.ContainerID = result.ContainerId
	fmt.Printf("task %s is running in container %s\n", t1.ID, t1.ContainerID)
	fmt.Println("Sleepy time")
	time.Sleep(time.Second * 30)
	fmt.Printf("stopping task %s\n", t1.ID)
	t1.State = task.Completed
	w1.AddTask(t1)
	result = w1.RunTask()
	if result.Error != nil {
		panic(result.Error)
	}
}
