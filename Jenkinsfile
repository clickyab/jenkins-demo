node { 
    checkout scm
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
