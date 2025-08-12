import { AppShell, Box, Stack } from '@mantine/core';
import { Outlet } from 'react-router';
import Header from '../header';
import AdminSidebar from './admin.sidebar';

function AdminLayout() {
  return (
    <>
      <AppShell
        header={{ height: 100 }}
        navbar={{
          width: 200,
          breakpoint: 'lg',
          collapsed: { mobile: true },
        }}
        py="md"
        px={{ base: '20px', md: 'xl' }}
        withBorder={true}
      >
        <AppShell.Header>
          <Header />
        </AppShell.Header>
        <AppShell.Navbar
          px="md"
          bg="white"
          withBorder
          style={{ overflow: 'hidden' }}
        >
          <Stack h="100%">
            <AdminSidebar />
          </Stack>
        </AppShell.Navbar>
        <AppShell.Main pos="relative">
          <Box pt={{ base: '20px', md: '70px' }}>
            <Outlet />
          </Box>
        </AppShell.Main>
      </AppShell>
      {/* <ScrollRestoration /> */}
    </>
  );
}

export default AdminLayout;
