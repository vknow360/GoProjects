# 📝 To-Do List CLI App

A simple and efficient **To-Do List CLI App** built using **Go** that allows users to manage their tasks effortlessly from the command line.

## 🚀 Features

✅ **Add Tasks** with priority, category, and optional deadline. \
✅ **List Tasks** with well-formatted output.\
✅ **Mark Tasks as Done** or **Undo** completion.\
✅ **Remove a Task** or **Clear All Tasks**.\
✅ **Persistent Storage** (tasks are saved in JSON format).\
✅ **Beautiful UI** with colors and proper alignment.\
✅ **Minimal Dependencies** (`fatih/color`, `mattn/go-runewidth`).

---

## 📦 Installation

### **Prerequisites**

-   Install [Go](https://go.dev/dl/) (>=1.18)

### **Clone the Repository**

```sh
git clone https://github.com/vknow360/GoProjects.git
cd todo-app
```

### **Build the Binary**

```sh
go build todo.go
```

## 📌 Usage

### **Add a Task**

```sh
$ todo add "Complete Golang project" high development 2025-03-15
✅ Task added: [1] Complete Golang project | Priority: high | Category: development | Deadline: 2025-03-15
```

### **List All Tasks**

```sh
$ todo list
📌 Your Tasks:
--------------------------------------------------------------------------------
ID   Status          Priority   Category      Deadline     Title
--------------------------------------------------------------------------------
1    ❌ Incomplete   high       development   2025-03-15   Complete Golang project
--------------------------------------------------------------------------------
```

### **Mark Task as Done**

```sh
$ todo done 1
✅ Task [1] marked as done!
```

### **Remove a Task**

```sh
$ todo remove 1
✅ Task [1] removed!
```

### **Clear All Tasks**

```sh
$ todo clear
✅ All tasks removed successfully!
```

### **Help Menu**

```sh
$ todo help
```
