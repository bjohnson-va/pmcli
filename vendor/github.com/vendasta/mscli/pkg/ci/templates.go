package ci

const (
	//JenkinsfileTemplate is the template for the Jenkinsfile to generate
	JenkinsfileTemplate = `#!groovy
@Library('shared-libraries') _

properties ([
    buildDiscarder(logRotator(artifactNumToKeepStr: '5', daysToKeepStr: '15')),
    disableConcurrentBuilds()
])

def parameters = [:]
parameters['label'] = "[[.Name]].${env.BRANCH_NAME}.${env.BUILD_NUMBER}".replaceAll("[^a-zA-Z0-9]+","_")
parameters['projectRoot'] = '[[.Name]]'
parameters['buildCommands'] = ["mscli app test", "mscli app vet", "mscli app lint"]
parameters['namespace'] = "[[.Name]]"

mscliNode(parameters) {}
`
)
