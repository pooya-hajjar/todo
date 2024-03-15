# mvc todo app

## prerequisites
* todo_app database in postgres

## usage

First of all, run the migrate file in the migration folder

```bash
migrate -path ./migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up
```

Create the .env file
```bash
touch .env
```

Put your variables in that file

```env
APP_PORT=3000
POSTGRES_DSN=postgres://postgres:<YOUR_PASSWORD>@localhost:5432/todo_app
JWT_SECRET_KEY=WhatsUpMan

GOOGLE_CLIENT_ID=<YOUR_CLIENT_ID>
GOOGLE_CLIENT_SECRET=<YOUR_CLIENT_SECRET>
DEFAULT_USER_PASSWORD=YoBuddy
```
The last 3 lines are use for login with Google (OAuth2) and are optional

Install dependencies and Run project 
```bash
go mod tidy
go run ./main.go
```


## Document
````
GET    /   
POST   /auth/signup                          username:string password:string avatar?:string
POST   /auth/signin                          username:string password:string
POST   /auth/google-login      
POST   /auth/google-callback 
POST   /auth/logout  
GET    /user/info                
GET    /user/search/:id      
PUT    /user/update                          username:string password:string avatar?:string status?:int  
GET    /top_ten      
GET    /tasks/get               
GET    /tasks/get/:id                        
POST   /tasks/add                            title:string priority?:int start_date?:datetime end_date?:datetime 
DELETE /tasks/delete/:id        
PATCH  /tasks/rename/:id                     ?title=string        
PATCH  /tasks/update/:id                     ?status=int   
````

## todo
- [ ] add cache system