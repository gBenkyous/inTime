# 手順書

## 崩しTDDで行う
崩しテスト駆動開発  
1. 動くそれっぽいアプリケーションを作る
2. テストケースを作る
3. テストに通るアプリケーションを作る
4. リファクタリング

## 画面作成場所
pagesに画面を作成する  
既存のfileこぴって貼り付けると楽です。  
既存フレームワークをこぴって貼り付けましょう。そこから修正するといいと思います。
例：  
mocカテゴリーのmocログイン画面であれば下記作成し、indexを作成する。  
/usr/src/app/src/pages/mocs/MocLogin  
-index.tsx  

必要に応じてカテゴリーは追加してください。ただし、カテゴリーを追加した場合には、不要カテゴリを作らないように、レビュー時のレビュワーを全員に設定してください。  

## ルーティング
### ページコンポーネントエクスポート
routes/category/exportedPages.tsにページコンポーネントをインポートして、簡単に呼び出せるようにしておく。  

### ルーティング作成
/usr/src/app/src/routes/category配下にpagesカテゴリーと同様のカテゴリルーティングを作成する。（既存のものがあれば追記）
ルーティング名は[pagesカテゴリ名]Routesとする。
例：pagesカテゴリーが下記のように
/usr/src/app/src/pages/mocs/MocLogin
mocsならば
routes/category/mocsRoutes.tsxを作成。

### ルーティング設定
routes/index.tsxにルーティングを設定する。

## 画面画面内コンテンツ作成
ググったり、ほかのやつを参考に作っていきましょう。  
基本的には  
- [コンポーネントに関して](https://qiita.com/Hashimoto-Noriaki/items/95d9fe027d169ce74218)
- [値の受け渡しに関して](https://qiita.com/daishiman/items/0f0592dc60cc844bf99b)



APIから渡される想定のデータmodelフォルダの下とかにjsonファイル作ってそれを呼び出してテストしながら開発する