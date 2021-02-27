# Restart Torrent Client

Restart torrent client with new download and upload speed settings.

**URL** : `/api/restart/downrate/{downrate}/uprate/{uprate}`

**Method** : `GET`

**Query Parameters** :

| Parameter  | Type    | Description                                                             |
| ---------- | ------- | ----------------------------------------------------------------------- |
| `downrate` | integer | Maximum download speed in Kbps. **Set 0 for unlimited download speed.** |
| `uprate`   | integer | Maximum upload speed in Kbps. **Set 0 to disable upload.**              |

## Success Response

**Code** : `200 OK`

**Main Object** :

| Name      | Type   | Description                                 |
| --------- | ------ | ------------------------------------------- |
| `success` | bool   | Indicates whether the query was successful. |
| `message` | string | Text message that describes the response.   |

## Examples

**Request** :

`GET http://localhost:9000/api/restart/downrate/4096/uprate/512`

**Success Response** :

```json
{
  "success": true,
  "message": "Restart torrent client."
}
```
