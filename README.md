**LIBRARY CLI**
----
   Library cli is used to interact with the library REST Api. 

* **CLI features**

 - login
 - logout
 - get all users
 - get specific user
 - get all books
 - get specific book
 - user takes book
 - user returns book
 - admin adds book
 - admin deletes book

* **Requirments**
  - go 1.12+

* **Building**
  - `go build -o library`

* **Common commands and flags:**

All of the flag used below are required for the specific commands and the cli can not be used without them.

  - `library login -e=<your email> -p=<password>` - logs the user
  - `library logout -t=<your jwt token>` - logouts the user
  - `library get-all -t=<your jwt token>` - show all books
  - `library get -i=<isbn> -t=<your jwt token` - shows book with the provided isbn
  - `library delete -i=<isbn> -t=<your jwt token` - deletes book with the provided isbn
  - `library save -i=<isbn> -title=<title> -a=<author> -u=<available units> -t=<your jwt token` - saves a book with the provided properties.
  - `library take -i=<isbn> -e=<user email> -t=<your jwt token` - user with this email takes the book
  - `library return -i=<isbn> -e=<user email> -t=<your jwt token` - user with this email returns the book
  - `library get-all-users -t=<your jwt token>` - shows all users
  - `library get-all -e=<user email> -t=<your jwt token>` - shows user with this email

  
*  **Finding commands**

    Use the `library --help` or `library -h` argument to get a complete list of available commands.
    
    Also use `library <command> --help` or `library <command> -h` to get cli reference information about the command.


*  **Using the cli**
   
   After logging the user recieves his token for using the cli. The user should provide his token when making a request with the cli, othwerwise he would not be able to access the Library Resp API.


* **Examples**

   * example of the get command for book: `library get -i="123456" -t="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiI2NTQwMmRlOS03MDI4LTQ5MzEtOTY4MS1iODkxNTUyN2NlYTMiLCJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MTA5ODA4NDcsInVzZXJfaWQiOjEsInVzZXJfcm9sZSI6IkFkbWluIn0.1uhHWJ81iG1UgCbqfZsPfwf-D5v5hh3oGpiEbshra18"`


   * example of the save command for book: `library save -i="123456" -n="Book Title" -a="Ivan Vazov" -u=12 -t="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiI2NTQwMmRlOS03MDI4LTQ5MzEtOTY4MS1iODkxNTUyN2NlYTMiLCJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MTA5ODA4NDcsInVzZXJfaWQiOjEsInVzZXJfcm9sZSI6IkFkbWluIn0.1uhHWJ81iG1UgCbqfZsPfwf-D5v5hh3oGpiEbshra18"`

  * example of the login command: `library login -e="misho@gmail.com" -p="12213"`