import { NavLink, Stack } from '@mantine/core';
import { NavLink as NavRouter } from 'react-router';

function AdminSidebar() {
  return (
    <Stack>
      {navs.map((nav) => {
        return <NavLink component={NavRouter} title={nav.text} to={nav.href} />;
      })}
    </Stack>
  );
}

export default AdminSidebar;

const navs = [
  {
    text: 'User',
    href: '/admin/user',
  },
  {
    text: 'Developer',
    href: '/admin/developer',
  },
];
