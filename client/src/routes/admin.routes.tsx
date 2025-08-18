import AdminLayout from '@/components/layout/admin/admin.layout';
import AdminDeveloperPage from '@/pages/Admin/Developer';
import AdminUserPage from '@/pages/Admin/User';
import { createRoute } from '@tanstack/react-router';
import { rootRoute } from './root-route';

const adminPage = createRoute({
  getParentRoute: () => rootRoute,
  path: '/admin',
  component: AdminLayout,
});

const adminUserPage = createRoute({
  getParentRoute: () => adminPage,
  path: '/user',
  component: AdminUserPage,
});

const adminDeveloperPage = createRoute({
  getParentRoute: () => adminPage,
  path: '/developer',
  component: AdminDeveloperPage,
});

export const adminPages = [adminPage, adminUserPage, adminDeveloperPage];
