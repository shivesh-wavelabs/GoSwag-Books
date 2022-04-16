# GoSwag-Books
Golang backend with Postgres DB

# Prerequisites
Install Go and add it to the PATH.

# Steps to get started
1. `go mod init GoSwag-Books`
2. `go mod tidy`
3. `go run main.go`

# To Access API Endpoints
1. http://localhost:8000/swagger/index.html

# Swagger UI

![swaggerUI-books](https://user-images.githubusercontent.com/87301008/159532649-3a30551b-3c59-4d3e-9435-35a1181f609e.png)


# Get All Books [GET]

![GetBooks-SwaggerUI](https://user-images.githubusercontent.com/87301008/159532713-c1fba193-3cba-436b-86bc-d1695ae6bc85.png)

# Post a Book [POST]

![PostBook SwaggerUI](https://user-images.githubusercontent.com/87301008/159539070-0b8ba841-1d81-4b76-8db1-772043290639.png)

# Update a Book [PUT]

![UpdateBook SwaggerUI](https://user-images.githubusercontent.com/87301008/159539146-ed6c0942-3f6b-4951-804a-ce396f0b1351.png)

# Delete a Book [DELETE]

![DeleteBook SwaggerUI](https://user-images.githubusercontent.com/87301008/159539625-875df645-7fb5-424c-9d6c-503261f4cb42.png)

# Docker Images
![GoSwag-Books-docker](https://user-images.githubusercontent.com/88183749/163683048-681fb665-9cbd-4872-a1e0-ebcd5d50f66d.png)

 `sudo docker build -t go-swag-books:latest .`
