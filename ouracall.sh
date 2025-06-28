#!/usr/bin/env bash

# curl -X GET https://api.ouraring.com/v2/usercollection/personal_info \
# curl --location --request GET 'https://api.ouraring.com/v2/usercollection/daily_activity?start_date=2025-01-31&end_date=2025-06-28' \
curl --location --request GET 'https://api.ouraring.com/v2/usercollection/daily_activity?start_date=2025-06-01&end_date=2025-06-28' \
-H "Authorization: Bearer $OURA_TOKEN"
