node { 
    checkout scm
    stage('Commit message') {
        checkout scm
	def TARGE = sh(script: 'mktemp', returnStdout: true).trim()
        sh "echo $TARGE && git log > $TARGE"
	def out = readFile TARGE
	slackSend color: 'green', message: out
}}
