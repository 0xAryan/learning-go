- Cannot level an empty file go complains so i change
  the old hello.go to hello.go.txt

- Need to create go.mod file with the name of the main file before running test
    - `go mod init <file_name>`

###Discipline
- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor