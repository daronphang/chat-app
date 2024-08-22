# Deployment

## AWS

Install dependencies on AWS-Linux

```sh
$ sudo yum update -y
$ sudo yum install -y docker
$ sudo service docker start
$ sudo usermod -a -G docker ec2-user

$ sudo curl -L https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
$ sudo chmod +x /usr/local/bin/docker-compose
$ docker-compose version

$ sudo yum install git -y
```

## Overlay Network

As all services are deployed in Docker, they must be connected to an overlay network in order for the containers to communicate with each other. DNS routing is automatically performed.

1. On primary node, declare it as swarm manager

```sh
$ docker swarm init
$ docker swarm join-token worker # for worker nodes
```

2. Create overlay network

```sh
$ docker network create -d overlay --attachable chatapp
```

3. On worker nodes, connect to swarm manager

```sh
$ docker swarm join --token <manager token>
```

4. Create a dummy container to connect to the overlay network to expose it on worker node (workaround for bug that does not connect the containers to the overlay network when using docker compose)

```sh
$ docker run -dit --name busybox --network chatapp busybox
```

2. Spin up Docker compose files in the following order: message-service, user-service, session-service, chat-service, chat-ui
