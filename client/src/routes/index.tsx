import { createBrowserRouter } from 'react-router';
import AuthPage from '../pages/Auth';
import { adminRoutes } from './admin.routes';
import { publicRoutes } from './public.routes';

export const routes = createBrowserRouter([
  { path: '/login', element: <AuthPage type="login" /> },
  { path: '/register', element: <AuthPage type="register" /> },
  ...adminRoutes,
  ...publicRoutes,
]);
