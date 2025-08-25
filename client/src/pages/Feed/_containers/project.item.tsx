import CardMain from '@/components/card-main';
import type { ProjectTypeItem } from '@/services/projects/projects.types';
import { ActionIcon, Avatar, Badge, Group, Stack, Text } from '@mantine/core';
import { IconMessage, IconThumbUp } from '@tabler/icons-react';
import { Link } from '@tanstack/react-router';

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
    <CardMain
      style={{
        border: '1px solid #e1e4e8',
        boxShadow: '0 1px 3px rgba(27,31,35,0.04)',
      }}
    >
      <Stack>
        <Group gap={8}>
          <Avatar src={project.user_avatar || undefined} radius="xl" size={28}>
            {project.user_name?.[0]}
          </Avatar>
          <Link to="/profile/$id" params={{ id: project.user_id }}>
            <Text size="sm" fw={500}>
              {project.user_name}
            </Text>
          </Link>
          <Text size="xs" c="dimmed">
            {project.created_at &&
              new Date(project.created_at).toLocaleDateString()}
          </Text>
        </Group>
        <Text fw="bolder" size="lg">
          {project.title}
        </Text>
        <Text size="sm" c="dimmed">
          {project.description}
        </Text>
        {project.tags && (
          <Group gap={4}>
            {project.tags.map((tag: string) => (
              <Badge key={tag} color="gray" variant="light">
                {tag}
              </Badge>
            ))}
          </Group>
        )}
        <Group>
          <Group gap={2}>
            <ActionIcon
              variant={isLike ? 'filled' : 'transparent'}
              c="black"
              onClick={handleClickLike}
            >
              <IconThumbUp />
            </ActionIcon>
            <Text size="sm">{project.likes ?? 0} Likes</Text>
          </Group>
          <Group gap={2}>
            <ActionIcon variant="transparent" c="black">
              <IconMessage />
            </ActionIcon>
            <Text size="sm">{project.comments ?? 0} Comments</Text>
          </Group>
        </Group>
      </Stack>
    </CardMain>
  );
}

export default ProjectItem;
