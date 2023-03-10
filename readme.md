# file-serverfor 

for receive files posted

body is the steam of file.
save to hash.

麻了，写个Createfile都卡。


## upload file

```txt
POST /api/upload?path=:path/to/file
```

```sh
curl --data-binary @filename http://localhost:3000/
```


## get file

```txt
GET /path/to/file
```