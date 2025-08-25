import AuthPage from '@/pages/Auth';
import { createRoute } from '@tanstack/react-router';
import { rootRoute } from './root-route';

const loginPage = createRoute({
  getParentRoute: () => rootRoute,
  path: '/auth',
  component: () => <AuthPage type="login" />,
});

const registerPage = createRoute({
  getParentRoute: () => rootRoute,
  path: '/register',
  component: () => <AuthPage type="register" />,
});

export const authPages = [loginPage, registerPage];
