# golang-viper-config-example

Simple example. Currently no tests implemented, no validity check of given parameters. Feel free to add tickets for more details, documentation, other examples

To run this example in linux use following:

``` 
go get github.com/devilsray/golang-viper-config-example
cd $GOPATH/src/github.com/devilsray/golang-viper-config-example/
go run *.go
```




### viper tester

```
docker run --rm -it  -v $PWD/testfolder:/dhmsconfig:ro  hollychen503/vipertest:latest
```

### alpine

```
apk update && apk upgrade && apk add --no-cache bash git openssh
```

```
docker run --rm -it -v $PWD/testfolder:/dhmsconfig hollychen503/alpine-git   bin/sh
```
