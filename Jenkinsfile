pipeline {
    agent { docker { image 'ganboonhong/reader' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}