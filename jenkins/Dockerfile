FROM jenkinsci/jenkins

USER root

# Adding docker to image
ADD "https://get.docker.com/builds/Linux/x86_64/docker-latest.tgz" /tmp/
RUN tar xzf /tmp/ && mv /tmp/docker/docker /usr/bin/docker

# Adding docker-compose to image
ADD "https://github.com/docker/compose/releases/download/1.23.1/docker-compose-$(uname -s)-$(uname -m)" /usr/local/bin/docker-compose
RUN chmod 755 /usr/local/bin/docker-compose

RUN usermod -a -G sudo jenkins
USER jenkins
