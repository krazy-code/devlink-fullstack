import { AppShell, Stack } from '@mantine/core';
import { Outlet } from '@tanstack/react-router';
import Header from './header';
import SideBar from './left-side-bar';
import RightSideBar from './right-side-bar';

const sideBarWidth = 400;

function Layout() {
  return (
    <>
      <AppShell
        header={{ height: 100 }}
        navbar={{
          width: sideBarWidth,
          breakpoint: 'lg',
          collapsed: { mobile: true },
        }}
        aside={{
          width: sideBarWidth,
          breakpoint: 'lg',
          collapsed: { mobile: true },
        }}
        withBorder={false}
      >
        <AppShell.Header>
          <Header />
        </AppShell.Header>
        <AppShell.Navbar p="md" style={{ overflow: 'hidden' }}>
          <SideBar />
        </AppShell.Navbar>
        <AppShell.Main pos="relative">
          <Stack w="100%" py="md">
            <Outlet />
          </Stack>
        </AppShell.Main>
        <AppShell.Aside p="md">
          <RightSideBar />
        </AppShell.Aside>
      </AppShell>
    </>
  );
}

export default Layout;
