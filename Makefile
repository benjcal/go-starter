init-db:
	cd db; \
	psql -U postgres -c 'DROP DATABASE go' -c 'CREATE DATABASE go'; \
	psql -U postgres go -f init.sql
