node { 
    checkout scm
    stage('Commit message') {
        checkout scm
        sh "make commit-check"
        sh "env"
        sh "ls -alh"
        sh "git log"
    }
    stage('Dependency') {
        sh "ls && make restore"
    }
    stage('CodeGen') {
        sh "make codegen"
    }
    stage('Build') {
        sh "make all"
    }
    stage('Lint') {
        sh "make lint"
    }
}
