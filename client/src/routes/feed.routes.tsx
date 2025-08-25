import Layout from '@/components/layout';
import FeedPage from '@/pages/Feed';
import DeveloperPage from '@/pages/Feed/Developers';
import ProfilePage from '@/pages/Feed/Profile';
import LandingPage from '@/pages/Landing';
import { createRoute } from '@tanstack/react-router';
import { rootRoute } from './root-route';

const feedLayout = createRoute({
  getParentRoute: () => rootRoute,
  path: '',
  component: Layout,
  id: 'feed-layout',
});

const landingPage = createRoute({
  getParentRoute: () => rootRoute,
  path: '/',
  component: LandingPage,
});

const feedPage = createRoute({
  getParentRoute: () => feedLayout,
  path: '/feed',
  component: FeedPage,
});

const developerPage = createRoute({
  getParentRoute: () => feedLayout,
  path: '/developer',
  component: DeveloperPage,
});

const developerDetailPage = createRoute({
  getParentRoute: () => feedLayout,
  path: '/developer/$id',
  component: DeveloperPage,
});

const profilePage = createRoute({
  getParentRoute: () => feedLayout,
  path: '/profile/$id',
  component: ProfilePage,
});

export const feedPages = [
  feedLayout,
  landingPage,
  feedPage,
  developerPage,
  developerDetailPage,
  profilePage,
];
