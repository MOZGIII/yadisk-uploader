# YaDisk Uploader

Uploads files to Yandex.Disk.

## Usage

### Getting token

First get an OAuth token with Yandex.Disk REST API access.

You can create your own OAuth app here: https://tech.yandex.ru/oauth/

Then see here for how to get token: https://tech.yandex.ru/oauth/doc/dg/tasks/get-oauth-token-docpage/

### Using the app

To upload to disk:

```
yadisk-uploader -token [your OAuth token] source/file.txt disk:/remote/file.txt
```

To upload to app directory:

```
yadisk-uploader -token [your OAuth token] source/file.txt app:/remote/file.txt
```

### From source

```
git clone https://github.com/MOZGIII/yadisk-uploader.git
cd yadisk-uploader/src
go get ./...
go run main.go -token [your OAuth token] source/file.txt disk:/remote/file.txt
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

