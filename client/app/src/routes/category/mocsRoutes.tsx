import { createRoute } from '../utils';
import { Moc,Moc2 } from '../exportedPages';
import SidebarLayout from 'src/layouts/SidebarLayout';
import { Navigate } from 'react-router-dom';

export const mocRoutes = [
  createRoute('mocs', <SidebarLayout />, [
    createRoute('', <Navigate to="moc" replace />),
    createRoute('moc', <Moc />),
    createRoute('moc2', <Moc2 />)
  ])
];


