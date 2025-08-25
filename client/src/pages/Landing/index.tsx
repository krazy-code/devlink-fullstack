import { Button, Container, Group, Stack, Text, Title } from '@mantine/core';
import { Link } from '@tanstack/react-router';

export default function LandingPage() {
  return (
    <Container size="xs" py={60}>
      <Stack align="center" gap={24}>
        <Title order={1} size={36} ta="center">
          Devlink Community
        </Title>
        <Text ta="center" c="dimmed" size="lg">
          Connect, share, and grow with fellow developers. Showcase your
          projects, follow others, and build your network.
        </Text>
        <Group gap={16}>
          <Button component={Link} to="/auth" variant="filled" color="blue">
            Login / Register
          </Button>
          <Button component={Link} to="/feed" variant="outline" color="blue">
            Explore Feed
          </Button>
        </Group>
      </Stack>
    </Container>
  );
}
