import { AppShell, Stack } from '@mantine/core';
import { Outlet } from '@tanstack/react-router';
import AdminHeader from './admin.header';
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
        bg="transparent"
        px={{ base: '20px', md: 'xl' }}
        withBorder={false}
      >
        <AppShell.Header>
          <AdminHeader />
        </AppShell.Header>
        <AppShell.Navbar p="md" bg="transparent" style={{ overflow: 'hidden' }}>
          <AdminSidebar />
        </AppShell.Navbar>
        <AppShell.Main pos="relative">
          <Stack py="md">
            <Outlet />
          </Stack>
        </AppShell.Main>
      </AppShell>
      {/* <ScrollRestoration /> */}
    </>
  );
}

export default AdminLayout;
