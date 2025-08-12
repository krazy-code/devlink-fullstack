import AdminLayout from '@/components/layout/admin/admin.layout';
import AdminUserPage from '@/pages/Admin/User';
import { createBrowserRouter } from 'react-router';
import Layout from '../components/layout';
import AuthPage from '../pages/Auth';
import DeveloperPage from '../pages/Developers';
import DeveloperDetailPage from '../pages/Developers/detail';
import MainPage from '../pages/Feed';

export const routes = createBrowserRouter([
  { path: '/login', element: <AuthPage type="login" /> },
  { path: '/register', element: <AuthPage type="register" /> },
  {
    path: '/admin',
    element: <AdminLayout />,
    children: [{ path: 'user', element: <AdminUserPage /> }],
  },
  {
    path: '/',
    element: <Layout />,
    children: [
      { index: true, element: <MainPage /> },
      { path: 'projects', element: <MainPage /> },
      {
        path: 'developer',
        children: [
          { index: true, element: <DeveloperPage /> },
          { path: ':id', element: <DeveloperDetailPage /> },
        ],
      },
    ],
  },
]);
