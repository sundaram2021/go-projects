package main 

import (
	"fmt"
)

type todo struct {
	id int32
	name string
	isDone bool
}

type Todo interface {
	createTodo(t todo) todo
	getTodo(id int32) todo
	deleteTodo(id int32) todo
	getAllTodos() []todo
	updateTodo(t todo) todo
}

type Todos struct {
	todos []todo
}


func (t *Todos) createTodo(td todo) todo {
	t.todos = append(t.todos, td)
	return td
}

func (t *Todos) deleteTodo(id int32) todo {
	var delete todo
	for i, ele := range t.todos{
		if ele.id == id {
			delete = ele
			t.todos = append(t.todos[:i],t.todos[i+1:]... )
			break
		}
	}
	return delete
}

func (t *Todos) getTodo(id int32) todo{
	var get todo
	for _, ele := range t.todos{
		if ele.id == id {
			get = ele
			// t.todos = append(t.todos[:i],t.todos[i+1:]... )
			break
		}
	}
	return get
}


func (t *Todos) getAllTodos() []todo{
	var allTodos []todo
	allTodos = append(allTodos, t.todos...)

	return allTodos
}


func (t *Todos) updateTodo(td todo) todo {
	var updatedTodo todo 

	for i, ele := range t.todos {
		if ele.id == td.id {
			t.todos[i] = td
			updatedTodo = td
			// t.todos = append(t.todos[:i],t.todos[i+1:]... )
			break
		}
	}


	return updatedTodo
}


func main() {
	tds := &Todos{}

	tds.createTodo(todo{id: 1, name: "learn go", isDone: false})
	tds.createTodo(todo{id: 2, name: "get some water", isDone: true})
	tds.createTodo(todo{id: 3, name: "get you work done", isDone: false})
	tds.createTodo(todo{id: 4, name: "get some sleep", isDone: true})
	tds.createTodo(todo{id: 5, name: "learn to do", isDone: false})

	all := tds.getAllTodos()
	fmt.Println("all todos: ", all)

	id3 := tds.getTodo(3)
	fmt.Println("Id3 todo : ", id3)

	dl3 := tds.deleteTodo(5)
	fmt.Println("Deleted todo: ", dl3)

	upd := tds.updateTodo(todo{id: 4, name: "get some tight sleep", isDone: false})
	fmt.Println("updated todo : ", upd)
}