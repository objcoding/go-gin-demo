#!groovy
pipeline {
    agent any

    environment {
        GIT_REPO = "https://github.com/objcoding/go-gin-demo.git"
        BUILD_IMAGE_SCRIPT_PATH = "build.sh"
    }
    stages {
        stage('获取代码') {
            when {
                anyOf {
                    branch 'sit'
                    branch 'qa'
                }
            }
            steps {
                deleteDir()
                git([url: "${GIT_REPO}", branch: "${BRANCH_NAME}"])
            }
        }

        stage('编译代码') {
            when {
                anyOf {
                    branch 'sit'
                    branch 'qa'
                }
            }
            steps {
                sh "mvn -U -am clean package -DskipTests"
            }
            post {
                failure {
                    mail to: 'zhangchenghui.dev@gmail.com',
                         subject: "jenkins 流水线失败",
                         body: "编译代码失败，请尽快使用以下git命令进行代码回滚，或尽快修复错误并提交：\n" +
                                    "git revert HEAD\n" +
                                    "git push\n"
                }
            }
        }

        stage('构建镜像') {
            when {
                anyOf {
                    branch 'sit'
                    branch 'qa'
                }
            }
            steps {
                sh "sh ${BUILD_IMAGE_SCRIPT_PATH} ${BRANCH_NAME}"
            }
            post {
                failure {
                    mail to: 'zhangchenghui.dev@gmail.com',
                         subject: "代码构建失败",
                         body: "代码构建失败，请尽快使用登陆jenkins查看错误日志！"
                }
            }
        }

        stage('更新sit服务') {
            when {
                anyOf {
                    branch 'sit'
                }
            }
            steps {
                sh " '/deploy.sh' ${active} "
            }
        }


        stage('更新qa服务') {
            when {
                anyOf {
                    branch 'qa'
                }
            }
            steps {
                sh " '/deploy.sh' ${active} "
            }
        }
    }
}