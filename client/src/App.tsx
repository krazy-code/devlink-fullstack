import { MantineProvider } from '@mantine/core';
import '@mantine/core/styles.css';
import { ModalsProvider } from '@mantine/modals';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { RouterProvider } from 'react-router';
import './App.css';
import { cssVariablesResolver, defaultTheme } from './lib/theme';
import { routes } from './routes';
import { ProfileProvider } from './services/profile/profile.provider';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: false,
      refetchOnWindowFocus: false,
    },
  },
});

function App() {
  return (
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
            <RouterProvider router={routes} />
          </ProfileProvider>
        </ModalsProvider>
      </MantineProvider>
    </QueryClientProvider>
  );
}

export default App;
