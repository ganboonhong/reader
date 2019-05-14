pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                go build -o main
            }
        }
        stage('Test') {
            steps {
                go test ./...
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}