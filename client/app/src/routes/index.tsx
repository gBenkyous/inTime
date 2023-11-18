import { Navigate, RouteObject } from 'react-router-dom';
import { dashboardRoutes } from './dashboardRoutes';
import { applicationRoutes } from './applicationRoutes';
import { componentRoutes } from './componentRoutes';
import { createRoute, createSubRoute } from './utils';

import BaseLayout from 'src/layouts/BaseLayout';
import { Overview, Status404, Status500, StatusComingSoon, StatusMaintenance } from './lazyLoadedPages';

const routes: RouteObject[] = [
  {
    path: '',
    element: <BaseLayout />,
    children: [
      createRoute('/', <Overview />),
      createRoute('overview', <Navigate to="/" replace />),
      createSubRoute('status',[
        createRoute('', <Navigate to="404" replace />),
        createRoute('404', <Status404 />),
        createRoute('500', <Status500 />),
        createRoute('maintenance', <StatusMaintenance />),
        createRoute('coming-soon', <StatusComingSoon />)
      ]), 
      { path: '*', element: <Status404 /> }     
    ]
  },
  ...dashboardRoutes,
  ...applicationRoutes,
  ...componentRoutes
];

export default routes;

/* 基本的なルーティング例
const routes: RouteObject[] = [
  {
    path: '/',
    element: <BaseLayout />,
    children: [
      { path: 'overview', element: <Overview /> },
      { 
        path: 'status/404',
        element: <Status404 />,
        children: [
          { path: 'aaa', element: <aaa404 /> },
      },
    ]
  },
  { path: '*', element: <Status404 /> }
];

*/
