FROM alpine

ADD main /main

ENTRYPOINT [ "/main" ]

#docker build -t hollychen503/vipertest .
#ocker run --rm -it  -v $PWD/config.yml:/dhmsconfig/config.yml hollychen503/vipertest:latest
