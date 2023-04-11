# docker-training

Purpose of this repository to learn docker and writing dockerfiles.
I have added descriptions with the commands in dockerfile. 

### Commands used to run the container

1. `docker run -p 8081:8080 gin-server` is the command used to run this application, `-p` is flag that is used to expose port for host. e.g `-p <hostPort>:<containerPort>`.
2. `docker run -p 8081:8080 -e VAR1=ValuePassedDuringRunningContainer gin-server` in this command, new this is `-e` that is used to pass value for the environment variable that is defined in the dockerfile.
3. `docker run -p 8081:8080 -e VAR1=ValuePassedDuringRunningContainer -d gin-server` in this command new thing is `-d` , that is used to run container in detached mode.

### Push image to docker hub.
1. First I created account to `hub.docker.com`.
2. I logged in into my account using `docker login` command.
3. I tagged my image using `docker tag <image> <dockerHubUser>/<image>` (if I don't do this then I will face an error smth like this `requested access to the resource is denied`.
4. Then I pushed my image using `docker push <dockerHubUser>/<image>`.

1st commit before rebase
2nd commit before rebase

commmit from branch that I want to rebase.