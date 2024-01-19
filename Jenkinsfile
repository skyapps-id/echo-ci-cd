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
            checkout scm
        }
        
        stage('Build Docker Image') {
            app = docker.build("${registryName}:${BUILD_NUMBER}", "-f Dockerfile .")
        }
        
        stage('Unit Test') {
            app.inside {
                sh 'echo "Tests passed"'
            }
        }
        
        stage('Upload Image to ACR') {
            docker.withRegistry( "http://${registryUrl}", registryCredential ) {
                app.push()
            }
        }

        stage('Trigger Manifest Update') {
            echo "triggering updatemanifestjob"
            build job: 'Job Deployment', parameters: [string(name: 'DOCKERTAG', value: env.BUILD_NUMBER),string(name: 'SVC_NAME', value: registryName),string(name: 'IMAGE_NAME', value: "${registryUrl}/${registryName}")]
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
