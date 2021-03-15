# TECHNICAL TEST SOFTWARE ENGINEER (GOLANG)
  
To start the API run the following command: go run *.go  
The API will start running in localhost:8000  
  
# ENDPOINT 1
This is the first endpoint: http://localhost:8000/api/sort  
This endpoint will receive an array of integers, sort it and move the repeated numbers to the end of the array. The data must be sent like this:  
  
![Screenshot](./images/sortBody.png)  
  
The response will return the sorted array:  
  
![Screenshot](./images/sortResponse.png)  
  
# ENDPOINT 2
This is the second endpoint: http://localhost:8000/api/user/id  
This enpoint will return the detailed info about the user associated with the id number:  
  
![Screenshot](./images/userResponse.png)  
  
If the user doesn't exist the response will be like this:  
  
![Screenshot](./images/userNotFound.png)  
  
All the info is stored in a dummy database found in the file database.go
