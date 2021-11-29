pipeline {
    agent any

    stages {
        stage('build') {
            steps {
                sh "PATH=$PATH:/usr/local/go/bin"
                sh "go build go-http-server.go"
            }
        }
        stage('zip') {
            steps {
                sh "zip go-http-server.zip go-http-server"
            }
        }
        stage('upload') {
            steps {
                nexusArtifactUploader{
                    nexusVersion: "nexus3"
                    protocol("http")
                    nexusUrl('nexus.shakhmin.ru:8081')
                    groupId('bin')
                    version('1')
                    repository('zip')
                    credentialsId('7481cfc2-21c5-473c-a509-dd2270d5c954')
                    artifact {
                        artifactId('go-http-server')
                        type('zip')
                        // classifier('')
                        file('go-http-server.zip')
                    }
                }
            }
        }
    }
}
