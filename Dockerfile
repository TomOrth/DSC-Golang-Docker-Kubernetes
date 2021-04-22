
# The "FROM" command selects the base image to build and execute subsequent instructions on,
# such as the ones below that we will go through. For this image we are using the official golang image. https://hub.docker.com/_/golang/
# This is pretty much a basic debian OS pre-installed with the latest Golang.
FROM golang

# The "RUN" command is to instruct any commands you want to run while in the building process,
# but not while the container is actually running.
RUN mkdir /server

# The "ADD" command lets us put local files on our computer onto the built image. 
# All you need to specify is what file or directory you want to add then were to put it.
# Here, we are taking simple_web.go and putting it into the server folder on our image.
ADD simple_web.go /server

# "WORKDIR" send the current folder that all commands with be executed in. Usually, when running
# a command, it starts a whole new bash session at the root but can now be specified with WORKDIR
WORKDIR /server

# Here, we use the run command to compile the added golang code to machine code in order to
# run it smoothly and easily.
RUN go build simple_web.go

# The "CMD" command is one of  the most usefull. Like its brother RUN, CMD takes commands
# to run at startup when the docker container is ran. For example, here we take our compiled code and run it.
CMD ["./simple_web"]

# "EXPOSE" informs Docker that the container listens on the specified network ports at runtime. 
# This then enables us to externaly talk to the container through port 8080
EXPOSE 8080 