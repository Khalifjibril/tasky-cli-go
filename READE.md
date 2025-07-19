# Tasky CLI To-Do Manager

![Go Lang](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)

## 📝 Description

**Tasky** is a lightweight and intuitive command-line interface (CLI) application built with **Go** that helps you efficiently manage your daily tasks. It provides core functionalities to add, view, update, and delete tasks, ensuring your tasks are persistently stored and readily available whenever you need them. Designed for simplicity and speed, Tasky is perfect for those who prefer working directly from the terminal.

## ✨ Features

* **Add Tasks:** Quickly add new tasks with automatically assigned unique IDs.
* **View Tasks:** Display a clear, numbered list of all your current tasks.
* **Update Tasks:** Modify the content of any existing task by simply providing its ID.
* **Delete Tasks:** Remove completed or unwanted tasks by their unique ID, with confirmation.
* **Persistent Storage:** All tasks are saved to a local `data.json` file, ensuring your data is retained across sessions.
* **Robust Input Handling:** Includes validation for numerical inputs and prevents creating empty tasks.

## 🚀 Getting Started

Follow these steps to get your Tasky CLI To-Do Manager up and running on your local machine.

### Prerequisites

* **Go:** Make sure you have Go installed (version 1.22 or higher recommended).
    * [Download Go](https://go.dev/dl/)

### Installation & Running

1.  **Clone the Repository:**
    Open your terminal or Command Prompt and clone the project to your local machine:
    ```bash
    git clone [https://github.com/Khalifjibril/tasky-cli-go.git](https://github.com/Khalifjibril/tasky-cli-go.git)
    ```
    *(Remember to replace `tasky-cli-go` if you chose a different repository name on GitHub)*

2.  **Navigate to the Project Directory:**
    ```bash
    cd tasky-cli-go
    ```

3.  **Run the Application:**
    You can run Tasky directly:
    ```bash
    go run main.go
    ```
    Or, you can build an executable for easier future use:
    ```bash
    # For Windows:
    go build -o tasky.exe

    # For macOS/Linux:
    go build -o tasky
    ```
    Then, run the executable:
    ```bash
    # For Windows:
    ./tasky.exe

    # For macOS/Linux:
    ./tasky
    ```

## 🎮 Usage

Once the application is running, follow the simple on-screen menu prompts:

1.  **Add a task:** Enter `1` and press Enter.
2.  **Check all tasks:** Enter `2` and press Enter.
3.  **Update a task:** Enter `3` and press Enter, then provide the task ID and new content.
4.  **Delete a task:** Enter `4` and press Enter, then provide the task ID.
5.  **Exit the application:** Enter `5` and press Enter.

## 🛠️ Technologies Used

* **Go Programming Language:** The core language used for building the application.
* **Standard Go Libraries:** `fmt`, `os`, `bufio`, `strings`, `encoding/json`, `slices`.

## 💡 Future Enhancements (Roadmap)

* **Task Completion Status:** Ability to mark tasks as "done."
* **Due Dates:** Add functionality to set and display due dates for tasks.
* **Task Prioritization:** Implement a system for high, medium, and low priority tasks.
* **Search/Filter:** Allow users to search for tasks by keywords or filter by status/priority.
* **Confirmation Prompts:** Add "Are you sure?" prompts for delete operations.
* **Improved UI/UX:** Explore more visually appealing terminal output.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
*(You might need to create a separate `LICENSE` file in your project root with the MIT license text if you want to include it. For a simple project, linking to it in the README is fine.)*

---