# DSC-Golang-Docker-Kubernetes

Made with :heart: By DSC ICNJ and DSC UMass




## Docker

Building the docker image

```bash 
sudo docker build -t golang-server .
```

But what is this you ask? Lets break it down.

First thing is the `build`, this tells docker to look for a docker file and run the commands in it.

`-t` is short for tag, which is for what you want to tag the docker container as to refer to later. 

`golang-server` is the tag name given.

And lastly the `.` which tells the build command where it can look for the Dockerfile.


One of the most amazing appeals of docker is how flexible and vast it actually is. No where better can you see it than on [Dockerhub](https://hub.docker.com/). Think of dockerhub as a npm for docker images. Theres anything you need from a nginx, mongo or minecraft server images that can all be ran in one command and run flawlessly on any x86 computer with docker.

