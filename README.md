# SIMPLE CRUD API USING GO-SQL
A simple program to do Create, Read, Update and Delete `User` data consists of name and age. This program runs on localhost at port 3000. The host and the port depends on environment variables on `.env` file.
## CREATE
To create data, hit route `/people` using `POST` method with a body request consists of JSON:
```json
{
    "name": "",
    "age": 
}
```
## READ ALL DATA
To read all data, hit route `/people` using `GET` method.
## READ
To read a data, hit route `/person/{id}` using `GET` method with a query id.
## UPDATE
To update data, hit route `/person/{id}` using `UPDATE` method with a body request consists of JSON:
```json
{
    "name": "",
    "age": 
}
```
## DELETE
To delete data, hit route `/person` using `DELETE` method with a body request consists of JSON:
```json
{
    "id": ""
}
```