package constants

var Register string = `mutation Register(
	$first_name: String!, 
	$last_name: String, 
	$username: String, 
	$email: String!, 
	$password: String) {
  insert_users_one(object: {first_name: $first_name, last_name: $last_name, username: $username, email: $email, password: $password}) {
    id
    first_name
    last_name
    username
    email
  }
}`
var SearchUser string = `query SearchUser($loginText: String!) {
  users(where: {_or: [ {email: {_eq: $loginText}}, {username: {_eq: $loginText}}]}) {
    id
    first_name
    last_name
    username
    email
    password
  }}
`
var OldSearchUser string = `query SearchUser(
  $loginText: String!) {   
    search_user(args: { search: $loginText }) {     
      id     
      first_name     
      last_name     
      email     
      phone_number     
      password   } 
}`
