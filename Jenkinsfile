pipeline {
    agent { docker { image 'golang:1.11.1' } }

    stages {
        stage('Build') {
            steps {
                echo 'Importing packages ...'
                sh 'go get -d -v ./...'
                echo 'Building binary ...'
                sh 'go build -o main'
                echo 'Done: building binary'
                // echo 'building new docker image'
                sh 'sudo docker build -t ganboonhong/reader .'
                // echo 'Done: building new docker image'
            }
        }
        stage('Test') {
            steps {
                echo 'Running tests'
                sh 'go test ./...'
                echo 'Done: testing'
            }
        }
        /*stage('Deploy') {
            steps {
                echo 'Deploying....'
                sh 'docker-compose down'
                sh 'docker-compose up -d'
            }
        }*/
    }
}