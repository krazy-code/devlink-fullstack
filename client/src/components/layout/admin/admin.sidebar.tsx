import CardMain from '@/components/card-main';
import { NavLink, Stack } from '@mantine/core';
import { NavLink as NavRouter } from 'react-router';

function AdminSidebar() {
  return (
    <CardMain>
      <Stack h="100%">
        {navs.map((nav) => {
          return (
            <NavLink
              key={nav.href}
              component={NavRouter}
              title={nav.text}
              to={nav.href}
              label={nav.text}
              variant="subtle"
            />
          );
        })}
      </Stack>
    </CardMain>
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
