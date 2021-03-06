SHELL := /bin/bash
SECRETS := $(shell readlink -f ./.env)

all: alertmanager.yml prometheus.yml blackbox.yml
	docker-compose up -d

alertmanager.yml: alertmanager.yml.in
	@test -f $(SECRETS) ; set -a ; source $(SECRETS) ; envsubst < $< > $@

prometheus.yml: alert.rules.yml

clean:
	rm alertmanager.yml

stop:
	docker-compose stop

down: clean
	docker-compose down -v
