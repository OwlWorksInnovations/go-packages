# How do I add more packages?

## Set up the package
You need to make a new folder then init the folder with its remote (github) url (including folder)
```
go mod init github.com/OwlWorksInnovations/go-packages/database
```

## The .go file (refer to test.go)
For the package to be usable by your app, it must use the correct package name and export its members (structs and functions) by capitalizing their first letter.
According to the Go Language Specification, any identifier that does not start with a capital letter is unexported and invisible to packages outside of its own folder.

## Set up github
You need to push everything to master/main branch then create a tag with the folder prefix
```
git tag database/v1.0.0
git push origin database/v1.0.0
```

# How to use it?
Once tagged, anyone can pull it
```
go get github.com/OwlWorksInnovations/go-packages/logger@v1.0.0
```
