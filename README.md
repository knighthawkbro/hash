What it can do:
- Create an hashset and perform functions as a datatype hash
- Runs through a driver program testing all the hash functions
- Runs through a seuss program to print all distinct words as a Hashset
- Loads a dictionary file and a test document and spell checks the words against a 140,000 word hashset

How to build yourself:
- DL [Golang](https://golang.org/dl/)
- CD to hash directory
- run ```go build hash.go```
- remove other executable ```rm hash```
- run ```.\hash.go```
- conversly you can just run ```go run hash.go``` for an instant execution of the program
- I think you may need to set your [$GOPATH](https://golang.org/doc/code.html#GOPATH) and create a src folder within for go build to work