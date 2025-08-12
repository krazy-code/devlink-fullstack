import AdminLayout from '@/components/layout/admin/admin.layout';
import AdminDeveloperPage from '@/pages/Admin/Developer';
import AdminUserPage from '@/pages/Admin/User';
import type { RouteObject } from 'react-router';

export const adminRoutes: RouteObject[] = [
  {
    path: '/admin',
    element: <AdminLayout />,
    children: [
      { path: 'user', element: <AdminUserPage /> },
      { path: 'developer', element: <AdminDeveloperPage /> },
    ],
  },
];
