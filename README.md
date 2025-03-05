### README

# GoClassificator

GoClassificator is a web application that allows you to manage MySQL database connections, execute scans on the databases, and view scan reports. The application is built using Go and the Echo framework.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

### Step 1: Clone the Repository

```sh
git clone https://github.com/sacortesro/GoClassificator.git
cd GoClassificator
```

### Step 2: Build and Run the Application with Docker compose

1. **Run the Docker compose**:

```sh
docker compose -f docker-compose-mysql.yml up --build -d
```

  * **go-classificator stopped**:

  Application container may stop. It occurs because can't stablish the database connection. Start manually the container.

  ```sh
    docker start go-classificator
  ```

2. **Build and Run the Docker of the test DB**:


```sh
cd employeesdb
docker build -t mysqldb-employees .
docker run --name employeedb-mysql -v mysql-employees:/var/lib/mysql -d -p 3307:3306 mysqldb-employees:latest
```

### Step 3: Access the Application

Open your web browser and navigate to `http://localhost:8080`.

## API Endpoints

### Create API Key

- **Endpoint**: `POST /apikey`
- **Description**: Creates a new API key.
- **Response**:
  ```json
  {
    "apiKey": "your-api-key"
  }
  ```

### Register MySQL Connection

- **Endpoint**: `POST /api/v1/database`
- **Description**: Registers a new MySQL database connection.
- **Request Body**:
  ```json
  {
    "host": "localhost",
    "port": 3306,
    "username": "root",
    "password": "password",
    "dbname": "mydatabase"
  }
  ```
- **Response**:
  ```json
  {
    "id": 1
  }
  ```

### Execute Scan

- **Endpoint**: `POST /api/v1/database/scan/:id`
- **Description**: Executes a scan on the specified database.
- **Path Parameter**:
  - `id`: The ID of the database connection.
- **Response**:
  ```json
  {
    "message": "Scan successful"
  }
  ```

### Get Scan Result

- **Endpoint**: `GET /api/v1/database/scan/:id`
- **Description**: Retrieves the result of a scan.
- **Path Parameter**:
  - `id`: The ID of the scan.
- **Response**:
  ```json
  {
    "date": "January 2, 2006 15:04:05",
    "databaseName": "mydatabase",
    "host": "localhost",
    "scanCount": 1,
    "totalTables": 10,
    "totalColumns": 50,
    "dataTypesSummary": [
      {
        "type": "String",
        "count": 20
      },
      {
        "type": "Integer",
        "count": 15
      },
      {
        "type": "Date",
        "count": 10
      },
      {
        "type": "Boolean",
        "count": 5
      }
    ],
    "tables": [
      {
        "name": "users",
        "columnCount": 5
      },
      {
        "name": "orders",
        "columnCount": 3
      }
    ]
  }
  ```

### Render HTML Report

- **Endpoint**: `GET /view/report/:id`
- **Description**: Renders the scan report as an HTML page.
- **Path Parameter**:
  - `id`: The ID of the scan.
- **Response**: HTML content of the scan report.

## Directory Structure

```
project-root/
├─── main.go           		
├── internal/                 
│   ├── api/                  
│   │   ├── auth/		        # Contiente la lógica para autenticación
│   │   ├── controllers/    # Controlador de las rutas   
│   │   ├── middleware/     # Middleware para autenticación  
│   │   └── routes.go      	
│   ├── config/               
│   ├── database/             
│   │   ├── models/         # Define los modelos de los datos
│   │   └── repository/     # Contiene la lógica de la base de datos
│   ├── logger/			
│   ├── security/           # Seguridad en los servicios (encripción)   
│   ├── services/       	  # Lógica de negocio de las rutas
├── test/                     
├── web/                          
│   └── templates/		      # Web (HTML)
├── .env              
├── Dockerfile                
├── docker-compose.yml        
├── go.mod                    
├── go.sum                    
└── README.md     

```