import { Box, Card, Center, Stack, Text } from '@mantine/core';

interface DeveloperItemProps {
  name: string;
}

function DeveloperItem({ name }: DeveloperItemProps) {
  return (
    <Card withBorder>
      <Stack>
        <Card.Section>
          <Center>
            <Box
              w={80}
              h={80}
              style={{
                borderRadius: '100%',
                border: '1px solid black',
              }}
            >
              <Text>Image</Text>
            </Box>
          </Center>
        </Card.Section>
        <Text>{name}</Text>
      </Stack>
    </Card>
  );
}

export default DeveloperItem;
