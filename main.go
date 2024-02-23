package main

import (
	"net/http"
    "errors"

	"github.com/gin-gonic/gin"
)

func getTodos(context *gin.Context){
     context.IndentedJSON(http.StatusOK,todos)
}

func addTodos(context *gin.Context){
	var newTodos todo

	if err := context.BindJSON(&newTodos); err !=nil {
		return
	}
    todos = append(todos, newTodos)
	context.IndentedJSON(http.StatusCreated,newTodos)
}
func getTodo(context *gin.Context){
	id := context.Param("id")
	
	todo, err := getTodoById(id); 
	if err != nil {
        context.IndentedJSON(http.StatusNotFound, gin.H{"message": "To do not found"})
		return
	}

	context.IndentedJSON(http.StatusOK,todo)
}

func toggleStatusOfTodo (context *gin.Context){
	id := context.Param("id")
	todo,err := getTodoById(id)

	if err != nil{
		context.IndentedJSON(http.StatusNotFound,gin.H{"message": "To do not found"})
		return
	}

	todo.Complated =!todo.Complated

	context.IndentedJSON(http.StatusOK,todo)
}

func getTodoById(id string) (*todo , error ){
for i ,t := range todos{
	if t.ID == id {
		return &todos[i], nil
	}
}

return nil,errors.New("todo not found")
}

func main(){
	router := gin.Default()

     router.GET("/todos",getTodos)
	 
     router.GET("/todos/:id",getTodo)
	 router.PATCH("/todos/:id",toggleStatusOfTodo)
	 router.POST("/todos",addTodos)
	router.Run(":8080")
}

type todo struct{
	ID          string `json:"id"`
	Item        string `json:"item"`
	Complated   bool `json:"complated"`
}

var todos = []todo{
	{ID:"1",Item:"eat breack fast",Complated: false},
	{ID:"2",Item:"call mom",Complated: false},
	{ID:"3",Item:"learn go",Complated: false},
}