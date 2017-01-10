
cap=cd champak && git pull && bundle exec cap production

all: deploy

deploy:
	@echo 'Deploy a new release'
	$(cap) deploy

rollback:
	@echo 'Rollback to previous release'
	$(cap) deploy:rollback


backup:
	@echo 'Backup data and sync to localhost'
	$(cap) app:pull
