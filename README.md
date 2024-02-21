# The Todoer 📝

## Overview
Todoer is a simple and straightforward tool designed to help you manage your tasks effectively by creating and generating todo files. With Todoer, you can easily organize your tasks, mark them as completed, and export them to a todo file for future reference.

## Features
- Create a new todo list
- Add tasks to the list
- Remove tasks from the list
- Mark tasks as completed or undone
- Export the todo list to a file

## Installation
Todoer is a standalone executable, so no installation is required. Simply download the executable file for your operating system from the releases page and run it to start using Todoer. An install shell script is also given if you 
want to run the actual source code.

## Usage
To start Todoer, simply execute the binary file. You will be greeted with a command-line interface where you can enter various commands to manage your todo list.

### Available Commands
- `new`: Create a new todo list.
- `add`: Add a new task to the list.
- `remove`: Remove a task from the list.
- `markdone`: Mark a task as done.
- `markfree`: Mark a task as undone.
- `list`: Display the list of tasks.
- `export`: Export the todo list to a file.
- `help`: Display the help message.

## Example Usage
```sh
$ ./todoer

Todoer > new
Enter List Name: SomeTodoList

Todoer > add
Enter task: Complete project documentation

Todoer > add
Enter task: Submit project presentation

Todoer > list

All inserted tasks:
1. [-] Complete project documentation
2. [-] Submit project presentation

Completed Tasks:

Current Pending Tasks:
1. [  ] Complete project documentation
2. [  ] Submit project presentation

Todoer > markd 1

Todoer > list

All inserted tasks:

1. [-] Complete project documentation
2. [-] Submit project presentation

Completed tasks:

1. [X] Complete project documentation

Current Pending Tasks:

2. [  ] Submit project presentation

Todoer > export

Todoer > exit

```

## Version
Todoer Version: 1.0

## Author
- **Aarav Shreshth**
- Email: aaravshreshth7@gmail.com

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Thank you for using Todoer! If you have any feedback or suggestions, feel free to reach out to me. Happy task managing! 📅✨
