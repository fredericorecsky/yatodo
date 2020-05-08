#!/usr/bin/env bash
#cd "$(dirname "$0")"
#source bash_tap.sh

HOST=localhost
PORT=9000
USER_NAME="testuser_"$RANDOM

echo "Create user , retrieve the token ( put on a variable )"
TOKEN=`/usr/bin/env curl -s -X POST $HOST:$PORT/users -d "{ \"username\": \"${USER_NAME}\" }"  |jq -r '.SecretToken'`
echo "SecretToken: $TOKEN"

# echo"Duplicate user"
#/usr/bin/env test -s -X POST $HOST:$PORT/users -d "{ \"username\": \"${USER_NAME}\" }"

echo "create a new todo list"
LIST_ID=`/usr/bin/env curl -s -H "Token:${TOKEN}" -X POST $HOST:$PORT/todolists -d "{ \"Name\": \"TODO TEST\" }" |jq -r '.ID'`
echo "ListID $LIST_ID"
/usr/bin/env curl -s -H "Token:${TOKEN}" -X POST $HOST:$PORT/todolists -d "{ \"Name\": \"TODO TEST 2\" }" >/dev/null

echo "get all todo lists from this user"
/usr/bin/env curl -s -H "Token:${TOKEN}" localhost:9000/todolists  |jq

echo "get a given todolist ( $LIST_ID ) from this user"
/usr/bin/env curl -s -H "Token:${TOKEN}" localhost:9000/todolists/$LIST_ID  |jq

echo "create a new item"
ITEM_ID=`/usr/bin/env curl -s -H "Token:${TOKEN}" -X POST localhost:9000/todolists/$LIST_ID/items -d "{ \"Description\":\"ITEM TEST\" }"  |jq -r '.ID'`
echo "ItemID $ITEM_ID"

echo "More one item"
/usr/bin/env curl -s -H "Token:${TOKEN}" -X POST localhost:9000/todolists/$LIST_ID/items -d "{ \"Description\":\"ITEM TEST 2\" }"  |jq -r '.ID'

echo "Get all items from this list"
/usr/bin/env curl -s -H "Token:${TOKEN}" localhost:9000/todolists/$LIST_ID/items  |jq

echo "create a new label for this item"
LABEL_ID=`/usr/bin/env curl -s -H "Token:${TOKEN}" -X POST localhost:9000/todolists/$LIST_ID/items/$ITEM_ID/labels -d '{ "Label":"semaforo" }' | jq -r '.ID'`

echo "create a new comment"
/usr/bin/env curl -s -H "Token:${TOKEN}" -X POST localhost:9000/todolists/$LIST_ID/items/$ITEM_ID/comments -d '{ "Comment":"Legal" }' | jq

echo "add a due date to item $ITEM_ID"
/usr/bin/env curl -s -H "Token:${TOKEN}" -X POST localhost:9000/todolists/$LIST_ID/items/$ITEM_ID/dueto -d '{ "DueTo":"2020-05-01" }' | jq

echo "get the item $ITEM_ID"
/usr/bin/env curl -s -H "Token:${TOKEN}" localhost:9000/todolists/$LIST_ID/items/$ITEM_ID | jq

echo "delete a due date on item $ITEM_ID"
/usr/bin/env curl -s -H "Token:${TOKEN}" -X DELETE localhost:9000/todolists/$LIST_ID/items/$ITEM_ID/dueto  | jq

echo "get the item $ITEM_ID"
/usr/bin/env curl -s -H "Token:${TOKEN}" localhost:9000/todolists/$LIST_ID/items/$ITEM_ID | jq

echo "delete label $LABEL_ID on item $ITEM_ID"
/usr/bin/env curl -s -H "Token:${TOKEN}" -X DELETE localhost:9000/todolists/$LIST_ID/items/$ITEM_ID/labels/$LABEL_ID  | jq

echo "get the item $ITEM_ID"
/usr/bin/env curl -s -H "Token:${TOKEN}" localhost:9000/todolists/$LIST_ID/items/$ITEM_ID | jq

echo "Delete the $ITEM_ID"
/usr/bin/env curl -s -H "Token:${TOKEN}" -X DELETE localhost:9000/todolists/$LIST_ID/items/$ITEM_ID | jq

echo "Get all items from this list"
/usr/bin/env curl -s -H "Token:${TOKEN}" localhost:9000/todolists/$LIST_ID/items  |jq