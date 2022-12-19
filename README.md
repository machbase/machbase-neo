# machbase

Machbase server



# quick start

## install machbase

Download archived file for your platform from [release](https://github.com/MACHBASE/machbase/releases).

```sh
unzip machbase-vX.Y.Z-platform-arch-edition.zip
```

## run machbase

```sh 
./machbase
```

## create example table

```sh
curl -o - -X POST http://127.0.0.1:4058/db/query \
            --data-urlencode "q=create tag table TAGDATA (name varchar(200) primary key, time datetime basetime, value double summarized, jsondata json)"
```

## insert

```sh
curl -o - -X POST http://127.0.0.1:4058/db/write -H "Content-Type: application/json" \
    -d "@post-data.json"
```


- `post-data.json`

    ```json
        {
            "table": "TAGDATA",
            "data": {
                "columns":["name", "time", "value", "jsondata"],
                "rows": [
                    [ "my-car", 1670380342000000000, 1.0001, "{\"speed\":\"32.1kmh\",\"lat\":37.38906,\"lon\":127.12182}" ],
                    [ "my-car", 1670380343000000000, 2.0002, "{\"speed\":\"65.4kmh\",\"lat\":37.38908,\"lon\":127.12189}" ]
                ]
            }
        }
    ```

## query

 ```sh
curl -o - http://127.0.0.1:4058/db/query --data-urlencode "q=select * from TAGDATA"
```