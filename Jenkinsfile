def remote = [:]
remote.name = "node-1"
remote.host = "18.142.43.34"
remote.allowAnyHosts = true

node {
    withCredentials([sshUserPrivateKey(credentialsId: 'ec2-mid-project-key', keyFileVariable: 'identity', passphraseVariable: '', usernameVariable: 'ubuntu')]) {
        remote.user = ubuntu
        remote.identityFile = identity
        stage("Connect EC2") {
            sshCommand remote: remote, command: 'whoami'
            sshCommand remote: remote, command: 'pwd'
        }
        stage("Pull Service") {
            sshCommand remote: remote, command: '''
                cd ci-cd-with-jenkins
                sudo git pull
            '''
        }
        stage("Backup Image") {
            sshCommand remote: remote, command: '''
                docker tag backend-service:latest backend-service:backup 
            '''
            sshCommand remote: remote, command: '''
                docker image ls | grep backend-service   
            '''
        }
        stage("Build a New Image") {
            sshCommand remote: remote, command: '''
                cd ci-cd-with-jenkins
                docker-compose build
            '''
            sshCommand remote: remote, command: '''
                docker image ls
            '''
        }
        stage("Down Service") {
            sshCommand remote: remote, command: '''
                cd ci-cd-with-jenkins
                docker-compose down
            '''
            sshCommand remote: remote, command: '''
                docker ps
            '''
        }
        stage("Up a New Service") {
            sshCommand remote: remote, command: '''
                cd ci-cd-with-jenkins
                docker-compose up -d
            '''
            sshCommand remote: remote, command: '''
                docker ps
            '''
        }
        stage("Remove Old Image") {
            sshCommand remote: remote, command: '''
                docker image rm backend-service:backup
            '''
            sshCommand remote: remote, command: '''
                docker image ls
            '''
            sshCommand remote: remote, command: '''
                df -h
            '''
        }
    }
}