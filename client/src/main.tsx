import { MantineProvider } from '@mantine/core';
import '@mantine/core/styles.css';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { RouterProvider } from '@tanstack/react-router';
import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import './App.css';
import './index.css';

import { ModalsProvider } from '@mantine/modals';
import { cssVariablesResolver, defaultTheme } from './lib/theme.ts';
import { router } from './routes/index.tsx';
import { ProfileProvider } from './services/profile/profile.provider.tsx';

declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router;
  }
}

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: false,
      refetchOnWindowFocus: false,
    },
  },
});

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <MantineProvider
        theme={defaultTheme}
        cssVariablesResolver={cssVariablesResolver}
      >
        <ModalsProvider
          modalProps={
            {
              // withCloseButton: false,
              // classNames: {
              //   body: classes.modalBody,
              //   header: classes.modalHeader,
              // },
            }
          }
        >
          <ProfileProvider>
            <RouterProvider router={router} />
          </ProfileProvider>
        </ModalsProvider>
      </MantineProvider>
    </QueryClientProvider>
  </StrictMode>
);
