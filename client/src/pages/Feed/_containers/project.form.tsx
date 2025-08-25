import SelectUser from '@/components/form/selects/select.user';
import useProject from '@/hooks/queries/useProject';
import type { ProjectFormBody } from '@/services/projects/projects.types';
import {
  Button,
  Group,
  LoadingOverlay,
  Modal,
  Select,
  Stack,
  Textarea,
  TextInput,
} from '@mantine/core';
import { useForm } from '@mantine/form';

interface ProjectFormProps {
  opened: boolean;
  onClose(): void;
  isEdit?: boolean;
}

function ProjectForm({ opened, onClose, isEdit = false }: ProjectFormProps) {
  const { createProject, updateProject } = useProject();

  const form = useForm<ProjectFormBody>({
    initialValues: {
      id: '',
      user_id: '',
      user_name: '',
      title: '',
      description: '',
      tech_stacks: '',
      project_url: '',
      created_at: '',
    },
  });
  const handleClose = () => {
    onClose();
    form.reset();
  };

  const handleSubmit = (values: ProjectFormBody) => {
    if (isEdit) updateProject.mutate(values);
    else createProject.mutate(values);
  };

  return (
    <Modal opened={opened} onClose={handleClose}>
      <LoadingOverlay
        visible={updateProject.isPending || createProject.isPending}
      />
      <form onSubmit={form.onSubmit(handleSubmit)}>
        <Stack>
          <TextInput
            {...form.getInputProps('title')}
            label="Title"
            placeholder="Insert Title"
          />
          <SelectUser
            label="Developer"
            placeholder="Choose Developer"
            {...form.getInputProps('user_id')}
          />
          <Textarea
            {...form.getInputProps('description')}
            label="Deskripsi"
            placeholder="Insert Deskripsi"
          />
          <Select data={['React', 'Vite', 'Golang']} />
          <TextInput
            {...form.getInputProps('project_url')}
            label="Project URL"
            placeholder="Insert Project URL"
          />
        </Stack>
        <Group>
          <Button>Close</Button>
          <Button>Submit</Button>
        </Group>
      </form>
    </Modal>
  );
}

export default ProjectForm;
