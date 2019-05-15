/*
Command to start Jenkins container

docker run \
-u root \
--name jenkins \
--detach \
--publish 8080:8080 \
-v $HOME/jenkins:/var/jenkins_home \
-v /var/run/docker.sock:/var/run/docker.sock \
-v $HOME/.docker:/root/.docker \
-v $HOME/ jenkinsci/blueocean:latest


*/
pipeline {
    /*
    An agent is typically a machine, or container, which connects to a Jenkins master and executes tasks when directed by the master.
    In this case, the agent is a docker container. We run the golang:1.11.1 image to get the container with the `go` that comes with `go` command.
    You can build  your own customized image and push it to docker hub in order to use it here.
    */
    agent { 
        docker { 
            image 'golang:1.11.1'
            // In this case, Jenkins container will be the host since the Jenkins container process triggered this pipeline
            // in order to use docker command, we need to map the docker command in 'jenkins' container to 'current testing' container (golang:1.11.1)
            args '-v /var/run/docker.sock:/var/run/docker.sock -v /usr/bin/docker:/usr/bin/docker'
        }
    }

    stages {
        stage('Build') {
            steps {
                echo 'Importing packages ...'
                sh 'go get -d -v ./...'

                echo 'Building binary ...'
                sh 'go build -o main'
                echo 'Done: building binary'

                echo 'building new docker image'
                sh 'docker build -t ganboonhong/reader .'
                echo 'Done: building new docker image'

                echo 'pushing new image'
                sh 'docker login'
                sh 'docker push ganboonhong/reader'
                echo 'Done: pushing new image'
            }
        }

        stage('Test') {
            steps {
                echo 'Running tests'
                sh 'go test ./...'
                echo 'Done: testing'
            }
        }

        stage('Deploy') {
            steps {
                echo 'Please run ./docker-compose-reload.sh manually to reload'
            }
        }
    }
}