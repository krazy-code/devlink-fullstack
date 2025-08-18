import { createRouter } from '@tanstack/react-router';
import { adminPages } from './admin.routes';
import { authPages } from './auth.routes';
import { feedPages } from './feed.routes';
import { rootRoute } from './root-route';

const routeTree = rootRoute.addChildren([
  ...adminPages,
  ...feedPages,
  ...authPages,
]);

export const router = createRouter({ routeTree });
