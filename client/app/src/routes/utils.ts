import React, { ReactNode } from 'react';
import { RouteObject } from 'react-router-dom';

/**
 * ルーティングを作成する関数
 * @param path パス
 * @param element ページの要素
 * @param childrenRoutes 子ルート
 * @returns 
 */
export const createRoute = (path: string, element: ReactNode, childrenRoutes?: RouteObject[]): RouteObject => ({
  path,
  element: element,
  children: childrenRoutes
});

/**
 * サブルートを作成する関数
 * /components/aaaa
 * /components/bbbb などの場合にcomponentsにはアクセスしないが配下のパスを設定したい場合に使用する
 * @param path パス
 * @param childrenRoutes 子ルート
 * @returns 
 */
export const createSubRoute = (path: string, childrenRoutes: RouteObject[]): RouteObject => ({
  path,
  children: childrenRoutes
});
