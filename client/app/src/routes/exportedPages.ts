import { lazy } from 'react';
import Loader from './Loader';
// 遅延読み込みじゃないもの
import Overview from '../content/overview';
import Status404 from '../content/pages/Status/Status404';
import Status500 from '../content/pages/Status/Status500';
import StatusComingSoon from '../content/pages/Status/ComingSoon';
import StatusMaintenance from '../content/pages/Status/Maintenance';
export { Overview, Status404, Status500, StatusComingSoon, StatusMaintenance };
// 遅延読み込み（呼び出されるまではロードしない）主ページなど
export const Dashboard = Loader(lazy(() => import('src/content/dashboards/Dashboard')));
export const Crypto = Loader(lazy(() => import('src/content/dashboards/Crypto')));
export const Messenger = Loader(lazy(() => import('src/content/applications/Messenger')));
export const Transactions = Loader(lazy(() => import('src/content/applications/Transactions')));
export const UserProfile = Loader(lazy(() => import('src/content/applications/Users/profile')));
export const UserSettings = Loader(lazy(() => import('src/content/applications/Users/settings')));
export const Buttons = Loader(lazy(() => import('src/content/pages/Components/Buttons')));
export const Modals = Loader(lazy(() => import('src/content/pages/Components/Modals')));
export const Accordions = Loader(lazy(() => import('src/content/pages/Components/Accordions')));
export const Tabs = Loader(lazy(() => import('src/content/pages/Components/Tabs')));
export const Badges = Loader(lazy(() => import('src/content/pages/Components/Badges')));
export const Tooltips = Loader(lazy(() => import('src/content/pages/Components/Tooltips')));
export const Avatars = Loader(lazy(() => import('src/content/pages/Components/Avatars')));
export const Cards = Loader(lazy(() => import('src/content/pages/Components/Cards')));
export const Forms = Loader(lazy(() => import('src/content/pages/Components/Forms')));
export const Moc = Loader(lazy(() => import('../pages/mocks/Mock')));
export const Moc2 = Loader(lazy(() => import('../pages/mocks/Mock2')));