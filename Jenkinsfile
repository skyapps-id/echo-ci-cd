pipeline {
    agent any

    environment {
    registryName = "echo-ci-cd"
    registryCredential = 'ACR'
    dockerImage = ''
    registryUrl = 'efishery.azurecr.io'
    
    stages {
        stage ('Checkout') {
            steps {
            checkout([$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[url: 'https://github.com/skyapps-id/echo-ci-cd.git']]])
            }
        }
        
        stage('Build Docker Image') {
            steps {
                script {
                    def imageName = "${registryName}:${BUILD_NUMBER}"
                    dockerImage = docker.build(imageName, "-f path/to/Dockerfile .")
                }
            }
        }
        
        stage('Unit Test') {
            steps {
                script {
                    dockerImage.inside {
                        sh 'go test ./...'
                    }
                }
            }
        }
        
        stage('Upload Image to ACR') {
            steps {
                script {
                    def imageName = "${registryName}:${BUILD_NUMBER}"
                    dockerImage.tag(imageName)
                    dockerImage.push('latest')
                }
            }
        }
    }
    
    post {
        always {
            script {
                dockerImage.remove() // Remove the Docker image after use
            }
        }
    }
}
