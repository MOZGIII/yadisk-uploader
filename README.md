# YaDisk Uploader

Uploads files to Yandex.Disk.

## Usage

### Getting token

First get an OAuth token with Yandex.Disk REST API access.

You can create your own OAuth app here: https://oauth.yandex.ru/client/new

Then see here for how to get token: https://yandex.ru/dev/oauth/doc/dg/tasks/get-oauth-token.html

For more info, visit this page: https://yandex.ru/dev/oauth/

### Obtaining the app

```
go install github.com/MOZGIII/yadisk-uploader/cmd/yadisk-upload@master
```

### Using the app

To upload to disk:

```
yadisk-upload -token [your OAuth token] source/file.txt disk:/remote/file.txt
```

To upload to app directory:

```
yadisk-upload -token [your OAuth token] source/file.txt app:/remote/file.txt
```

### From source

```
git clone https://github.com/MOZGIII/yadisk-uploader.git
cd yadisk-uploader/cmd/yadisk-upload
go get ./...
go run main.go -token [your OAuth token] source/file.txt disk:/remote/file.txt
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

