# boilerplate

boilerplate for beginner. Using go fiber as base framework. This bolerplate compatible with fly.io CI/CD deployment.


## Folder Structure

This boilerplate have several folder with different function, such as:
* url : same as routes, this folder act to route URL in browser into controller
* config : all of apps configuration like: database, api, token.
* contoller : all of endpoint function
* model : all of data structure using in this apps


## Rename Apps

If you want to rename apps, please delete go.mod and go.sum file first, then type command in your terminal or cmd :

```sh
go mod init gocroot
go mod tidy
```

Then replace package import with the name gocroot/... in the main.go, controller.go dan url.go

