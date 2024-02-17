import { createRoute } from '../utils';
import { Dashboard, Crypto, Messenger } from '../exportedPages';
import SidebarLayout from 'src/layouts/SidebarLayout';
import { Navigate } from 'react-router-dom';

export const dashboardRoutes = [
  createRoute('dashboards', <SidebarLayout />, [
    createRoute('', <Navigate to="crypto" replace />),
    createRoute('crypto', <Crypto />),
    createRoute('messenger',<Messenger />)
  ])
];
