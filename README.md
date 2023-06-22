# gRPC_APIs

## GoとgRPCを利用したバックエンドサーバー


## github workflow で使うトークンの設定方法
### トークンの生成
1. 右上のアイコンからSettings
1. 左下の "Developer settings"
1. 左の "Personal access tokens" の "Tokens (classic)"
1. 右上の "Generate new token" の "Generate new token (classic)"
1. "Note"でトークン名, "Expiration"で有効期限, "Select scopes"で権限を指定し"Generate token"で生成

### トークンの登録
1. 使いたいリポジトリのSettingsを開く
1. 左の"Secrets and variables"の"Actions"を開く
1. "New repository sercret"で、"Secret"に先ほどのトークンを張る、"Name"はworkflow内で使う名称


