UpDb: 
	docker build -t my-mysql-image .
rundDB:
	docker run -d --name my-mysql-cont -p 3306:3306 my-mysql-image
down: docker stop my-mysql-cont
	docker system prune
buildApp: 
	docker build -t back ./Umag
runApp: 
	docker run -d --name backend1 -p 8080:8080  back



	

	