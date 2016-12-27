#!groovy

properties ([
    pipelineTriggers([cron('10 H/4 * * *')]),
    buildDiscarder(
        logRotator(
            artifactDaysToKeepStr: '',
            artifactNumToKeepStr: '',
            daysToKeepStr: '5', numToKeepStr: ''
        )
    ),
])

node {
    stage("Checkout") {
        git branch: "master",
            credentialsId: 'github-gianarb',
            poll: false,
            url: "git@github.com:gianarb/argh.git"
    }
    stage("Build") {
        sh 'echo $PATH'
        sh 'make clean'
        sh 'docker run --rm --name argh-${BUILD_NUMBER} -e GOPATH=/usr -e GOROOT=/usr/local/go -v /opt/jenkins/workspace/argh/:/usr/src/myapp -d -w /usr/src/myapp golang:1.7 sleep 500'
        sh 'docker exec argh-${BUILD_NUMBER} go get ./...'
        sh 'docker exec argh-${BUILD_NUMBER} go build -o ./build/argh .'
        sh 'docker rm -fv argh-${BUILD_NUMBER}'
        sh 'cat feeds.txt | ./build/argh generate ./docs'
        sh 'git commit -a -m "rebuilding site `date`"'
        sh 'git push origin master'
    }
}
