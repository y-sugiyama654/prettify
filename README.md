# Prettify

Prettifyは、Go言語で書かれたシンプルで使いやすいファイル整形CLIツールです。このツールは、コードを美しく整形し可読性を向上させることを目的としています。

## 特徴

- コードの自動整形
- シンプルなコマンドラインインターフェース
- 高速な処理速度

## インストール

以下の手順でPrettifyをインストールできます。

1. リポジトリをクローンします。

   ```bash
   git clone https://github.com/y-sugiyama654/prettify.git
   ```

2. ディレクトリに移動します。

   ```bash
   cd prettify
   ```

3. Goモジュールを初期化します。

   ```bash
   go mod tidy
   ```

## 使い方

Prettifyを使用するには、以下のコマンドを実行します。

```bash
// JSONフォーマット 
go run main.go -type json test.json

// YAMLフォーマット
go run main.go -type yaml test.yaml

// XMLフォーマット
go run main.go -type toml test.toml
```


