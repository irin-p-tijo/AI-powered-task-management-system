# AI-powered-task-management-system

## Overview

This project is a task management system that utilizes AI-powered task suggestions, user authentication. It enables users to create, assign, and track tasks efficiently, with the added feature of AI-powered task breakdowns to enhance productivity. The system is built using Golang 
## Features

- **User Authentication**: JWT-based authentication for secure user sessions.
- **Task Management**: Users can create, assign, and track tasks.
- **AI-Powered Task Suggestions**: Smart task breakdowns and suggestions using OpenAI/Gemini API to improve task completion.
- **Cloud Deployment**: Deployed on Render, Fly.io, or similar cloud platforms for scalability.

## Tech Stack

- **Backend**: Golang (Gin)
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **AI Integration**: OpenAI
- **Containerization**: Docker

## Features Breakdown

### 1. **User Authentication**
   - JWT-based authentication ensures secure login and token management.
### 2. **Task Management**
   - Users can create new tasks, assign them to team members, and track progress.

### 3. **AI-Powered Task Suggestions**
   - The AI feature leverages the OpenAI/Gemini API to suggest task breakdowns.
   - AI analyzes the task's description and breaks it into actionable steps, providing smarter planning.

## API Endpoints

### Authentication

- **POST  /user/usersignup**: user Sign up
- **POST  /user/userlogin**: User login.
### Tasks

- **POST /user/task/createTask**:Create a new task. 
- **GET /user/task/trackTask**: Get all tasks for the authenticated user.
- **POST /user/task/assignTask**: Assignes tasks.
- **POST /user/task/suggestTask**:Suggest tasks


## Setup & Installation

### Prerequisites

- Go 1.18 or later
- PostgreSQL 
- Docker 

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/ai-task-management.git
   cd ai-task-management
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up the database:
   - Create a PostgreSQL 
   - Configure your database connection in the `.env` file.

4. Start the server locally:

   ```bash
   go run main.go
   ```

   By default, the server runs on `http://localhost:8080`.


### Docker 

To run the application with Docker, use the following commands:

1. Build the Docker image:

   ```bash
   docker build -t ai-task-management .
   ```

2. Run the container:

   ```bash
   docker run -p 8080:8080 ai-task-management
   ```


## License

This project is licensed under the MIT License – see the [LICENSE](LICENSE) file for details.