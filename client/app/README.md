# 情報一覧

## Create React Appで始める

### 利用可能なスクリプト

プロジェクトディレクトリで、以下のコマンドを実行できます。

#### `yarn install`

必要なライブラリをインストールします。

#### `yarn start`

開発モードでアプリを実行します。\
ブラウザで[http://localhost:3000](http://localhost:3000)を開いてください。  
dockerで開いている場合3000ではなく80ポートが割り当てられているため[http://localhost](http://localhost)で開いてください。

編集するとページがリロードされます。\
コンソールにはlintエラーも表示されます。

#### `yarn test`

インタラクティブなウォッチモードでテストランナーを起動します。\
詳細は[テストの実行](https://facebook.github.io/create-react-app/docs/running-tests)のセクションを参照してください。

#### `yarn run build`

`build`フォルダにアプリを本番用にビルドします。\
Reactを本番モードで正しくバンドルし、最高のパフォーマンスを得るためにビルドを最適化します。

ビルドは圧縮され、ファイル名にはハッシュが含まれます。\
アプリはデプロイする準備ができました！

詳細は[デプロイ](https://facebook.github.io/create-react-app/docs/deployment)のセクションを参照してください。

#### `yarn run eject`

**注意: これは一方通行の操作です。一度`eject`すると、元に戻ることはできません！**

ビルドツールや設定の選択に満足できない場合は、いつでも`eject`できます。このコマンドは、プロジェクトから単一のビルド依存関係を削除します。

代わりに、すべての設定ファイルと間接的な依存関係（webpack、Babel、ESLintなど）をプロジェクトに直接コピーして、完全に制御できるようにします。`eject`以外のすべてのコマンドは引き続き動作しますが、コピーされたスクリプトを指すようになりますので、調整できます。この時点であなたは自分でやるしかありません。

`eject`を使わなくても構いません。厳選された機能セットは、小規模から中規模のデプロイに適しており、この機能を使う義務を感じる必要はありません。しかし、このツールは、あなたがそれをカスタマイズできる準備ができたときに役に立たないということは理解しています。

### もっと詳しく知る

[Create React Appのドキュメント](https://facebook.github.io/create-react-app/docs/getting-started)で詳しく学ぶことができます。

Reactを学ぶには、[Reactのドキュメント](https://reactjs.org/)をチェックしてください。


## 技術スタック
- react [Create React App](https://github.com/facebook/create-react-app)
- rendux [Redux](https://redux.js.org/)と[Redux Toolkit](https://redux-toolkit.js.org/)
- MUI(旧Material-UI) yarn add @mui/material @emotion/react @emotion/styled @mui/icons-material
  - テンプレート機能を活用
    - https://www.creative-tim.com/product/argon-dashboard-material-ui?AFFILIATE=128200&ref=admin-dashboards.com
    - 下記は参考用
      - https://github.com/devias-io/material-kit-react
- 依存関係管理 yarn
- 
