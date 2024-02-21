package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	tasknumber int
	info       string
	completed  bool
}

type Todoer struct {
	name      string
	buffer    string
	todotasks []Task
}

var currentTd *Todoer = nil
var taskIndex int = 1

func new() *Todoer {
	var name string
	fmt.Printf("New list name: ")
	fmt.Scanf("%s", &name)

	if len(name) > 20 {
		t_error("length of list name exceeds 20 chars.")
	} else if name == "" {
		t_error("no input name provided.")
	}

	buff := fmt.Sprintf("*Todo* %s\n", name)

	return &Todoer{name: name, buffer: buff}
}

func t_error(log string) {
	fmt.Printf("todoer: %s\n", log)
}

func help() {
	fmt.Println(`
# Copyright (c) Aarav Shreshth, 2024

# todoer is an open-source / free software and is
# freely distributable and available under the MIT License.

*** TODOER Help Page ***

todoer is a straightforward tool that enables you to generate a basic to-do list. 
Below is a compilation of potential commands ([<command>]<args..> - (usage)):

[new] <list name> ~ Creates a new todo list.

[add] <task> ~ Creates a new (non completed) task with the <str> as info.	

[list] - Displays the all inserted, pending and completed tasks.

[markdone] <index> ~ Marks the provided task (index) as done.

[markfree] <index> ~ Marks the provided task (index) as free (not completed).

[delete] - Deletes the current todo list.

[remove] <index> ~ Removes a task from the list.

[exit] ~ Exits the program.

[export] ~ Exports the buffer to current path.

[help] ~ Prints the help page.

Note:  Please be aware that commands marked with '~' may require further user input, 
	   although they can also function with arguments.

	`)
}

func (t *Todoer) addTask() {
	var info string

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter task: ")
	info, _ = reader.ReadString('\n')
	info = strings.TrimSpace(info)

	if info == " " || info == "" {
		t_error("no task info provided.")
		return
	}
	tsk := Task{tasknumber: taskIndex, info: info, completed: false}
	t.todotasks = append(t.todotasks, tsk)
	taskIndex++
}

func introAndSomeFun() {
	fmt.Println(`
The Todoer (todo list generator)
Made by Aarav Shreshth with LOVE in India

** Free Palestine! **

Type 'help' for a list of possible commands.
	`)
}

func (t *Todoer) exportTasks() string {
	for i := 0; i < len(t.todotasks); i++ {
		ct := t.todotasks[i]
		state := " "
		if ct.completed {
			state = "X"
		}

		t.buffer += fmt.Sprintf("\n%d. [%v] - %v\n", i+1, state, ct.info)
	}
	return t.buffer
}

func removeElement(slice []Task, index int) []Task {
	// Ensure the index is within bounds
	if index < 0 || index >= len(slice) {
		fmt.Println("Invalid index")
		return slice
	}

	// Remove the element by creating a new slice without it
	slice = append(slice[:index], slice[index+1:]...)

	// Update tasknumbers for tasks after the removed index
	for i := index; i < len(slice); i++ {
		slice[i].tasknumber = i + 1
	}

	return slice
}

func markTask(slice []Task, index int, completed bool) []Task {
	// Ensure the index is within bounds
	if index < 0 || index >= len(slice) {
		fmt.Println("Invalid index")
		return slice
	}

	slice[index].completed = completed
	return slice
}

func main() {
	introAndSomeFun()

	for {
		var current_arg string
		fmt.Printf("Todoer > ")
		fmt.Scanf("%s", &current_arg)

		// if err != nil {
		// 	t_error("error reading input.")
		// 	continue
		// }

		switch current_arg {
		case "help":
			help()

		case "new":
			if currentTd == nil {
				currentTd = new()
			} else {
				t_error("a list already exists in context, type `delete` to delete the current todo list.")
			}

		case "list":
			if currentTd != nil {
				if len(currentTd.todotasks) != 0 {
					fmt.Printf("\nAll Inserted Tasks:\n")
					for _, ct := range currentTd.todotasks {
						state := "-"
						fmt.Printf("\n%d. [%v] - %v\n", ct.tasknumber, state, ct.info)
					}

					fmt.Printf("\nCompleted tasks:\n")
					for _, ct := range currentTd.todotasks {
						if ct.completed {
							state := "X"
							fmt.Printf("\n%d. [%v] - %v\n", ct.tasknumber, state, ct.info)
						}

					}

					fmt.Printf("\nCurrent Pending Tasks:\n")

					for _, ct := range currentTd.todotasks {
						if !ct.completed {
							fmt.Printf("\n%d. [%v] - %v\n", ct.tasknumber, " ", ct.info)
						}

					}
					fmt.Printf("\n")
				} else {
					fmt.Println("There are no pending tasks.")
				}
			} else {
				t_error("no list initialized.")
			}

		case "delete":
			if currentTd != nil {
				var ans string
				fmt.Printf("Delete current list (Y/n): ")
				_, err := fmt.Scanf("%s", &ans)
				if err != nil {
					t_error("error reading input.")
					continue
				}
				if ans == "y" {
					// Set currentTd to nil to indicate no active list
					currentTd = nil
					taskIndex = 1
				}
			} else {
				t_error("no list initialized.")
			}

		case "add":
			if currentTd != nil {
				currentTd.addTask()
			} else {
				t_error("cannot add a new task; no list initialized.")
			}

		case "remove":
			var index int
			fmt.Printf("Index of the task to delete: ")

			_, err := fmt.Scanf("%d", &index)

			if err != nil {
				t_error("expected integer for index.")
				continue
			}

			var found bool
			for i, ct := range currentTd.todotasks {
				if index == ct.tasknumber {
					currentTd.todotasks = removeElement(currentTd.todotasks, i)
					found = true
					break
				}
			}

			if !found {
				t_error(fmt.Sprintf("\ncould not find any task associated with index (%d)", index))
			}

		case "markdone":
			var index int
			fmt.Printf("Index of the task to mark done: ")

			_, err := fmt.Scanf("%d", &index)

			if err != nil {
				t_error("expected integer for index.")
				continue
			}

			currentTd.todotasks = markTask(currentTd.todotasks, index-1, true)

		case "markfree":
			var index int
			fmt.Printf("Index of the task to mark to free: ")

			_, err := fmt.Scanf("%d", &index)

			if err != nil {
				t_error("expected integer for index.")
				continue
			}

			currentTd.todotasks = markTask(currentTd.todotasks, index-1, false)

		case "exit":
			os.Exit(0)

		case "export":
			if currentTd != nil {
				ex := currentTd.exportTasks()

				file, err := os.Create(currentTd.name + ".todo")
				if err != nil {
					t_error("could not export todo file.")
					continue
				}

				fmt.Fprint(file, ex)
			} else {
				t_error("no list initialized.")
			}

		case "":
			// Do nothing for an empty input

		default:
			t_error(fmt.Sprintf("unknown command `%s`.\ntype `help` for more info.", current_arg))
		}
	}
}
