import type { RouteObject } from 'react-router';
import Layout from '../components/layout';
import MainPage from '../pages/Feed';
import DeveloperPage from '../pages/Feed/Developers';
import DeveloperDetailPage from '../pages/Feed/Developers/detail';

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
