pipeline {
    
    agent any
    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub')
    }

    stages {
        stage('checkout') {
            steps {
               checkout scmGit(branches: [[name: '*/dev']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/MarlonDeOcampo/SimpleCrudGolang.git']])
            }
        }
        stage('build image') {
            steps {
                script {
                    sh "docker build -t alhon05/golangsimplecrud:v$BUILD_ID ."
                }
            }
        }
        stage('Login to Docker Hub') {      	
            steps{                
                    sh 'docker --version'
                    sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'                		
                    echo 'Login Completed'      
                }           
        }
        stage('upload image') {
            steps {
                sh "docker push alhon05/golangsimplecrud:v$BUILD_ID"
            }
        }
        stage('redeploy stack') {
            steps {
                sshPublisher(
                    publishers: [
                        sshPublisherDesc(
                            configName: "marlon", 
                            transfers: [sshTransfer(
                                execTimeout: 120000,
                                execCommand: "cd ~/Documents/SimpleCrudGolang;VERSION=$BUILD_ID docker stack deploy -c stack-main-global.yml main2"
                            )]
                        ) 
                    ]
                )
            }
        } 
    }
    post {
        always {
            sh 'docker logout'
            echo 'Job Completed...'   
        }
    }
}





           
	