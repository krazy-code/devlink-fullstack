import { Box, Button, Card, Group, Stack, Text } from '@mantine/core';
import CardMain from '../card-main';

function RightSideBar() {
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
        <Button>Post</Button>
      </Stack>
    </CardMain>
  );
}

export default RightSideBar;
