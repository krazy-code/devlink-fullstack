import {
  Button,
  Card,
  Group,
  LoadingOverlay,
  PasswordInput,
  Space,
  Stack,
  Text,
  TextInput,
  Title,
} from '@mantine/core';
import { useForm } from '@mantine/form';
import { Link } from 'react-router';
import useAuth from '../../hooks/queries/useAuth';
import type { AuthFormBody } from '../../services/auth/auth.types';

interface AuthPageProps {
  type: 'login' | 'register';
}

function AuthPage({ type = 'login' }: AuthPageProps) {
  const { login, register } = useAuth();

  const form = useForm<AuthFormBody>({
    initialValues: { name: '', email: '', password: '' },
    validate: {
      name(value) {
        if (!value && type === 'register') return 'Required';
      },
      email(value) {
        if (!value) return 'Required';
        if (!/^[\w-.]+@([\w-]+\.)+[\w-]{2,4}$/.test(value))
          return 'Email not valid';
      },
      password(value) {
        if (!value) return 'Required';
        if (value.length < 8) return 'Min 8';
        if (value.length > 16) return 'Max 16';
      },
      confirm_password(value, values) {
        if (type === 'register') {
          if (!value && values.password) return 'Required';
          if (value !== values.password) return "Password doesn't match";
        }
      },
    },
  });

  const handleSubmit = (values: AuthFormBody) => {
    if (type === 'login') login.mutate(values);
    else register.mutate(values);
  };

  return (
    <main>
      <LoadingOverlay visible={login.isPending || register.isPending} />
      <Card withBorder p="md" shadow="md">
        <form onSubmit={form.onSubmit(handleSubmit)}>
          <Stack>
            <Title>{type.toUpperCase()}</Title>
            {type === 'register' && (
              <TextInput
                label="Name"
                placeholder="Insert Name"
                {...form.getInputProps('name')}
              />
            )}
            <TextInput
              label="Email"
              placeholder="Insert Email"
              {...form.getInputProps('email')}
            />
            <PasswordInput
              label="Password"
              placeholder="Insert Password"
              {...form.getInputProps('password')}
            />
            {type === 'register' && (
              <PasswordInput
                label="Confirm Password"
                placeholder="Insert Confirm Password"
                {...form.getInputProps('confirm_password')}
              />
            )}
            <Button type="submit">Submit</Button>
          </Stack>
        </form>
        <Space h={10} />
        <Group>
          <Text>Doesn't have account?</Text>
          <Button component={Link} to="/register" variant="transparent">
            Register
          </Button>
        </Group>
      </Card>
    </main>
  );
}

export default AuthPage;
