# IMS Monitoring Platform
This is a web API used to monitor and perform certain actions related to the [IMS business solution](https://github.com/LanternNassi/IMS), such as backing up clients' SQL Server local databases and tracking their payments. The IMS desktop application periodically (`1 day` , `1 week`) sends created backup files of the local databases to this API server, which then calculates the backup parameters, including the individual cost of the backup. Additionally, it holds information about clients and the type of IMS desktop application instances they are running, whether multi-instance or single-instance.

## Features included 
- Billing monthly backup usage
- clients management (Client information )
- Backups management(Calculating sizes for the different backups of the clients' databases)
- Bills management(Issuing invoices to the different clients for their monthly usage)
- Handling payments(Processing payments paid in by the clients to determine validity)

### Clients
  The monitoring platform handles all clients information such as name , address for later usage in other modules.
  
  > Example adding a client for an installation.
  
  ```sh
  curl -X POST http://localhost:8080/clients \
  -H "Content-Type: application/json" \
  -d '{
    "FirstName": "John",
    "LastName": "Doe",
    "Email": "john.doe@example.com",
    "Phone": "123-456-7890",
    "Address": "123 Main St, Springfield, USA",
    "BusinessName": "Doe Enterprises",
    "Status": "Active",
    "ValidTill": "2025-12-31T23:59:59Z"
  }'

  ```

### Bills Management
  This is to issue invoices to the clients about their monthly database backup size usage . This is automatically generated when a backup file is posted under a client's name .It includes the individual cost of every database backup .

### Handling payments
  This is to ensure that the client has fully made payments for their usage. `This is yet to be implemented automatically using the Airtel money API`.

### Client Installations
  This is to add the type of solution installations the client is running . 
  ```sh
  curl -X POST http://localhost:8080/Installations \
  -H "Content-Type: application/json" \
  -d '{
    "ClientID": "123",
    "Installation_type": "Multi-instance",
    "Computer_name": "Sales-PC",
    "IMS_version": "v2.5",
    "Operating_system": "Windows 10",
    "RAM": "16GB",
    "Processor": "Intel Core i7",
    "Active": "true"
  }'

  ```

### Backups Management
  This feature helps handle backups sent in by the clients and compute their total size and bill them accordingly . 
  
  > Example of what the post request from the IMS business solution looks like :
  ```sh
  curl -X POST http://localhost:8080/backups \
  -H "Content-Type: application/json" \
  -d '{
    "ClientID": "123",
    "Name": "Weekly Backup",
    "Backup": "base64_encoded_data_here",
    "Size": 2048,
  }'

  ```



  
## Installation

To set up the IMS Monitoring Platform locally, follow these steps:

1. **Clone the repository**:
    ```sh
    git clone https://github.com/LanternNassi/IMSController.git
    cd IMSController
    ```

2. **Install dependencies**:
    Ensure you have Go installed. Then, run:
    ```sh
    go mod tidy
    ```

3. **Set up environment variables**:
    Configure your environment variables as needed or create a .env file, for example:
    ```sh
    export DB_HOST=your_database_host
    export DB_USER=your_database_user
    export DB_PASSWORD=your_database_password
    ```

4. **Run the application**:
    ```sh
    go run .
    ```

5. **Access the API**:
    Open your web browser and navigate to `http://localhost:8080`.

   
## Using Make and Docker-Compose

For users with `make` and `docker-compose`, you can set up the project using the following commands:

1. **Build and run the services**:
   > Start all the required containers for the API . By default the containers are built and started automatically.
    ```sh
    make start
    ```

3. **Running tests**:
   > Run tests for the API using netlify
    ```sh
    make test
    ```


> The frontend for this API is still under development 
