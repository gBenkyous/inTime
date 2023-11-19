import { createRoute } from '../utils';
import { Buttons, Modals, Accordions, Tabs, Badges, Tooltips, Avatars, Cards, Forms } from '../lazyLoadedPages';
import SidebarLayout from 'src/layouts/SidebarLayout';
import { Navigate } from 'react-router-dom';

export const componentRoutes = [
  createRoute('components', <SidebarLayout />, [
    createRoute('', <Navigate to="buttons" replace />),
    createRoute('buttons', <Buttons />),
    createRoute('modals', <Modals />),
    createRoute('accordions', <Accordions />),
    createRoute('tabs', <Tabs />),
    createRoute('badges', <Badges />),
    createRoute('tooltips', <Tooltips />),
    createRoute('avatars', <Avatars />),
    createRoute('cards', <Cards />),
    createRoute('forms', <Forms />)
  ])
];
