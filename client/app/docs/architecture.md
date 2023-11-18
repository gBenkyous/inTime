# パッケージ構成
/usr/src/app  
├── README.md                 # プロジェクトの説明やセットアップ方法  
├── package.json              # プロジェクト依存関係とスクリプト  
├── yarn.lock                 # Yarn依存関係のロックファイル  
├── public                    # 静的ファイル（HTML, faviconなど）  
│   └── index.html  
├── src                       # ソースコード  
│   ├── App.tsx               # アプリケーションのルートコンポーネント  
│   ├── index.tsx             # エントリポイント  
│   ├── router.tsx            # アプリケーションのルーティングを管理  
│   ├── assets                # 静的資産（画像、スタイル、フォント）  
│   ├── components            # 再利用可能なUIコンポーネント  
│   ├── hooks                 # カスタムフック  
│   ├── layouts               # サイドバーなど共通ページレイアウト  
│   ├── pages                 # 各ページのメインビューなどルーティングと密接に関連するコンポーネント  
│   ├── store                 # 状態管理（Redux/MobXなど）  
│   ├── models                # データモデル  
│   ├── services              # 外部サービスのインターフェース、ビジネスロジック  
│   ├── utils                 # ユーティリティ関数  
│   ├── constants             # アプリケーション全体で使用される定数  
├── tsconfig.json             # TypeScript設定  
├── .env                      # 環境変数  
└── tests                     # テストファイル  
    ├── unit                  # ユニットテスト  
    └── integration           # 統合テスト  
