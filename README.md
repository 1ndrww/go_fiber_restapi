#  Simple Go Rest API for [Fiber](https://github.com/gofiber/fiber)


## Usage

1. Download [Go Fiber](https://github.com/gofiber/fiber) by using: 

   ```bash
   $ go get -u github.com/gofiber/fiber
   ```

2. Download [Go-UUID](https://github.com/google/uuid) by using:

   ```bash
   $ go get -u github.com/google/uuid
   ```

3. Download [Go Validator](https://github.com/google/uuid) by using:

   ```bash
   $ https://github.com/go-playground/validator
   ```


4. Or Run the `go get .` command to download all package on this project. 

   ```bash
   $ go get .
   ```

5. Run this command to start the server:

   ```bash
   go run main.go
   ```

6. Horray! You're ready to Go 🚀

7. This is an example request for create user : 

   ```go
   {
    "user_name" : "dsfsdf",
    "email" :"3rr",
    "phone" :"ewtwe",
    "full_name" :"dsfdfdsf"
   }
   ```
   example of response create user :

   ```go
   {
   "code": 200,
   "data": {
   "ID": "b4affbaf-b5a0-47d0-8dd2-cdfbda4ba218",
   "UserName": "dsfsdf",
   "Email": "3rr",
   "Phone": "ewtwe",
   "FullName": "dsfdfdsf",
   "CreatedAt": "1697297887",
   "UpdatedAt": ""
   },
   "error": null
   }
   ```
8. example of response get all users :

   ```go
   {
       "code": 200,
       "data": [
           {
               "ID": "b4affbaf-b5a0-47d0-8dd2-cdfbda4ba218",
               "UserName": "dsfsdf",
               "Email": "3rr",
               "Phone": "ewtwe",
               "FullName": "dsfdfdsf",
               "CreatedAt": "1697297887",
               "UpdatedAt": ""
           }
       ],
       "error": null
   }
   ```
9. This is an example response get by id users failed: 
   ```go
   {
       "code": 404,
       "data": null,
       "error": "user id is not found"
   }
   ```




