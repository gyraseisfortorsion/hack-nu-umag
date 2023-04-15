.SILENT:

all:
	build
	run

build: 
	docker build -t my-mysql-image .
run:
	docker run -d --name my-mysql-cont -p 3306:3306 my-mysql-image
	

	