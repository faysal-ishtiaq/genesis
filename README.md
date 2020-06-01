# Genesis

An opinionated API scaffolding tool using gorm and gin-gonic

## Get started

go get https://github.com/faysal-ishtiaq/genesis

- mkdir appName
- cd appName
- genesis init -n appName -e sqlite
- go generate # to install dependencies and intializing go module
- genesis add -s serviceName

Then add serviceName.InitDB() in initDB() function of main.go, add serviceName.BindRoutes() in main function in main.go, change model.go, dto.go, api.go the way you want it!

### Example
Let's build a test app with Todo service

- mkdir test && cd test
- cd test
- genesis init -n test -e sqlite
- go generate
- genesis add -s todo -m task

Now, we have to add our service in main.go

in main.go/initDB(), write:
```go
todo.InitDB(db)
```
before `return db` statement.

in main.go/main(), write:
```go
todo.BindRoutes(db, router)
```
before 
```go
	err := router.Run()
	if err != nil {
		panic(err)
    }
```
    
Now, our `Task` model in our `todo` service might have a task tile and status to show if it's done.

modify `todo/model.go` like following:

```go
package todo

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Title string
	Done  bool
}
```
<b>Note: don't panic!</b> you just have to write `Title string` and	`Done  bool`

Update `todo/dto.go` like following

```go
package todo

type TaskDTO struct {
	ID    uint `json:"id,string,omitempty"`
	Title string `json:"title"`
	Done  bool `json:"done"`
}
```

Then update ToTask and ToTaskDTO function in `todo/mapper.go` like follwoing
```go
func ToTask(taskDTO TaskDTO) Task {
	return Task{
		Title: taskDTO.Title,
		Done:  taskDTO.Done,
	}
}

func ToTaskDTO(task Task) TaskDTO {
	return TaskDTO{
		Title: task.Title,
		Done:  task.Done,
	}
}
```

One last step. In your `todo/api.go` add these lines before `todo.TodoService.SaveTask(task)`:
```go
	task.Title = taskDTO.Title
	task.Done = taskDTO.Done
```

and you're done. Now run `go run main.go` and your API is up!