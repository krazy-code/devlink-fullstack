import CardMain from '@/components/card-main';
import type { ProjectTypeItem } from '@/services/projects/projects.types';
import { ActionIcon, Box, Group, Stack, Text } from '@mantine/core';
import { IconMessage, IconThumbUp } from '@tabler/icons-react';

function ProjectItem({
  project,
  handleClickLike,
  isLike,
}: {
  project: ProjectTypeItem;
  handleClickLike(): void;
  isLike: number;
}) {
  return (
    <CardMain>
      <Stack>
        <Group gap={8}>
          <Box w={20} h={20} bg="black" style={{ borderRadius: '100%' }}></Box>
          <Text size="sm">{project.user_name}</Text>
        </Group>
        <Text fw="bolder">{project.title}</Text>
        <Text>{project.description}</Text>
        <Group>
          <Group gap={2}>
            <ActionIcon
              variant={isLike ? 'filled' : 'transparent'}
              c="black"
              onClick={handleClickLike}
            >
              <IconThumbUp />
            </ActionIcon>
            <Text size="sm">21 Likes</Text>
          </Group>
          <Group gap={2}>
            <ActionIcon variant="transparent" c="black">
              <IconMessage />
            </ActionIcon>
            <Text size="sm">Reply</Text>
          </Group>
        </Group>
      </Stack>
    </CardMain>
  );
}

export default ProjectItem;
