champak - A open source e-commerce solution by spree.
---
For ubuntu (16.04.1 LTS)

## Create deploy user
* Add user
```
useradd -s /bin/bash -m deploy
passwd -l deploy
```
* Upload your id_rsa.pub
```
scp ~/.ssh/id_rsa.pub deploy@www.change-me.com:/tmp
```

* Ssh no-password login 
```
su - deploy
mkdir ~/.ssh
chmod 700 ~/.ssh
cat /tmp/id_rsa.pub >> ~/.ssh/authorized_keys
```

* set no-password sudo
```
EDITOR=vim visudo
```
and add line:
```
deploy ALL=(ALL) NOPASSWD: ALL
```

## Install ruby
* install rbenv
```
sudo apt-get install -y git build-essential make libssl-dev libreadline-dev

git clone https://github.com/rbenv/rbenv.git ~/.rbenv
git clone https://github.com/rbenv/ruby-build.git ~/.rbenv/plugins/ruby-build
git clone https://github.com/rbenv/rbenv-vars.git ~/.rbenv/plugins/rbenv-vars
```

* Modify your ~/.zshrc file instead of ~/.bash_profile
```
echo 'export PATH=$HOME/.rbenv/bin:$PATH' >> ~/.bashrc
echo 'eval "$(rbenv init -)"' >> ~/.bashrc
```
* ruby(need re-login)
```
rbenv install 2.2.5
rbenv local 2.2.5
gem install bundler
```

## Deployment

* Prepare
```
```
* Create database
```
psql -U postgres
CREATE DATABASE db-name WITH ENCODING = 'UTF8';
CREATE USER user-name WITH PASSWORD 'change-me';
GRANT ALL PRIVILEGES ON DATABASE db-name TO user-name;
```
* Tasks
```
```

## Issues

* 'Peer authentication failed for user', open file "/etc/postgresql/9.5/main/pg_hba.conf" change line:
```
    local   all             all                                     peer
```

to:
```
    local   all             all                                     md5
```


 

