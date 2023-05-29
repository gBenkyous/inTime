/**
=========================================================
* Argon Dashboard 2 MUI - v3.0.0
=========================================================

* Product Page: https://www.creative-tim.com/product/argon-dashboard-material-ui
* Copyright 2022 Creative Tim (https://www.creative-tim.com)

Coded by www.creative-tim.com

 =========================================================

* The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
*/

import React from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import App from "./App";
// ここでの App は、src/App.tsx で定義されています。
// src/App.tsx で定義されている App は、src/App.tsx で定義されている、App コンポーネントです。

// import とはなんですか？
// import は、モジュール、特に別のファイルにあるコードを読み込むための構文です。
// 例えば、あるファイルで定義された関数を、別のファイルで使いたい場合、
// その関数が定義されているファイルを import することで、その関数を使うことができます。

// Soft UI Context Provider
import { ArgonControllerProvider } from "./context";

// react-perfect-scrollbar component
import PerfectScrollbar from "react-perfect-scrollbar";
// fromとはなんですか？
// from は、import とセットで使われる構文です。
// formは、import するファイルのパスを指定します。
// 例えば、import PerfectScrollbar from "react-perfect-scrollbar" という構文があったとき、
// "react-perfect-scrollbar" は、react-perfect-scrollbar フォルダの index.js ファイルを指します。

// module "/usr/src/app/node_modules/react-perfect-scrollbar/lib/index"とは
// node_modules とは、プロジェクトで使用するパッケージがインストールされるフォルダです。
// 例えば、react-perfect-scrollbar パッケージをインストールすると、
// node_modules/react-perfect-scrollbar フォルダが作成され、その中にパッケージのファイルが格納されます。
// node_modules フォルダは、プロジェクトのルートフォルダに作成されます。


// react-perfect-scrollbar styles
import "react-perfect-scrollbar/dist/css/styles.css";
// importで {} を使うとはなんですか？逆に使わないとどうなりますか？
// {} は、import するファイルの中から、特定の変数や関数を読み込むための構文です。
// 例えば、import { ArgonControllerProvider } from "context" という構文があったとき、
// { ArgonControllerProvider } は、context フォルダの index.js ファイルから、
// ArgonControllerProvider という変数を読み込むための構文です。
// 逆に、import ArgonControllerProvider from "context" という構文だと、
// context フォルダの index.js ファイルから、default で export された変数を読み込むことになります。

// 下記は、index.html に書かれている <div id="root"></div> を取得しています。
const container = document.getElementById("root");
const root = createRoot(container!);

root.render(
  <BrowserRouter>
  
    <ArgonControllerProvider>
      <PerfectScrollbar>
        <App />
      </PerfectScrollbar>
    </ArgonControllerProvider>
  </BrowserRouter>
);
