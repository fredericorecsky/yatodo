# Yet Another Todo 

## Basics:

Yet another todo is a golang app that provides a service for a todo list frontend app

It has configuration now to run bare machine and on docker/docker compose. Can be easily 
deployed on k8s using the images + config.

### Database connection

For all environments the database setup is made setting os environment variables:

#### For sqllite3

    DB_BACKEND=sqlite3
    DB_DSN = <filename>  
    
#### For Mysql

    MYSQL_HOST=hostname or 127.0.0.1 for localhost
    MYSQL_DATABASE=yatodo
    MYSQL_USER=...
    MYSQL_PASSWORD=...
    
    DB_BACKEND=mysql
    DB_DBN=$MYSQL_USER:$MYSQL_PASSWORD@tcp($MYSQL_HOST)$MYSQL_DATABASE?charset=utf8&parseTime=True&loc=Local
    
 For pratical reasons the application uses the full DB_DSN already composed from the mysql vars
 
 ### Building / Running

 We have a Makefile to build , run , test
 
 To just build it locally. 
    
     make build 

 To run it locally also
    
    make run
    
 To build a docker image, ( golang default image )
 
    make build-docker
    
 Also it has a run docker
 
    make run-docker
    
 It has some api tests made with curl, they will do the basic set of operations with the todo list
 
    make test-curl
    
 Not fully implemented, some tests on db
 
    make test-db
    
 
 Running from docker compose:
 
    docker-compose build
    docker-compose up
    
 