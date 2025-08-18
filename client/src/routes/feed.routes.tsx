import Layout from '@/components/layout';
import FeedPage from '@/pages/Feed';
import DeveloperPage from '@/pages/Feed/Developers';
import { createRoute } from '@tanstack/react-router';
import { rootRoute } from './root-route';

const feedLayout = createRoute({
  getParentRoute: () => rootRoute,
  path: '/',
  component: Layout,
});

const feedPage = createRoute({
  getParentRoute: () => feedLayout,
  path: '/feed',
  component: () => {
    return <FeedPage />;
  },
});

const developerPage = createRoute({
  getParentRoute: () => feedLayout,
  path: '/developer',
  component: () => {
    return <DeveloperPage />;
  },
});

const developerDetailPage = createRoute({
  getParentRoute: () => feedLayout,
  path: '/developer/$id',
  component: () => {
    return <DeveloperPage />;
  },
});

export const feedPages = [
  feedLayout,
  feedPage,
  developerPage,
  developerDetailPage,
];
