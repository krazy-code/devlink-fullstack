import { AppShell, Box, Stack } from '@mantine/core';
import { Outlet } from 'react-router';
import Header from './header';
import SideBar from './side-bar';

function Layout() {
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
            <SideBar />
          </Stack>
        </AppShell.Navbar>
        <AppShell.Main pos="relative">
          <Box pt={{ base: '20px', md: '70px' }}>
            <Outlet />
          </Box>
        </AppShell.Main>
        <AppShell.Navbar
          px="md"
          bg="white"
          withBorder
          style={{ overflow: 'hidden', left: 'calc(100vw - 200px)', right: 0 }}
        >
          <Stack h="100%">
            <SideBar />
          </Stack>
        </AppShell.Navbar>
      </AppShell>
      {/* <ScrollRestoration /> */}
    </>
  );
}

export default Layout;
