# -----------------------------------------------------------------
# Users
# -----------------------------------------------------------------


# Create user
POST {{protocol}}://{{hostname}}/api/v1/users
{
  "username": "{{user_name}}"
}

HTTP/1.1 201

[Asserts]
header "content-type" contains "application/json"
jsonpath "$.username" == "{{user_name}}"


# Delete user
DELETE {{protocol}}://{{hostname}}/api/v1/users/{{user_id}}

HTTP/1.1 204


# Get user
GET {{protocol}}://{{hostname}}/api/v1/users/{{user_id}}

HTTP/1.1 200

[Asserts]
header "content-type" contains "application/json"
jsonpath "$.username" == "{{user_name}}"
jsonpath "$.id" == {{user_id}}


# Update user
PUT {{protocol}}://{{hostname}}/api/v1/users/{{user_id}}
{
  "username": "{{user_name}}"
}

HTTP/1.1 204
