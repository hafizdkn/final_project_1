# final_project_1

***Create task***
----
* **URL**

  ```
  localhost:8080/todos
  ```

* **Method:**

  ```
  POST
  ```

* **Data Params**

   **Required:**  
  ```json
   {
     "task":"squat jump 1 sets"
   }
  ```

* **Success Response:**
 
  ```json
  {
    "status": 201,
    "message": "Succes create todo task",
    "payload": {
        "Id": 1,
        "Task": "squat jump 1 sets",
        "Completed": false,
        "CreatedAt": "2022-10-19T17:13:37.941965084+07:00",
    }
  }
  ```
 
* **Error Response:**
  ```json
  {
    "status": 400,
    "message": "Bad request",
    "error": "Key: 'TodoInput.Task' Error:Field validation for 'Task' failed on the 'required' tag"
  }
  ```
  
***Get all task***
----
* **URL**

  ```
  localhost:8080/todos
  ```

* **Method:**

  ```
  GET
  ```
  
* **Success Response:**
 
  ```json
  {
    "status": 200,
    "message": "Success get all todo task",
    "payload": [
        {
            "Id": 1,
            "Task": "squat jump 1 sets",
            "Completed": false,
            "CreatedAt": "2022-10-19T17:13:37.941965084+07:00",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": "0001-01-01T00:00:00Z"
        },
        {
            "Id": 2,
            "Task": "push up 1 sets",
            "Completed": true,
            "CreatedAt": "2022-10-19T17:13:37.941965084+07:00",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": "0001-01-01T00:00:00Z"
       }
    ]
  }
  ```
  OR
  ```json
  []
  ```
  
***Get todo by id***
----
* **URL**

  ```
  localhost:8080/todos
  ```

* **Method:**

  ```
  GET
  ```
*  **URL Params**

   **Required:**
 
   `id=[integer]`
   
  
* **Success Response:**
 
  ```json
  {
    "status": 200,
    "message": "Success get todo",
    {
        "Id": 1,
        "Task": "squat jump 1 sets",
        "Completed": false,
        "CreatedAt": "2022-10-19T17:13:37.941965084+07:00",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": "0001-01-01T00:00:00Z"
    }
  }

* **Error Response:**
  ```json
  {
    "status": 500,
    "message": "Internal server error",
    "error": "record not found"
  }
  ```
  
* **Sample Call:**
  ```
  localhost:8080/todos/1
  ```
  
***Update todo***
----
* **URL**

  ```
  localhost:8080/todos
  ```

* **Method:**

  ```
  PUT
  ```
*  **URL Params**

   **Required:**
 
   `id=[integer]`
   
* **Data Params**

   **Required:**  
  ```json
  {
      "task":"squat jump 1 sets",
      "completed":true
  }
  ```
  
* **Success Response:**
 
  ```json
  {
    "status": 200,
    "message": "Success get todo",
    {
        "Id": 1,
        "Task": "squat jump 1 sets",
        "Completed": true,
        "UpdatedAt": "2022-10-19T18:06:33.423415679+07:00",
    }
  }
  
* **Error Response:**
  ```json
  {
    "status": 500,
    "message": "Internal server error",
    "error": "record not found"
  }
  ```
  
* **Sample Call:**
  ```
  localhost:8080/todos/1
  ```

***Delete todo***
----
* **URL**

  ```
  localhost:8080/todos
  ```

* **Method:**

  ```
  DELETE
  ```
*  **URL Params**

   **Required:**
 
   `id=[integer]`
   
* **Success Response:**
 
  ```json
  {
    "status": 200,
    "message": "Success delete todo"
  }
  ```

* **Error Response:**
  ```json
  {
    "status": 500,
    "message": "Internal server error",
    "error": "record not found"
  }
  ```

  
* **Sample Call:**
  ```
  localhost:8080/todos/1
  ```
