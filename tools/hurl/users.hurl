# -----------------------------------------------------------------
# Users
# -----------------------------------------------------------------

# Create user
POST http://{{HOST}}:{{PORT}}/api/v1/users
{
  "username": "{{USER_NAME}}"
}

HTTP/1.1 201

# Delete user
DELETE http://{{HOST}}:{{PORT}}/api/v1/users/{{USER_ID}}
HTTP/1.1 204

# Get user
GET http://{{HOST}}:{{PORT}}/api/v1/users/{{USER_ID}}
HTTP/1.1 200

