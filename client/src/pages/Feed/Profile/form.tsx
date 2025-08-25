import useUser from '@/hooks/queries/useUser';
import type { UserFormBody, UserTypeItem } from '@/services/users/users.type';
import { Button, FileInput, Modal, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';

interface ProfileFormProps {
  opened: boolean;
  onClose(): void;
  user?: UserTypeItem;
}

function ProfileForm({ onClose, opened, user }: ProfileFormProps) {
  const form = useForm<UserFormBody>({
    initialValues: {
      name: user?.name || '',
      email: user?.email || '',
      bio: user?.bio || '',
      avatar: null,
    },
  });

  const { updateUser } = useUser();

  const handleClose = () => {
    onClose();
    form.reset();
  };

  const handleSubmit = (values: typeof form.values) => {
    const formData = new FormData();

    formData.append('name', values.name);
    formData.append('bio', values.bio || '');
    formData.append('email', values.email);
    formData.append('avatar', values.avatar || '');

    updateUser.mutate(
      { id: user?.id || '', body: formData },
      {
        onSuccess: handleClose,
        onError: (err) => {
          console.log('Error updating user:', err);
        },
      }
    );
  };

  return (
    <Modal
      title="Edit Profile"
      size="lg"
      centered
      opened={opened}
      onClose={handleClose}
    >
      <form onSubmit={form.onSubmit(handleSubmit)}>
        <FileInput
          label="Input Avatar"
          placeholder="Input avatar"
          {...form.getInputProps('avatar')}
        />
        <TextInput
          label="Name"
          placeholder="Insert your name"
          {...form.getInputProps('name')}
        />
        <TextInput
          label="Email"
          placeholder="Insert your email"
          {...form.getInputProps('email')}
        />
        <TextInput
          label="Bio"
          placeholder="Insert your bio"
          {...form.getInputProps('bio')}
        />
        <Button type="submit">Save</Button>
      </form>
    </Modal>
  );
}

export default ProfileForm;
