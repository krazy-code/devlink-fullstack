import useDevelopers from '@/hooks/queries/useDeveloper';
import type { DeveloperFormBody } from '@/services/developers/developers.types';
import {
  Button,
  Card,
  Group,
  LoadingOverlay,
  Modal,
  Select,
  Stack,
  TextInput,
} from '@mantine/core';
import { useForm } from '@mantine/form';
import { useParams } from 'react-router';

interface AdminDeveloperFormProps {
  isEdit?: boolean;
  opened: boolean;
  onClose(): void;
}

function AdminDeveloperForm({
  isEdit = false,
  opened,
  onClose,
}: AdminDeveloperFormProps) {
  const { createDeveloper, updateDeveloper } = useDevelopers();

  const params = useParams();
  const developerId = params.id as string;

  const form = useForm<DeveloperFormBody>({
    initialValues: {
      bio: '',
      github: '',
      location: '',
      user_id: 0,
      website: '',
    },
  });

  const handleClose = () => {
    onClose();
    form.reset();
  };

  const handleSubmit = (values: DeveloperFormBody) => {
    const body = {
      ...values,
      user_id: Number(values.user_id),
    };
    if (isEdit) updateDeveloper.mutate({ ...body, id: developerId });
    else createDeveloper.mutate(body);
  };

  return (
    <Modal
      opened={opened}
      onClose={handleClose}
      title="Form Developer"
      pos="relative"
      centered
    >
      <LoadingOverlay
        visible={updateDeveloper.isPending || createDeveloper.isPending}
      />
      <Card>
        <form onSubmit={form.onSubmit(handleSubmit)}>
          <Stack>
            <Select
              label="User"
              placeholder="Insert User"
              {...form.getInputProps('user_id')}
              data={[{ value: '1', label: 'Putra' }]}
            />
            <TextInput
              label="Bio"
              placeholder="Insert Bio"
              {...form.getInputProps('bio')}
            />
            <TextInput
              label="Github"
              placeholder="Insert Github"
              {...form.getInputProps('github')}
            />
            <TextInput
              label="Location"
              placeholder="Insert Location"
              {...form.getInputProps('location')}
            />

            <TextInput
              label="Website"
              placeholder="Insert Website"
              {...form.getInputProps('website')}
            />
            <Group>
              <Button variant="outline" type="button">
                Close
              </Button>
              <Button type="submit">Create</Button>
            </Group>
          </Stack>
        </form>
      </Card>
    </Modal>
  );
}

export default AdminDeveloperForm;
