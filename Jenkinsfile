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
    }

    stages {
        stage('Checkout') {
            steps {
                // Checks out the source code from a version control system (e.g., Git).
                git url: 'https://github.com/pari-ss/JenkinsDockerStudy.git', branch: 'master'
            }
        }

        stage('Build') {
            steps {
                // Builds the Go application.
                sh 'go build -o jenkinsdockerstudy .' 
            }
        }

       /*  stage('Test') {
            steps {
                // Runs Go tests.
                sh 'go test -v ./...' 
            }
        } */

         stage('Deploy') {
            when {
                expression { currentBuild.result == null || currentBuild.result == 'SUCCESS' }
            }
            steps {
                echo 'Deploying the application...'
                sh 'go run main.go' // Replace with your deployment command
                // Add your deployment commands here, e.g.,
                // sh 'kubectl apply -f k8s/deployment.yaml'
                // sh 'scp target/myapp.war user@server:/opt/tomcat/webapps/'
            }
         }
        // Add more stages as needed for your CI/CD workflow, e.g., 'Docker Build', 'Deploy', etc.
    }

    /* post {
        // Actions to perform after the pipeline completes, regardless of success or failure.
        always {
            cleanWs() // Cleans up the workspace on the agent
        }
        // Add other post-build actions like notifications, archiving artifacts, etc.
    } */
}
