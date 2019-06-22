# Deploying Gofka using Ansible

## Docker for Mac users
```
sudo easy_install pip
sudo -H pip install docker==2.2.1 ansible==2.5.2 jinja2==2.9.6 couchdb==1.1 httplib2==0.9.2 requests==2.10.0
```

**Activate docker0 network**  
This is an optional step for local deployment. The OpenWhisk deployment via Ansible uses the docker0 network interface to deploy OpenWhisk and it does not exist on Docker for Mac environment.  

An expedient workaround is to add alias for docker0 network to loopback interface.  
```
sudo ifconfig lo0 alias 172.17.0.1/24  
```

## Deploying Using Kafka
```
ansible-playbook -i environments/local/hosts setup.yml
ansible-playbook -i environments/local kafka.yml -e mode=deploy
```

**Failed to import docker-py**  
After brew install ansible, the following lines are printed out:  
```
mkdir -p ~/Library/Python/2.7/lib/python/site-packages
echo '/usr/local/lib/python2.7/site-packages' > ~/Library/Python/2.7/lib/python/site-packages/homebrew.pth
```