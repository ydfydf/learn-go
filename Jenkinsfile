pipeline {
    agent {
        label 'YdfTest'
    }

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
            }
        }
        stage('Test') {
            steps {
                echo 'Run nginx docker'
                sh './echo-test.sh'
            }
        }
    }
}
