import React, { ReactNode } from 'react';
import { RouteObject } from 'react-router-dom';

export const createRoute = (path: string, element: ReactNode, childrenRoutes?: RouteObject[]): RouteObject => ({
  path,
  element: element,
  children: childrenRoutes
});

export const createSubRoute = (path: string, childrenRoutes: RouteObject[]): RouteObject => ({
  path,
  children: childrenRoutes
});
