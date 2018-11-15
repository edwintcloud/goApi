# goApi
My first API in golang. This project is a work in progress and learning process for me.

## Environment
This project uses a .env file which is not uploaded to github, please see .env-example file for reference.

## Databases
This project uses MySql, please make sure MySql is setup properly and a table is created. With proper .env file setup, this project will work as is with mysql. 

This project also uses MongoDB, make sure the latest version is installed and `mongod` is running.

## Setup
Make sure you have govendor installed and `$GOPATH/bin` in your path. Go should be installed and your go path should be setup correctly. Make sure to clone the repo in your go path.
```sh
git clone https://github.com/edwintcloud/goApi.git goApi
govendor sync
go run server.go
```