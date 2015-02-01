
Requirements for Test:
------------------

* Web application that allows user to create simple user objects (firstName & lastName). User object request should be sent via a REST/JSON interaction with the go server.
*  Go server Must Support for Postgresql DBMS.
*  Supports relevant views by using Revel Web Framework for Go
*  Go server must support REST API for creating user objects by using REST API Client (HTTP POST/PUT)
*  Extra Credit:  Secure REST API by using basic authentication


Components for test that must be provided:
------------------

*  Simple Revel user interface implementation where a user can self register their identity by submitting user name, password, first name, last name.
*  Go server that must accept REST/JSON requests from client and Revel implementation to create this user in the PostgresSQL DB.
*  Rest client that sends REST requests to server for identity creation. 
