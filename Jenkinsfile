pipeline {
    agent any

    environment {
        serviceName = "echo-ci-cd"
        registryCredential = 'ACR'
        dockerImage = ''
        registryUrl = 'efishery.azurecr.io'
    }
    
    stages {
        stage ('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Unit Test') {
            agent {
                docker {
                    image 'golang:1.19'
                }
            }
            environment { GOCACHE = "${WORKSPACE}" }
            steps {
                sh 'go version'
                sh 'go test ./...'            
            }
        }
        
        stage('Build Docker Image') {
            steps {
                script {
                    dockerImage = docker.build("${serviceName}:${BUILD_NUMBER}", "-f Dockerfile .")
                }
            }
        }
        
        stage('Upload Image to ACR') {
            steps {
                script {
                    docker.withRegistry( "http://${registryUrl}", registryCredential ) {
                        dockerImage.push()
                    }
                }
            }
        }

        stage('Trigger Manifest Update') {
            steps {
                script {
                    echo "Triggering update manifest job for ${ENV}"
                    build job: 'Job Deployment', parameters: [
                        string(name: 'ENV', value: ENV),
                        string(name: 'SVC_NAME', value: serviceName),
                        string(name: 'IMAGE_NAME', value: "${registryUrl}/${serviceName}"),
                        string(name: 'DOCKER_TAG', value: env.BUILD_NUMBER)
                    ]
                }
            }
        }
    }
    
    post {
        always {
            script {
                echo "Cleaning"
                sh "docker rmi ${serviceName}:${env.BUILD_NUMBER}"
                sh "docker rmi ${registryUrl}/${serviceName}:${env.BUILD_NUMBER}"
            }
        }
    }
}
