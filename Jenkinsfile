pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh 'mkdir -p $GOPATH/src/github.com/dtchanpura/cd-go; rsync -avz $WORKSPACE/* $GOPATH/src/github.com/dtchanpura/cd-go/'
        sh 'cd $GOPATH/src/github.com/dtchanpura/cd-go; make build'
      }
    }
    stage('Test') {
      steps {
        sh 'cd $GOPATH/src/github.com/dtchanpura/cd-go; make test'
      }
    }
  }
}