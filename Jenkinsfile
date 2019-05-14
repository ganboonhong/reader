pipeline {
    agent { docker { image 'golang:1.11.1' } }

    stages {
        stage('Build') {
            steps {
                sh 'go build -o main'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}