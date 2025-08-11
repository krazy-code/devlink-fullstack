import useDevelopers from '@/hooks/queries/useDevelopers';
import type { DeveloperFormBody } from '@/services/developers/developers.types';
import { Button, Stack, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { useParams } from 'react-router';

interface DeveloperFormProps {
  isEdit?: boolean;
}

function DeveloperForm({ isEdit = false }: DeveloperFormProps) {
  const { createDevelopers, updateDevelopers } = useDevelopers();

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

  const handleSubmit = (values: DeveloperFormBody) => {
    if (isEdit) updateDevelopers.mutate({ ...values, id: developerId });
    else createDevelopers.mutate(values);
  };

  return (
    <div>
      <form onSubmit={form.onSubmit(handleSubmit)}>
        <Stack>
          <TextInput />
          <Button type="submit">Create</Button>
        </Stack>
      </form>
    </div>
  );
}

export default DeveloperForm;
