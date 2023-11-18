import { createRoute, createSubRoute } from './utils';
import { Transactions, UserProfile, UserSettings } from './lazyLoadedPages';
import SidebarLayout from 'src/layouts/SidebarLayout';
import { Navigate } from 'react-router-dom';

export const applicationRoutes = [
  createRoute('management', <SidebarLayout />, [
    createRoute('', <Navigate to="transactions" replace />),
    createRoute('transactions', <Transactions />),
    createSubRoute('profile', [
      createRoute('', <Navigate to="details" replace />),
      createRoute('details', <UserProfile />),
      createRoute('settings', <UserSettings />)
    ])
  ])
];
