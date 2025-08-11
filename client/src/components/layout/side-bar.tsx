import { Box, Card, Group, Stack, Text } from '@mantine/core';

function SideBar() {
  return (
    <Stack>
      <Card radius="lg" p="sm" shadow="sm">
        <Group>
          <Box
            bg="gray"
            w={20}
            h={20}
            style={{
              borderRadius: '100%',
            }}
          >
            P
          </Box>
          <Text>Putra Satria</Text>
        </Group>
      </Card>
      <Card radius="lg" p="sm" shadow="sm"></Card>
      <Card radius="lg" p="sm" shadow="sm"></Card>
    </Stack>
  );
}

export default SideBar;
