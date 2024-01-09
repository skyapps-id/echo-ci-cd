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
                checkout([$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[url: 'https://github.com/skyapps-id/echo-ci-cd.git']]])
            }
        }
        
        stage('Build Docker Image') {
            steps {
                script {
                    def imageName = "echo-ci-cd:${BUILD_NUMBER}"
                    dockerImage = docker.build(registryName, "-f Dockerfile .")
                }
            }
        }
        
        // stage('Unit Test') {
        //     steps {
        //         script {
        //             dockerImage.inside {
        //                 sh 'go test ./...'
        //             }
        //         }
        //     }
        // }
        
        stage('Upload Image to ACR') {
            steps {
                script {
                    def imageName = "echo-ci-cd:${BUILD_NUMBER}"
                    dockerImage.tag(registryName)
                    docker.withRegistry( "http://${registryUrl}", registryCredential ) {
                        dockerImage.push('latest')
                    }
                }
            }
        }
    }
    
   /*  post {
        always {
            script {
                dockerImage.remove() // Remove the Docker image after use
            }
        }
    } */
}
