# -----------------------------------------------------------------
# Root
# -----------------------------------------------------------------


# API OK
GET {{protocol}}://{{hostname}}
HTTP 200


# Health Check
GET {{protocol}}://{{hostname}}/health
HTTP 200

[Asserts]
header "content-type" contains "application/json"
jsonpath "$.db.message" == "It's healthy"
jsonpath "$.db.status" == "up"
