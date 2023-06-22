# gRPC_APIs

## GoとgRPCを利用したバックエンドサーバー


## github workflow で使うトークンの設定方法
### トークンの生成
1. 右上のアイコンからSettings
1. 左下の "Developer settings"
1. 左の "Personal access tokens" の "Tokens (classic)"
1. 右上の "Generate new token" の "Generate new token (classic)"
1. "Note"でトークン名, "Expiration"で有効期限, "Select scopes"で権限を指定し"Generate token"で生成
<img width="960" alt="github_token" src="https://github.com/VolatileMint/gRPC_APIs/assets/85790876/773ce14d-a7cf-4bd7-82c9-0b93d7373b82">


### トークンの登録
1. 使いたいリポジトリのSettingsを開く
1. 左の"Secrets and variables"の"Actions"を開く
1. "New repository sercret"で、"Secret"に先ほどのトークンを張る、"Name"はworkflow内で使う名称
<img width="960" alt="トークン" src="https://github.com/VolatileMint/gRPC_APIs/assets/85790876/1d4f1797-b544-4578-bc46-a396a264c09f">
