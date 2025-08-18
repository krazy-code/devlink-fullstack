import DevlinkLogo from '@/components/logo';
import { Box, Button, Group } from '@mantine/core';
import { Link } from '@tanstack/react-router';
import classes from '../Header.module.css';

export default function AdminHeader() {
  return (
    <Box>
      <header className={classes.header}>
        <Group justify="space-between" h="100%">
          <DevlinkLogo />

          <Group visibleFrom="sm">
            <Button variant="default" component={Link} to="/">
              Go to Feed
            </Button>
          </Group>
        </Group>
      </header>
    </Box>
  );
}
