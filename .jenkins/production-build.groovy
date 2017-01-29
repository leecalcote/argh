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
            url: "https://github.com/gianarb/argh"
    }
    stage("Build") {
        sh('make clean')
        sh('docker run --rm --name argh-${BUILD_NUMBER} -e GOPATH=/usr -e GOROOT=/usr/local/go -v ${PWD}:/usr/src/myapp -d -w /usr/src/myapp golang:1.7 sleep 500')
        sh('docker exec argh-${BUILD_NUMBER} go get ./...')
        sh('docker exec argh-${BUILD_NUMBER} env GOOS=linux GOARCH=386 go build -o ./build/argh .')
        sh('docker rm -fv argh-${BUILD_NUMBER}')
        sh('cat feeds.txt | ./build/argh generate ./docs')
        sh('git commit -a -m "rebuilding site `date`"')
        withCredentials([[
                        $class: 'UsernamePasswordMultiBinding',
                        credentialsId: 'github-gianarb',
                        usernameVariable: 'GIT_USERNAME',
                        passwordVariable: 'GIT_PASSWORD'
                ]]) {
            sh('git push https://${GIT_USERNAME}:${GIT_PASSWORD}@github.com/gianarb/argh')
        }
    }
}
