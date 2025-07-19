package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
	"log"
	"io"
)

type Task struct {
	Id int
	Content string
}

const Create int = 1
const Check int = 2
const Update int = 3
const Delete int = 4

var counter int

func init() {
	tasks := readData()
	if len(tasks) == 0 {
		counter = 1
		return
	}
	maxId := 0
	for _, task := range tasks {
		if task.Id > maxId {
			maxId = task.Id
		}
	}
	counter = maxId + 1
}

func main() {
	fmt.Println("Welcome to Tasky - your simple CLI to-do manager!")

	for {
		choice := getChoice()

		switch choice {
			case Create:
				create()
			case Check:
				fmt.Println("These are your existing tasks")
				 check()
			case Update:
				update()
			case Delete:
				 delete()
			default:
				fmt.Println("Invalid choice!")
		}
	}	
}

	
func getChoice() int {
	fmt.Println("Choose an option")
	fmt.Println("1. Create your to-do list:")
	fmt.Println("2. Check your to-do list")
	fmt.Println("3. update your to-do list")
	fmt.Println("4. Delete") 
	var choice  int 
	fmt.Scanln(&choice)
	return choice 
}


func create() {
	task := getTask()
	var tasks = readData()
	tasks = append(tasks, task)
	saveData(tasks)
	fmt.Printf("Your new task is %v and its id is %v\n", task.Content, task.Id) 
}

func getTask() Task {
	fmt.Println("What is your new task?")
	var id = counter 
	var content  string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		content = scanner.Text()
		content = strings.TrimSpace(content)
	} else if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err) 
	}
	counter++
	return Task{Id: id, Content: content}
}



func check() {
	tasks := readData()
	for _, task := range tasks {
		fmt.Printf("id = %v, task = %v \n", task.Id, task.Content)
	}
}


func delete() {
	fmt.Println("These are your existing tasks")
	check()
	fmt.Println("Please put in the Id of the task you want to delete.")
	var idToDelete int
	_, err := fmt.Scanln(&idToDelete)
		if err != nil {
			fmt.Println("Invalid inpt. Please enter a vaid number for the task ID")
			bufio.NewReader(os.Stdin).ReadString('\n')
			return 
		}
	var tasks = readData()
	found := false

	
	for index, task := range tasks {
		if task.Id == idToDelete {
			tasks = slices.Delete(tasks, index, index+1)
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Task with ID %d not found. No task was deleted.\n", idToDelete)
		return
	}

	saveData(tasks)
	fmt.Printf("Task with %d deleted succesifuly \n", idToDelete)
	fmt.Println("your updated tasks are:")
	check()
}

func update() {
	fmt.Println("These are your existing tasks")
	check()
	fmt.Println("Please put in the Id of the task you want to update.")
	var id int
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Invalid input! Please put in a valid number for the task ID")
		bufio.NewReader(os.Stdin).ReadString('\n')
		return
	}
	bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Update your chosen task:")
	var content string 
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		content = scanner.Text()
		content = strings.TrimSpace(content)
	} else if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err) 
		return
	}
	if content == "" {
		fmt.Println("Task content cannot be empty. Update cancelled")
		return
	}

	var tasks = readData()
	found := false
	for index, task := range tasks {
		if task.Id == id {
			tasks[index].Content = content
		found = true
		break
	}
}
	if !found {
		fmt.Printf("Task with ID %d not found. No task was updated.\n", id)
		return
	}
	saveData(tasks)
	fmt.Printf("Task with ID %d successifuly updated.\n ", id)
	fmt.Println("Your udated tasks are:")
	check()
}

func readData() []Task {
	file, err := openFile("data.json", os.O_RDONLY|os.O_CREATE)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("could not read the users.json file: %v", err)
	}
	var tasks = []Task {}
	json.Unmarshal(bytes, &tasks)
	return tasks
}


func openFile(name string, flags int) (*os.File, error) {
	file, err := os.OpenFile(name, flags, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func saveData(tasks []Task) {
	bytes, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		log.Fatalf("could not convert data into json")
	}
	file, err := openFile("data.json",  os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()
	_, err = file.Write(bytes)
	if err != nil {
		log.Fatalf("could not update the database: %v", err)
	}
}

