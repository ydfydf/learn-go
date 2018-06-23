pipeline {
    agent Test

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
            }
        }
        stage('Test') {
            steps {
                echo 'Run nginx docker'
                sh /root/go/run-nginx-docker.sh 
            }
        }
    }
}
