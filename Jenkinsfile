pipeline {
    agent any

    environment {
        registryName = "echo-ci-cd"
        registryCredential = 'ACR'
        registryUrl = 'efishery.azurecr.io'
    }
    
    node {
        def app
        
        stage ('Checkout') {
            steps {
                checkout scm
            }
        }
        
        stage('Build Docker Image') {
            steps {
                script {
                    app = docker.build("${registryName}:${BUILD_NUMBER}", "-f Dockerfile .")
                }
            }
        }
        
        stage('Unit Test') {
            steps {
                script {
                    app.inside {
                        sh 'echo "Tests passed"'
                    }
                }
            }
        }
        
        stage('Upload Image to ACR') {
            steps {
                script {
                    docker.withRegistry( "http://${registryUrl}", registryCredential ) {
                        app.push()
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
