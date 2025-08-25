import { useListFollows } from '@/hooks/queries/useFollow';
import { ProfileContext } from '@/services/profile/profile.provider';
import { Box, Card, Group, Loader, Stack, Text } from '@mantine/core';
import { useContext } from 'react';
import CardMain from '../card-main';

function LeftSideBar() {
  const { profile } = useContext(ProfileContext);
  const { data, isLoading } = useListFollows(profile?.id || '');

  return (
    <CardMain>
      <Stack h="100%" gap={8}>
        <Card radius="lg" p="sm" shadow="sm">
          <Group>
            <Box bg="gray" w={28} h={28} style={{ borderRadius: '100%' }}>
              {profile?.name?.[0] || '?'}
            </Box>
            <Text fw={500}>{profile?.name || 'User'}</Text>
          </Group>
        </Card>
        <Card radius="lg" p="sm" shadow="sm">
          <Text size="sm" fw={600} mb={4}>
            Following
          </Text>
          {isLoading && <Loader size="xs" />}
          {!isLoading && (!data || data.length === 0) && (
            <Text size="xs" c="dimmed">
              You are not following anyone yet.
            </Text>
          )}
          {!isLoading && data && data.length > 0 && (
            <Stack gap={4}>
              {data.map((follow) => (
                <Group key={follow.id} gap={6}>
                  <Box
                    bg="blue"
                    w={20}
                    h={20}
                    style={{ borderRadius: '100%' }}
                  />
                  <Text size="sm">{follow.followed_id}</Text>
                  {/* TODO: Replace followed_id with developer/user name by joining with user/dev data */}
                </Group>
              ))}
            </Stack>
          )}
        </Card>
      </Stack>
    </CardMain>
  );
}

export default LeftSideBar;
