# -----------------------------------------------------------------
# Stories
# -----------------------------------------------------------------


# Create story
POST {{protocol}}://{{hostname}}/api/v1/stories
{
  "userId": {{user_id}},
  "description": "{{description}}",
  "points": "{{points}}",
  "status": "{{status}}",
  "title": "{{title}}",
  "type": "{{type}}"
}

HTTP/1.1 201

[Asserts]
header "content-type" contains "application/json"
jsonpath "$.userId" == {{user_id}}


# Delete story
DELETE {{protocol}}://{{hostname}}/api/v1/stories/{{story_id}}

HTTP/1.1 204


# Get story
GET {{protocol}}://{{hostname}}/api/v1/stories/{{story_id}}

HTTP/1.1 200

[Asserts]
header "content-type" contains "application/json"
jsonpath "$.id" == {{story_id}}


# List stories
GET {{protocol}}://{{hostname}}/api/v1/stories

HTTP/1.1 200


# Update story
PUT {{protocol}}://{{hostname}}/api/v1/stories/{{story_id}}
{
  "userId": {{user_id}},
  "description": "{{description}}",
  "points": "{{points}}",
  "status": "{{status}}",
  "title": "{{title}}",
  "type": "{{type}}"
}

HTTP/1.1 204
