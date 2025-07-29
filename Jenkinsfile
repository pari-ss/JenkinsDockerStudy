pipeline {
    agent any // Specifies that any available agent can execute the pipeline

    tools {
        // Configures the Go tool. 'go-1.x.x' should match a Go installation configured in Jenkins' Global Tool Configuration.
        go 'go1.22.4' 
    }

    environment {
        // Sets environment variables specific to your Go project, if needed.
        // For example, to enable Go modules:
        GO111MODULE = 'on'
        //DOCKERHUB_CREDENTIALS = credentials('pariss')
        DOCKER_HUB_CREDENTIAL_ID = 'pariss'
        DOCKER_IMAGE = 'pariss/test-app' 
    }

    stages {
        stage('Checkout') {
            steps {
                // Checks out the source code from a version control system (e.g., Git).
                echo "checking out repo"
                git url: 'https://github.com/pari-ss/JenkinsDockerStudy.git', branch: 'master'
            }
        }

        stage('Run Docker Build') {
            steps {
                // Builds the Go application.
                //sh 'go build -o jenkinsdockerstudy .' 
                script{
                    echo "starting docker build"
                    sh "docker build -t ${DOCKER_IMAGE}:latest ."
                    echo "docker built successfully"
               }
            }
        }

        stage('push to docker hub'){
                steps{
                    echo "pushing to docker hub"
                    script{
                        docker.withRegistry('https://index.docker.io/v1/', 'pariss'){
                            echo "while image push"
                            docker.image("${DOCKER_IMAGE}:latest").push()
                        }
                    }
                    echo "done"
                }
        }    
    }
   // post {
        // Actions to perform after the pipeline completes, regardless of success or failure.
      //  always {
       //     cleanWs() // Cleans up the workspace on the agent
      //  }
        // Add other post-build actions like notifications, archiving artifacts, etc.
   // }
}