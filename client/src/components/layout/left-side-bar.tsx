import { Box, Card, Group, Stack, Text } from '@mantine/core';
import CardMain from '../card-main';

function LeftSideBar() {
  return (
    <CardMain>
      <Stack h="100%">
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
    </CardMain>
  );
}

export default LeftSideBar;
