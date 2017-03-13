node { 
    checkout scm
    stage('Commit message') {
        sh "make commit-check"
        sh "env"
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
