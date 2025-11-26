cat > Jenkinsfile <<'EOF'
pipeline {
    agent any

    triggers {
        githubPush()
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build Docker image') {
            steps {
                sh '''
                echo "Construyendo imagen Docker..."
                docker build -t miapp:latest .
                '''
            }
        }

        stage('Enviar a AWS EC2 y desplegar') {
            steps {
                sshagent(['aws-ec2-docker']) {
                    sh '''
                    echo "Desplegando en EC2..."

                    ssh -o StrictHostKeyChecking=no ubuntu@13.60.186.175 '
                        echo "Actualizando repositorio..."
                        cd /home/ubuntu/Devops || git clone https://github.com/adurancar-oss/Devops.git && cd Devops
                        git pull origin main

                        echo "Construyendo imagen en EC2..."
                        docker build -t miapp:latest .

                        echo "Deteniendo contenedor actual si existe..."
                        docker stop miapp || true
                        docker rm miapp || true

                        echo "Iniciando contenedor..."
                        docker run -d --name miapp -p 80:80 miapp:latest
                    '
                    '''
                }
            }
        }
    }

    post {
        success {
            echo "Despliegue completado con éxito"
        }
        failure {
            echo "El despliegue falló"
        }
    }
}
EOF
