pipeline {
    agent any

    environment {
        registryName = "echo-ci-cd"
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
                    dockerImage = docker.build("${registryName}:${BUILD_NUMBER}", "-f Dockerfile .")
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
                    echo "Triggering update manifest job"
                    build job: 'Job Deployment', parameters: [string(name: 'DOCKERTAG', value: env.BUILD_NUMBER),string(name: 'SVC_NAME', value: registryName),string(name: 'IMAGE_NAME', value: "${registryUrl}/${registryName}")]
                }
            }
        }
    }
    
    post {
        always {
            script {
                echo "Clear images"
                sh "docker rmi ${registryName}:${env.BUILD_NUMBER}"
                sh "docker rmi ${registryUrl}/${registryName}:${env.BUILD_NUMBER}"
            }
        }
    }
}
