FROM mysql:latest

# Set environment variables
ENV MYSQL_ROOT_PASSWORD=123asd123

# Copy SQL script to container
COPY ./migrations/create_tables.sql /docker-entrypoint-initdb.d/
COPY ./migrations/insert_sales.sql /docker-entrypoint-initdb.d/
COPY ./migrations/insert_supplies.sql /docker-entrypoint-initdb.d/

# Expose port 3306
EXPOSE 3306
