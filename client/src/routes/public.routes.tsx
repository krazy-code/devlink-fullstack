import type { RouteObject } from 'react-router';
import Layout from '../components/layout';
import DeveloperPage from '../pages/Developers';
import DeveloperDetailPage from '../pages/Developers/detail';
import MainPage from '../pages/Feed';

export const publicRoutes: RouteObject[] = [
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
];
