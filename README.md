# oapi-calendar-sync-demo 🚀

## 概要

Larkカレンダー同期アプリケーション（Go言語で実装） 📅

## 参照URL

このデモは以下のURLを参考にしています：

[カレンダーイベント同期の準備](https://open.larksuite.com/document/home/calendar-event-sync/prepare)

## 機能

- 📋 カレンダーのリスト表示
- 📅 イベントのリスト表示
- 🔔 カレンダーのサブスクリプション追加

## 使用技術

| 技術 | 説明 |
|------|------|
| Go言語 | メイン言語 |
| Lark OpenAPI | Larkとの連携 |
| Gin | Webフレームワーク |
| Logrus | ロギング |
| Docker | コンテナ化 |
| TailScale | リモートアクセス |

## インストールと設定

1. リポジトリのクローン：
   ```bash
   git clone https://github.com/your-username/oapi-calendar-sync-demo.git
   cd oapi-calendar-sync-demo
   ```

2. 設定ファイルの編集：
   `conf/config.yml` を編集し、Larkアプリの認証情報を設定します。

3. アプリケーションの起動：
   ```bash
   docker compose up -d
   ```

4. TailScale Funnelでエンドポイント設置:
   ```bash
   tailscale funnel -bg --https=443 8089
   ```
