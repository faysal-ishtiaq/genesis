# Genesis

An opinionated API scaffolding tool using gorm and gin-gonic

## Get started

go get https://github.com/faysal-ishtiaq/genesis

- mkdir appName
- cd appName
- genesis init -n appName -e sqlite
- go generate # to install dependencies and intializing go module
- genesis add -s serviceName

Then add serviceName.InitDB() in initDB() function of main.go, add serviceName.BindRoutes() in main function in main.go, change model.go, dto.go, api.go the way you want it!

### Example
Let's build a test app with Todo service

- mkdir test && cd test
- cd test
- genesis init -n test -e sqlite
- genesis add -s todo -m task
