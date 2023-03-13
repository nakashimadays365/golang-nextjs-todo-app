local-start:
	APP_ENV=dev MYSQL_USER=user MYSQL_PASS=password MYSQL_HOST=localhost MYSQL_NAME=todo_db air

dev:
	cd web/todo-app/ && npm run dev

