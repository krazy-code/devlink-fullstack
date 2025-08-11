import useDevelopers from '@/hooks/queries/useDeveloper';
import type { DeveloperFormBody } from '@/services/developers/developers.types';
import { Button, Card, Group, Modal, Stack, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { useParams } from 'react-router';

interface DeveloperFormProps {
  isEdit?: boolean;
  opened: boolean;
  onClose(): void;
}

function DeveloperForm({
  isEdit = false,
  opened,
  onClose,
}: DeveloperFormProps) {
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
    if (isEdit) updateDeveloper.mutate({ ...values, id: developerId });
    else createDeveloper.mutate(values);
  };

  return (
    <Modal
      opened={opened}
      onClose={handleClose}
      title="Form Developer"
      centered
    >
      <Modal.Body>
        <Card>
          <form onSubmit={form.onSubmit(handleSubmit)}>
            <Stack>
              <TextInput />
              <Group>
                <Button variant="outline" type="button">
                  Close
                </Button>
                <Button type="submit">Create</Button>
              </Group>
            </Stack>
          </form>
        </Card>
      </Modal.Body>
    </Modal>
  );
}

export default DeveloperForm;
