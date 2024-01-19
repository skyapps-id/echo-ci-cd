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
        
        stage('Build Docker Image') {
            steps {
                script {
                    dockerImage = docker.build("${registryName}:${BUILD_NUMBER}", "-f Dockerfile .")
                }
            }
        }
        
        stage('Unit Test') {
            steps {
                script {
                    dockerImage.inside {
                        // sh 'go test ./...'
                        echo "Tests passed"
                    }
                }
            }
        }
        
        stage('Upload Image to ACR') {
            steps {
                script {
                    // dockerImage.tag("${BUILD_NUMBER}")
                    docker.withRegistry( "http://${registryUrl}", registryCredential ) {
                        dockerImage.push()
                    }
                }
            }
        }

        stage('Trigger Manifest Update') {
            steps {
                script {
                    echo "triggering updatemanifestjob"
                    build job: 'Job Deployment', parameters: [string(name: 'DOCKERTAG', value: env.BUILD_NUMBER),string(name: 'SVC_NAME', value: registryName),string(name: 'IMAGE_NAME', value: "${registryUrl}/${registryName}")]
                }
            }
        }
    }
    
    /* post {
        always {
            script {
                dockerImage.remove() // Remove the Docker image after use
            }
        }
    } */
}
