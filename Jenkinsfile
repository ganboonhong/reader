pipeline {
    agent { docker { image 'golang:1.11.1' } }

    stages {
        stage('Build') {
            steps {
                echo 'Building....'
                sh 'go get -d -v ./...'
                sh 'go build -o main'
                echo 'Done building'
            }
        }
        stage('Test') {
            steps {
                echo 'Running tests'
                sh 'go test ./...'
                echo 'Done testing'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}