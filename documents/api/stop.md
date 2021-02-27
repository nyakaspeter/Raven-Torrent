# Stop Server

Stop the server and exit from the application.

**URL** : `/api/stop`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Main Object** :

| Name      | Type   | Description                                 |
| --------- | ------ | ------------------------------------------- |
| `success` | bool   | Indicates whether the query was successful. |
| `message` | string | Text message that describes the response.   |

## Examples

**Request** :

`GET http://localhost:9000/api/stop`

**Success Response** :

```json
{
  "success": true,
  "message": "Server stopped."
}
```
