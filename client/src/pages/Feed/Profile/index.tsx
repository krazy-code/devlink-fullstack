import CardMain from '@/components/card-main';
import useUser from '@/hooks/queries/useUser';
import { siteConfig } from '@/services/_config/site-config';
import { ProfileContext } from '@/services/profile/profile.provider';
import { Box, Button, Card, Group, Image, Stack } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { useParams } from '@tanstack/react-router';
import { useContext } from 'react';
import ProfileForm from './form';

function ProfilePage() {
  const params = useParams({
    from: '/feed-layout/profile/$id',
    select: (params) => ({ id: params.id }),
  });

  const [openedForm, modalForm] = useDisclosure();

  const profileId = params.id;
  const { detailUser } = useUser();
  const { profile } = useContext(ProfileContext);
  const queryDetail = detailUser(profileId);
  const data = queryDetail.data;
  return (
    <CardMain>
      <Stack>
        {profile?.id === profileId && (
          <Button onClick={modalForm.open}>Edit Profile</Button>
        )}
        <Card shadow="sm" p="md" radius="md">
          <Stack>
            <Box
              pos="absolute"
              top={0}
              bottom="30%"
              left={0}
              right={0}
              bg="primary"
            />
            <Box
              pos="absolute"
              bottom={0}
              left={0}
              right={0}
              top="70%"
              bg="black"
            />
            <Group style={{ zIndex: 1 }} pos="relative" align="end" gap={16}>
              <Box
                w={70}
                h={70}
                pos="relative"
                style={{ borderRadius: '100%', overflow: 'hidden' }}
              >
                <Image
                  src={
                    data?.avatar
                      ? `${siteConfig?.host}${data?.avatar}`
                      : undefined
                  }
                  fit="cover"
                  style={{
                    maxWidth: '100%',
                    maxHeight: '100%',
                    placeItems: 'center',
                  }}
                />
              </Box>
            </Group>
          </Stack>
        </Card>
      </Stack>
      <ProfileForm opened={openedForm} onClose={modalForm.close} user={data} />
    </CardMain>
  );
}

export default ProfilePage;
