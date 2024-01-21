pipeline {
    agent any

    environment {
        serviceName = "echo-ci-cd"
        registryCredential = 'ACR'
        dockerImage = ''
        registryUrl = 'efishery.azurecr.io'
        authorName = ''
        authorEmail = ''
    }
    

    stages {
        stage ('Checkout') {
            steps {
                checkout scm
                script {
                    authorName = sh(returnStdout: true, script: "git log -1 --pretty=format:'%an'")
                    echo "${authorName}"
                    authorEmail = sh(returnStdout: true, script: "git log -1 --pretty=format:'%ae'")
                    echo "${authorEmail}"
                }
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
                    dockerImage = docker.build("${serviceName}--${ENV}:${BUILD_NUMBER}", "-f Dockerfile .")
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
                        string(name: 'DOCKER_TAG', value: env.BUILD_NUMBER),
                        string(name: 'AUTHOR_NAME', value: authorName),
                        string(name: 'AUTHOR_EMAIL', value: authorEmail)
                    ]
                }
            }
        }
    }
    
    post {
        always {
            script {
                echo "Cleaning"
                sh "docker rmi ${serviceName}--${ENV}:${env.BUILD_NUMBER}"
                sh "docker rmi ${registryUrl}/${serviceName}--${ENV}:${env.BUILD_NUMBER}"
            }
        }
        success {
            echo "Release Success"
        }
        failure {
            echo "Release Failed"
        }
        cleanup{
            echo "Clean up in post work space"
            cleanWs()
        }
    }
}
